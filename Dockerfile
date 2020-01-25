FROM golang as build-env

LABEL author="Michele Caci <michele.caci@gmail.com>"

RUN go get -v -d -u -t github.com/mcaci/briscola-serv/... \
    google.golang.org/grpc \
    github.com/golang/protobuf/protoc-gen-go \
    github.com/golang/protobuf/proto; \
    go install github.com/golang/protobuf/protoc-gen-go; \
    curl -L https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip -o /tmp/protoc.zip; \
    apt-get update; \
    apt-get install unzip -y; \
    mkdir -p /home/protoc; \
    unzip /tmp/protoc.zip -d /home/protoc

WORKDIR /go/src/github.com/mcaci/briscola-serv

RUN /home/protoc/bin/protoc pb/briscola.proto pb/compare.proto pb/count.proto pb/points.proto --go_out=plugins=grpc:.; \
    CGO_ENABLED=0 go build -o briscolad cmd/briscolad/main.go;

FROM scratch

WORKDIR /app
COPY --from=build-env /go/src/github.com/mcaci/briscola-serv/briscolad /app

EXPOSE 8080 8081

ENTRYPOINT ["./briscolad"]
