DOCKER_VERSION ?= latest
CONFIGURATION_FILE ?= default.yaml
CONFIGURATION_FROM ?= file:/$(CONFIGURATION_FILE)

project_dir := $(dir $(mkfile_path)$(abspath $(lastword $(MAKEFILE_LIST))))

paf:
	echo $(project_dir)
build:
	CGO_ENABLED=0 go build -o dist/swaggerui  ./...
docker-build: build
	docker build -t swaggerui:$(DOCKER_VERSION) .
docker-run: docker-build
	docker run --rm -v $(project_dir)$(CONFIGURATION_FILE):/$(CONFIGURATION_FILE) -e CONFIGURATION_FROM=$(CONFIGURATION_FROM) --name swaggerui -p 8080:8080 swaggerui:latest
