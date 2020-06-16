.PHONY : test

CURDIR=$(shell pwd)
DOCKER_TARGET= hub.deepin.com/dcmc/cms
DOCKER_BUILDER_TARGET=$(DOCKER_TARGET).builder
TARGET=gingo_server

build:
	go build -o $(CURDIR)/$(TARGET) main.go

run: build
	[ -f conf.yaml ] || cp conf.example.yaml conf.yaml
	SERVER_ROOT=$(CURDIR) $(CURDIR)/$(TARGET)

docker-build:
	docker image inspect golang >/dev/null || docker pull golang:alpine
	docker image inspect alpine >/dev/null || docker pull alpine

docker:
	DOCKER_BUILDKIT=1 docker build -f Dockerfile -t $(DOCKER_TARGET) .
	DOCKER_BUILDKIT=1 docker build -f Dockerfile -t $(DOCKER_TARGET):dev .

docker-release:
	docker push $(DOCKER_TARGET)
	docker push $(DOCKER_TARGET):dev
docker-push:
	docker push $(DOCKER_TARGET)

migrate:
	echo SERVER_ROOT=$(CURDIR)

test:
	echo cd $(CURDIR)/test; SERVER_ROOT=$(CURDIR) ginkgo -v -r -cover

clean:
	rm $(CURDIR)/$(TARGET)