all: start

.PHONY: build
build: service-a service-b

.PHONY: start
start: build
	./scripts/start.sh

.PHONY: stop
stop:
	./scripts/stop.sh

.PHONY: service-a
service-a:
	@docker build -t service-a -f service-a/Dockerfile .

.PHONY: service-b
service-b:
	@docker build -t service-b -f service-b/Dockerfile .

