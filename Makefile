DOCK=docker-compose -p clusapp
DOCKER=$(DOCK) -f deployments/docker-compose.yml
DOCKER_DEV=$(DOCK) -f deployments/docker-compose.dev.yml

all: build test

build:
	./do.sh goget
	./do.sh build

install: build
	./do.sh install

test:
	./do.sh test
	./do.sh coverage

start:
	vouch

dev:
	./do.sh watch

docker:
	$(DOCKER) up ${args}

docker-build:
	$(DOCKER) build ${args}

docker-stop:
	$(DOCKER) stop ${args}

docker-down:
	$(DOCKER) down ${args}

docker-logs:
	$(DOCKER) logs ${args}

docker-dev:
	$(DOCKER_DEV) up ${args}

docker-dev-build:
	$(DOCKER_DEV) build ${args}

docker-dev-stop:
	$(DOCKER_DEV) stop ${args}

docker-dev-down:
	$(DOCKER_DEV) down ${args}

docker-dev-logs:
	$(DOCKER_DEV) logs ${args}

# .PHONY is used for reserving tasks words
.PHONY: build start dev \
	docker docker-build docker-stop docker-down docker-logs \
	docker-dev docker-dev-build docker-dev-stop docker-dev-down docker-dev-logs
