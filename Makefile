.PHONY : test

CURDIR=$(shell pwd)
DOCKER_TARGET= hub.docker.com/tangheng1995/gingo/server
TARGET=gingo_server

build:
	go build -mod vendor -v -o $(CURDIR)/$(TARGET) cmd/main.go

run: build
	SERVER_ROOT=$(CURDIR) $(CURDIR)/$(TARGET)

docker-build:
	docker image inspect golang >/dev/null || docker pull golang:alpine
	docker image inspect alpine >/dev/null || docker pull alpine

docker:
	DOCKER_BUILDKIT=1 docker build -f deployments/Dockerfile -t $(DOCKER_TARGET) .

docker-release:
	docker push $(DOCKER_TARGET)
docker-push:
	docker push $(DOCKER_TARGET)

test:
	go test $(CURDIR)/test -v -coverprofile=coverage.out -covermode=count -coverpkg=./...

coverage: test
	find . -name "coverage.out.*" -exec tail +2 {} >> coverage.out \;
	go tool cover -html coverage.out -o coverage.html
	go tool cover -func=coverage.out
	golangci-lint run -E=golint --out-format checkstyle ./... > report.xml

analyze: coverage
	sonar-scanner -X