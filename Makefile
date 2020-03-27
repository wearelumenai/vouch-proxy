DOCK=docker-compose -p clusapp
DOCKER=$(DOCK) -f deployments/docker-compose.yml
DOCKER_DEV=$(DOCK) -f deployments/docker-compose.dev.yml
# Variable for filename for store running procees id
PID_FILE=/tmp/clusauth.pid
# We can use such syntax to get main.go and other root Go files.
GO_FILES=$(wildcard *.go)

all: build test

.PHONY: get
get:
	go get -v ./...

.PHONY: build
build: get
	go build

.PHONY: install
install: get
	go install

.PHONY: test
test:
	./do.sh test
	./do.sh coverage

.PHONY: start
start:
	vouch-proxy

.PHONY: start-pid
start-pid:
	vouch-proxy
	
#& echo $$! > $(PID_FILE)

# go run $(GO_FILES) & echo $$! > $(PID_FILE)
# You can also use go build command for start task
# start:
#   go build -o /bin/clusauth . && \
#   /bin/clusauth & echo $$! > $(PID_FILE)

# Stop task will kill process by ID stored in PID_FILE (and all child processes by pstree).
.PHONY: stop
stop:
	pkill vouch-proxy || exit 0
#kill -9 `cat $(PID_FILE)`

# -kill `pstree -p \`cat $(PID)\` | tr "\n" " " |sed "s/[^0-9]/ /g" |sed "s/\s\s*/ /g"`

# Before task will only prints message. Actually, it is not necessary. You can remove it, if you want.
.PHONY: before
before: build
	@echo "STOPED clusauth" && printf '%*s\n' "40" '' | tr ' ' -

# Restart task will execute stop, before and start tasks in strict order and prints message.
.PHONY: restart
restart: stop before start-pid
	@echo "STARTED clusauth" && printf '%*s\n' "40" '' | tr ' ' -

# Serve task will run fswatch monitor and performs restart task if any source file changed. Before serving it will execute start task.
.PHONY: dev
dev: start-pid
	fswatch -or --event=Updated ./ | \
	xargs -n1 -I {} make restart

# dev:
	# ./do.sh watch

.PHONY: docker-up
docker-up:
	$(DOCKER) up -d ${args}

.PHONY: docker-build
docker-build:
	$(DOCKER) build ${args}

.PHONY: docker-stop
docker-stop:
	$(DOCKER) stop ${args}

.PHONY: docker-down
docker-down:
	$(DOCKER) down ${args}

.PHONY: docker-logs
docker-logs:
	$(DOCKER) logs -f ${args}

.PHONY: docker-restart
docker-restart:
	$(DOCKER) restart ${args}

.PHONY: docker-config
docker-config:
	$(DOCKER) config ${args}

.PHONY: docker-dev-up
docker-dev-up:
	$(DOCKER_DEV) up -d ${args}

.PHONY: docker-dev-build
docker-dev-build:
	$(DOCKER_DEV) build ${args}

.PHONY: docker-dev-stop
docker-dev-stop:
	$(DOCKER_DEV) stop ${args}

.PHONY: docker-dev-down
docker-dev-down:
	$(DOCKER_DEV) down ${args}

.PHONY: docker-dev-logs
docker-dev-logs:
	$(DOCKER_DEV) logs -f ${args}

.PHONY: docker-dev-restart
docker-dev-restart:
	$(DOCKER_DEV) restart -f ${args}

.PHONY: docker-dev-config
docker-dev-config:
	$(DOCKER_DEV) config -f ${args}
