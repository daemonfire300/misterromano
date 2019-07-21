# Mister Romano

A sample project utilizing:

* Go Modules
* An empty Makefile to prove that we know that later we might need a build tool
* A Dockerfile because we want to containerize our application
* Gorilla Mux for URL parameters
* A roman numbers conversion library that looks good enough based on its unittests


# Usage

`go mod vendor` if a local vendor is needed, i.e. IDE autocompletion still relies on a local vendor folder for certain IDEs, or just use docker to build an artifact `docker build . -t romano1`
and run with `docker run --detach -it --rm -p 8080:8080 --name romano romano1`.

# Open API

The project contains a very thinly documented open-api spec. Might be missing certain things.
Just there to prove that we prove that we can document an api using state-of-the-art
tools.

# Outlook

This project does use the standard http library and intentionally does not utilize
Go's capability to serve via TLS. We assume this application will be containerized and
deployed behind a reverse proxy, e.g., NGINX, Traeffik or Istio (Service Mesh). Thus
the TLS handling (https, encryption, yada yada) will be handled by that component.

Multiple roman numbers can lead to the same arabic number, thus this service does not provide
a disambiguous 1-to-1 mapping.

Since this service is state-less it is a prime target for horizontal auto-scaling, i.e.,
just add more machines running this service. For example using Kubernetes + Istio and 
Kubernetes native auto-scaling features (ReplicaSet+Service = WIN).
