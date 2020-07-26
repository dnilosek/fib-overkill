GOSRC=./...
COVERDIR=.cover
COVERFILE=.cover.out
VERSIONFILE=version
BINPATH=bin

WORKERPATH=worker/cmd
WORKERTARGET=worker
WORKERBIN=$(patsubst %, ${BINPATH}/%, $(WORKERTARGET))
WORKERDOCKER=build/docker/dockerfile.worker
WORKERIMAGE=fib-overkill-worker

APIPATH=api/cmd
APITARGET=fib-api
APIBIN=$(patsubst %, ${BINPATH}/%, $(APITARGET))
APIDOCKER=build/docker/dockerfile.api
APIIMAGE=fib-overkill-api

WEBDOCKER=build/docker/dockerfile.web
WEBIMAGE=fib-overkill-web

BINARIES=$(WORKERBIN) $(APIBIN)

ENV		?= test
PORT		?= 8080
BUILD_VERSION	?= $(shell cat $(VERSIONFILE) | head -n 1)
BUILD_NUMBER	?= 0
DOCKER_TAG	?= $(ENV)

.DEFAULT_GOAL=test

dep:
	@go get -v -d $(GOSRC)

test:
	@go test -v -race -coverprofile $(COVERFILE) $(GOSRC)

vet:
	@go vet $(GOSRC)

cover: test
	@mkdir -p $(COVERDIR)
	@go tool cover -html=$(COVERFILE) -o $(COVERDIR)/index.html
	@cd $(COVERDIR) && python -m SimpleHTTPServer $(PORT)

run-worker:
	@go run worker/cmd/worker.go

run-api:
	@go run api/cmd/fib-api.go

build: $(BINARIES)

$(WORKERBIN): $(WORKERPATH)/worker.go
	@CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo \
		-o $@ $(WORKERPATH)/worker.go

$(APIBIN): $(APIPATH)/fib-api.go
	@CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo \
		-o $@ $(APIPATH)/fib-api.go

docker-build-web:
	docker build -t $(WEBIMAGE) -f $(WEBDOCKER) .

docker-build-api:
	docker build -t $(APIIMAGE) -f $(APIDOCKER) .

docker-build-worker:
	docker build -t $(WORKERIMAGE) -f $(WORKERDOCKER) .

docker-build: docker-build-web docker-build-api docker-build-worker

deploy-dev:
	kubectl apply -k build/k8s/dev

destroy-dev:
	kubectl delete -k build/k8s/dev

build-and-deploy: build docker-build deploy-dev

destroy: clean destroy-dev clean-images

clean-images:
	@docker rmi $(WEBIMAGE) $(APIIMAGE) $(WORKERIMAGE)

clean:
	@rm -rf $(BINPATH)

setup-logging:
	@helm repo add elastic https://helm.elastic.co
	@curl -O https://raw.githubusercontent.com/elastic/helm-charts/master/elasticsearch/examples/minikube/values.yaml
	@helm install --name-template elasticsearch elastic/elasticsearch -f ./values.yaml 
	@rm ./values.yaml
	@helm install --name-template kibana elastic/kibana
	@helm install --name-template filebeat elastic/filebeat
	@helm install --name-template metricbeat elastic/metricbeat

remove-logging:
	@helm repo add elastic https://helm.elastic.co
	@helm uninstall elasticsearch
	@helm uninstall kibana
	@helm uninstall filebeat
	@helm uninstall metricbeat

start-logging:
	@kubectl port-forward deployment/kibana-kibana 5601

.PHONY: dep test vet cover run-worker run-api build docker-build-web docker-build-api docker-build-worker docker-build deploy-dev destroy-dev build-and-deploy destroy clean-images clean setup-logging start-logging
