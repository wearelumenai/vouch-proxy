DOCKER=docker-compose -f deployments/docker-compose.yml

all: build test

build:
	./do.sh goget
	./do.sh build

start:
	./vouch-proxy

docker:
	$(DOCKER) up ${args}

docker-build:
	$(DOCKER) build ${args}

docker-stop:
	$(DOCKER) stop ${args}

docker-down:
	$(DOCKER) down ${args}

# .PHONY is used for reserving tasks words
.PHONY: build start \
	docker docker-build docker-stop docker-down
