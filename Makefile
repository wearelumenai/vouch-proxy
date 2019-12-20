DOCKER=docker-compose -f deployments/docker-compose.yml

all: docker-build build test

build:
	./do.sh goget
	./do.sh build

start:
	./vouch-proxy

docker:
	$(DOCKER) up ${args} vouch-proxy

docker-build:
	$(DOCKER) build ${args}

docker-stop:
	$(DOCKER) stop ${args}

docker-down:
	$(DOCKER) down ${args}

# .PHONY is used for reserving tasks words
.PHONY: build start docker docker-build docker-stop docker-down
