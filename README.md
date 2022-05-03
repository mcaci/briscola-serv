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
export TAG=0.1.3; CGO_ENABLED=0 go build -o briscolad main.go; docker build -t mcaci/briscola-serv:$TAG .; rm briscolad
```

Running (can also use --detach):

```sh
docker run --rm -it -p 4000:8080 -p 8081:8081 mcaci/briscola-serv:0.0.1
```

Pushing:

```sh
docker push mcaci/briscola-serv:0.0.1
```

### Deploying on KinD steps examples

Here are the commands to run:

```sh
kind create cluster --config ./deployment/kind/conf.yaml
helm install briscola-serv ./deployment/app
```

To test the deployment it is possible to run either of the two after adjusting the IP address to the one taken from the load balancer's external address:

```sh
$ curl -XPOST -d '{"number":1}' http://172.18.255.200:8080/points
{"points":11}
$ go run main.go -grpc 172.18.255.200:8081 -cli points 1
11
```

For more information read Kind's LoadBalancer [documentation](https://kind.sigs.k8s.io/docs/user/loadbalancer/).

When done run `kind delete cluster` to dispose of it.
