default: start
project:=ms-workspace-demo
service:=ms-demo-golang
COMMIT_HASH = $(shell git rev-parse --verify HEAD)

.PHONY: start
start: 
	docker-compose -p ${project} up -d

.PHONY: stop
stop: 
	docker-compose -p ${project} down

.PHONY: restart
restart: stop start

.PHONY: logs
logs: 
	docker-compose -p ${project} logs -f ${service}

.PHONY: ps
ps: 
	docker-compose -p ${project} ps

.PHONY: shell
shell: 
	docker-compose -p ${project} exec ${service} sh

.PHONY: build
build:
	docker-compose -p ${project} build --no-cache

.PHONY: clean
clean: stop build start

.PHONY: dep-add
dep-add:
	docker-compose -p ${project} exec ${service} go get -u ${package}

.PHONY: dep-verify
dep-verify:
	docker-compose -p ${project} exec ${service} go mod verify

.PHONY: commit-hash
commit-hash:
	@echo $(COMMIT_HASH)

.PHONY: build-release
build-release:
	docker build --target release -t local/${service}:${COMMIT_HASH} .

.PHONY: run-release
run-release:
	docker run -d --name ${service}_${COMMIT_HASH} -p :3737 local/${service}:${COMMIT_HASH} /main
	docker logs -f ${service}_${COMMIT_HASH}