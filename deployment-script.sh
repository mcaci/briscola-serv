#!/bin/sh
kind create cluster --config conf/kind.yaml
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/namespace.yaml
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/metallb.yaml
#kubectl get pods -n metallb-system --watch
#docker network inspect -f '{{.IPAM.Config}}' kind
kubectl apply -f conf/metallb-cm.yaml
kubectl apply -f blueprints/deployment.yaml
#curl -XPOST -d '{"number":1}' http://172.19.255.200:8080/points