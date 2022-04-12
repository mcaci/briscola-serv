# briscola-serv

Service for operations related to briscola game

## How to run

Two terminals are needed to run the services for the briscola game

### Terminal 1

Terminal 1 is where you start the server that listens to briscola commands

```shell
go run main.go -d
```

### Terminal 2

Terminal 2 is where you execute briscola commands

#### Using HTTP transport layer

```shell
$ curl -XPOST -d '{"number":1}' http://localhost:8080/points
{"points":11}
$ curl -XPOST -d '{"numbers":[1, 2, 3]}' http://localhost:8080/count
{"points":21}
$ curl -XPOST -d '{"firstCardNumber":1, "firstCardSeed":2, "secondCardNumber":3, "secondCardSeed":1, "briscolaSeed":1}' http://localhost:8080/compare
{"secondCardWins":true}
```

#### Using gRPC transport layer

```shell
$ go run main.go -cli points 1
11
$ go run main.go -cli points 10
4
$ go run main.go -cli count 1 2 10 3 5 8 6 4 4 2 5
27
$ go run main.go -cli compare 1 2 3 1 5 
# server side error
$ go run main.go -cli compare 1 2 3 1 3
false
$ go run main.go -cli compare 1 2 3 1 1
true
```

### Docker steps examples

Building: 

```sh
CGO_ENABLED=0 go build -o briscolad main.go
docker build -t mcaci/briscola-serv:v0.0.1 .
rm briscolad
```

Running (can also use --detach):

```sh
docker run --rm -it -p 4000:8080 -p 8081:8081 mcaci/briscola-serv:v0.0.1
```

Pushing:

```sh
docker push mcaci/briscola-serv:v0.0.1
```
