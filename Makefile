GOSRC=./...
COVERDIR=.cover
COVERFILE=.cover.out
VERSIONFILE=version
BINPATH=bin

WORKERPATH=worker/cmd
WORKERTARGET=worker
WORKERBIN=$(patsubst %, ${BINPATH}/%, $(WORKERTARGET))

APIPATH=api/cmd
APITARGET=fib-api
APIBIN=$(patsubst %, ${BINPATH}/%, $(APITARGET))

BINARIES=$(WORKERBIN) $(APIBIN)

ENV		?= test
PORT		?= 8080
BUILD_VERSION	?= $(shell cat $(VERSIONFILE) | head -n 1)
BUILD_NUMBER	?= 0
DOCKER_TAG	?= $(ENV)

.DEFAULT_GOAL=test

print-%:  
	@echo $* = $($*)

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

clean:
	@rm -rf $(BINPATH)

compose: build
	 docker-compose build

.PHONY: dep test vet cover run-worker run-api build clean compose
