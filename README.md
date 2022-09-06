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
export TAG=0.1.8; CGO_ENABLED=0 go build -o briscolad main.go; docker build --rm -t mcaci/briscola-serv:$TAG .; rm briscolad
```

Running (can also use --detach):

```sh
docker run --rm -d -p 4000:8080 -p 8081:8081 --name briscola-serv mcaci/briscola-serv:$TAG
```

Pushing:

```sh
docker push mcaci/briscola-serv:$TAG
```

### Build for debugging

Without optimization to avoid surprises while setting breakpoints.

```sh
docker build --rm -t mcaci/briscola-serv-debug:$TAG .
```

Here -N will disable optimization and `-l` disable inlining. This removes surprises when setting breakpoints.

To be confirmed if push is needed.

```sh
docker push mcaci/briscola-serv-debug:$TAG
```

Next step:
Start the debug with an ephemeral container

```sh
POD_NAME=`kubectl get pod -l app.kubernetes.io/name=briscola-serv --no-headers -o custom-columns=':metadata.name'`
# for console
kubectl debug -it $POD_NAME --image=busybox:1.28 --target=briscola-serv
# missing to solve the ptrace issue
kubectl debug -it $POD_NAME --image=mcaci/briscola-serv-debug:latest --target=briscola-serv
# done with integrated container
kubectl attach -it $POD_NAME -c debug
```

### Deploying on KinD steps examples

Here are the commands to run:

```sh
kind create cluster --config ./deployment/kind/conf.yaml
kubectl create ns briscola-serv-ns
kubectl create ns metallb-system
helm install briscola-serv ./deployment/briscola-serv
kubectl label namespace default istio-injection=enabled --overwrite
# kind delete pod
```

If the tag needs to be changed run `helm upgrade briscola-serv ./deployment/briscola-serv --set image.tag=$TAG`.

If a change in the dependencies (like metallb) is needed run `helm dependency update ./deployment/briscola-serv/`.

To test the deployment it is possible to run either of the two after adjusting the IP address to the one taken from the load balancer's external address:

```sh
$ curl -XPOST -d '{"number":1}' http://172.18.255.200:8080/points
{"points":11}
$ go run main.go -grpc 172.18.255.200:8081 -cli points 1
11
```

For more information read Kind's LoadBalancer [documentation](https://kind.sigs.k8s.io/docs/user/loadbalancer/).

When done run `kind delete cluster --name briscola-serv-cluster` to dispose of it.

### More debugging info

https://kubernetes.io/docs/tasks/configure-pod-container/share-process-namespace/

curl -v -XPATCH -H "Content-Type: application/json-patch+json" \
'http://127.0.0.1:8001/api/v1/namespaces/default/pods/nginx-8f458dc5b-wkvq4/ephemeralcontainers' \
--data-binary @- << EOF
[{
"op": "add", "path": "/spec/ephemeralContainers/-",
"value": {
"command":[ "/bin/sh" ],
"stdin": true, "tty": true,
"image": "nicolaka/netshoot",
"name": "debug-strace",
"securityContext": {"capabilities": {"add": ["SYS_PTRACE"]}},
"targetContainerName": "nginx" }}]
EOF

kubectl proxy --port=8080

curl -v -XPATCH -H "Content-Type: application/json-patch+json" "http://127.0.0.1:8080/api/v1/namespaces/default/pods/$POD_NAME/ephemeralcontainers" --data-binary @- << EOF
[{
"op": "add", "path": "/spec/ephemeralContainers/-",
"value": {
"command":[ "dlv", "attach", "1" ],
"stdin": true, "tty": true,
"image": "mcaci/briscola-serv-debug":latest,
"name": "debug-briscola-serv",
"securityContext": {"capabilities": {"add": ["SYS_PTRACE"]}},
"targetContainerName": "briscola-serv" }}]
EOF

kubectl exec -it $POD_NAME -c debug-briscola-serv" -- bash

"dlv", "attach", "1"