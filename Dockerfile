FROM golang

LABEL author="Michele Caci <michele.caci@gmail.com>"

RUN ["git", "config", "--global", "url.\"git@github.com:\".insteadOf", "\"https://github.com/\""]
RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " > /root/.ssh/config

RUN go get -v -d -u -t github.com/mcaci/briscola-serv/...
RUN go get -v -d -u -t google.golang.org/grpc
RUN go get -v -d -u -t github.com/golang/protobuf/protoc-gen-go
RUN go get -v -d -u -t github.com/golang/protobuf/proto
RUN go install google.golang.org/grpc
RUN go install github.com/golang/protobuf/proto
RUN go install github.com/golang/protobuf/protoc-gen-go
RUN curl -L https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip -o /tmp/protoc.zip
RUN apt-get update
RUN apt-get install unzip -y
RUN mkdir -p /home/protoc
RUN unzip /tmp/protoc.zip -d /home/protoc

WORKDIR /go/src/github.com/mcaci/briscola-serv

RUN ["/home/protoc/bin/protoc", "pb/briscola.proto", "pb/compare.proto", "pb/count.proto", "pb/points.proto", "--go_out=plugins=grpc:."]
RUN go build ./...

EXPOSE 8080 8081

ENTRYPOINT ["go", "run", "cmd/briscolad/main.go"]