#!/usr/bin/env bash

kubectl apply -f service.yml
kubectl config set-context --current --namespace=etcd
kubectl port-forward svc/grpc-proxy-server 8080:8080
