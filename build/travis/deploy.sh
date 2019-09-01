#!/usr/bin/env bash

# Build binaries and images
make build
docker build -t dnilosek/fib-overkill-api -f api/dockerfile .
docker build -t dnilosek/fib-overkill-worker -f worker/dockerfile .
docker build -t dnilosek/fib-overkill-web ./web

# Push to dockerhub
docker push dnilosek/fib-overkill-api
docker push dnilosek/fib-overkill-worker
docker push dnilosek/fib-overkill-web
