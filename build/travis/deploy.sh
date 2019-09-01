#!/usr/bin/env bash

# Build binaries and images (dep needs to be run already)
make build
docker build -t dnilosek/fib-overkill-api:latest -t dnilosek/fib-overkill-api:$SHA -f api/dockerfile .
docker build -t dnilosek/fib-overkill-worker:latest -t dnilosek/fib-overkill-worker:$SHA -f worker/dockerfile .
docker build -t dnilosek/fib-overkill-web:latest -t dnilosek/fib-overkill-web:$SHA ./web

# Push to dockerhub
docker push dnilosek/fib-overkill-api:latest
docker push dnilosek/fib-overkill-api:$SHA

docker push dnilosek/fib-overkill-worker:latest
docker push dnilosek/fib-overkill-worker:$SHA

docker push dnilosek/fib-overkill-web:latest
docker push dnilosek/fib-overkill-web:$SHA

# Apply k8s configs
kubectl version
kubectl apply -f build/k8s
kubectl rollout restart
