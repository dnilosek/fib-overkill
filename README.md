# Fibonacci-Overkill

This is a wildly over-egineered fibonacci calculator, created primarliy for testing a multi-service deploy using kubernetes. 

- [Design](#design)
- [Build](#build)
  * [Requirements](#requirements)
  * [Pre-build requirements](#pre-build-requirements)
  * [Building](#building)
  * [Running](#running)

## Design
![Design Diagram](assets/diag.jpg?raw=true "Diagram")

The goal of this system is to have a user enter the index of a fibonacci number to be calculated. The API layer managers storing that number in [Postgres](https://www.postgresql.org/) (you know, for storage reasons) and also sends the index off to be calculated by a fibonnaci calculation worker. This is done using a pub/sub design with [Redis](https://redis.io/) as the mediator. The UI is written using [React.js](https://reactjs.org/), and the API service as well as the worker are written in [Go](https://golang.org/).

The system is deployed for local sandboxing with [Kubernetes](https://kubernetes.io/) via [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)

## Build
### Requirements
- Minikube 1.11.0
- Golang 1.12.0
- NPM 5.8.0

### Pre-build requirements
Just a couple things to ensure we can use minikube as a sanbox environment. First we will be using the internal docker repository to store the service images, so we have to connect to the minikube docker environment (unfortunetly this needs to be done every time you open up a shell):

```bash
eval $(minikube docker-env)
```

Then we just have to make sure minikube will let us use the ingress

```bash
minikube addons enable ingress
```
### Building
To build and deploy everything to minikube:

```bash
make build-and-deploy
```
If you want to remove everything (K8s deploy, docker images, and binaries)
```bash
make destroy
```
### Running
The  UI and API service will be available via the minikube IP, this can be found using:
```bash
minikube ip
```
You may get a security warning trying to access the site as there is no SSL cert. You should see something that looks like this:

![UI](assets/ui.png?raw=true "UI")

