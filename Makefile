COMMIT_SHA ?= $(shell git rev-parse HEAD)
REPONAME ?= noirai
IMAGE_NAME ?= noirai-otel-collector
MIGRATOR_IMAGE_NAME ?= noirai-schema-migrator
CONFIG_FILE ?= ./config/default-config.yaml
DOCKER_TAG ?= latest

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOPATH ?= $(shell go env GOPATH)
GOTEST=go test -v $(RACE)
GOFMT=gofmt
FMT_LOG=.fmt.log
IMPORT_LOG=.import.log

CLICKHOUSE_HOST ?= 127.0.0.1
CLICKHOUSE_PORT ?= 9000

LD_FLAGS ?=


.PHONY: install-tools
install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v2.6.0

.DEFAULT_GOAL := test-and-lint

.PHONY: test-and-lint
test-and-lint: test fmt lint

.PHONY: test
test:
	go test -count=1 -v -race -cover ./...

.PHONY: build
build:
	go build -o .build/${GOOS}-${GOARCH}/noirai-otel-collector ./cmd/noiraiotelcollector
	go build -o .build/${GOOS}-${GOARCH}/noirai-schema-migrator ./cmd/noiraischemamigrator

.PHONY: amd64
amd64:
	make GOARCH=amd64 build

.PHONY: arm64
arm64:
	make GOARCH=arm64 build

.PHONY: build-all
build-all: amd64 arm64

.PHONY: run
run:
	go run cmd/noiraiotelcollector/main.go --config ${CONFIG_FILE}

.PHONY: fmt
fmt:
	@echo Running go fmt on query service ...
	@$(GOFMT) -e -s -l -w .

.PHONY: build-and-push-noirai-collector
build-and-push-noirai-collector:
	@echo "------------------"
	@echo  "--> Build and push noirai collector docker image"
	@echo "------------------"
	docker buildx build --platform linux/amd64,linux/arm64 --progress plain \
		--no-cache --push -f cmd/noiraiotelcollector/Dockerfile \
		--tag $(REPONAME)/$(IMAGE_NAME):$(DOCKER_TAG) .

.PHONY: build-noirai-collector
build-noirai-collector:
	@echo "------------------"
	@echo  "--> Build noirai collector docker image"
	@echo "------------------"
	docker build --build-arg TARGETPLATFORM="linux/amd64" \
		--no-cache -f cmd/noiraiotelcollector/Dockerfile --progress plain \
		--tag $(REPONAME)/$(IMAGE_NAME):$(DOCKER_TAG) .

.PHONY: build-noirai-schema-migrator
build-noirai-schema-migrator:
	@echo "------------------"
	@echo  "--> Build schema migrator docker image"
	@echo "------------------"
	docker build --build-arg TARGETPLATFORM="linux/amd64" \
		--no-cache -f cmd/noiraischemamigrator/Dockerfile --progress plain \
		--tag $(REPONAME)/$(MIGRATOR_IMAGE_NAME):$(DOCKER_TAG) .

.PHONY: build-and-push-noirai-schema-migrator
build-and-push-noirai-schema-migrator:
	@echo "------------------"
	@echo  "--> Build and push schema migrator docker image"
	@echo "------------------"
	docker buildx build --platform linux/amd64,linux/arm64 --progress plain \
		--no-cache --push -f cmd/noiraischemamigrator/Dockerfile \
		--tag $(REPONAME)/$(MIGRATOR_IMAGE_NAME):$(DOCKER_TAG) .

.PHONY: lint
lint:
	@echo "Running linters..."
	@$(GOPATH)/bin/golangci-lint -v --config .golangci.yml run && echo "Done."

.PHONY: install-ci
install-ci: install-tools

.PHONY: test-ci
test-ci: lint

.PHONY: migrator
migrator:
	@echo "------------------"
	@echo "--> Running schema migrator for $(CLICKHOUSE_HOST):$(CLICKHOUSE_PORT)"
	@echo "------------------"
	go run cmd/noiraischemamigrator/main.go sync --dsn "clickhouse://$(CLICKHOUSE_HOST):$(CLICKHOUSE_PORT)" --dev
