#!/bin/sh
kind create cluster --config kind/conf.yaml
helm repo add metallb https://metallb.github.io/metallb
helm install metallb metallb/metallb --version v0.12.1 --values metallb/values.yaml
helm install briscola-serv ./deployment