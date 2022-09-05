kind create cluster --config ./deployment/kind/conf.yaml
helm install briscola-serv ./deployment/briscola-serv
# kubectl label namespace default istio-injection=enabled --overwrite

