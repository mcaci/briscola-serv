kind create cluster --config ./deployment/kind/conf.yaml
kubectl apply -f https://projectcontour.io/quickstart/contour.yaml
kubectl patch daemonsets -n projectcontour envoy -p '{"spec":{"template":{"spec":{"nodeSelector":{"ingress-ready":"true"},"tolerations":[{"key":"node-role.kubernetes.io/control-plane","operator":"Equal","effect":"NoSchedule"},{"key":"node-role.kubernetes.io/master","operator":"Equal","effect":"NoSchedule"}]}}}}'
helm install metallb metallb/metallb
helm install briscola-serv ./deployment/briscola-serv

# install doesn't work at first try
# helm upgrade briscola-serv ./deployment/briscola-serv

# istio
# kubectl label namespace default istio-injection=enabled --overwrite

