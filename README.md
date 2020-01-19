# briscola-serv

Service for operations related to briscola game

## How to run

Two terminals are needed to run the services for the briscola game

### Terminal 1

Terminal 1 is where you start the server that listens to briscola commands

```shell
go run cmd/briscolad/main.go
```

### Terminal 2

Terminal 2 is where you execute briscola commands

#### Using HTTP transport layer

```shell
$ curl -XPOST -d '{"number":1}' http://localhost:8080/points
11
```

#### Using gRPC transport layer

```shell
$ go run cmd/briscolacli/main.go points 1
11
$ go run cmd/briscolacli/main.go points 10
4
$ go run main.go cmd/briscolacli/count 1 2 10 3 5 8 6 4 4 2 5
27
$ go run cmd/briscolacli/main.go compare 1 2 3 1 5
false
$ go run cmd/briscolacli/main.go compare 1 2 3 1 1
true
```
