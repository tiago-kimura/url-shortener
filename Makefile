.PHONY: usage build test staticcheck get-linter lint docker-build integration-tests run

OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

NOW = $(shell date -u '+%Y%m%d%I%M%S')

GO := go
GO_LINTER := golint
DOCKER := docker
BUILDOS ?= $(shell go env GOHOSTOS)
BUILDARCH ?= amd64
GOFLAGS ?=
ECHOFLAGS ?=
ROOT_DIR := $(realpath .)

BIN := shortener

PKGS = $(shell $(GO) list ./...)

ENVFLAGS ?= CGO_ENABLED=0
BUILDENV ?= GOOS=$(BUILDOS) GOARCH=$(BUILDARCH)
BUILDFLAGS ?= -a -installsuffix cgo $(GOFLAGS) $(GO_LINKER_FLAGS)
EXTLDFLAGS ?= -extldflags "-lm -lstdc++ -static"

usage: Makefile
	@echo $(ECHOFLAGS) "to use make call:"
	@echo $(ECHOFLAGS) "    make <action>"
	@echo $(ECHOFLAGS) ""
	@echo $(ECHOFLAGS) "list of available actions:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'

## test: run unit tests
test:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Running unit tests $(NO_COLOR)"
	@$(LOCAL_VARIABLES) $(ENVFLAGS) $(GO) test $(GOFLAGS) $(PKGS) --cover

## staticcheck: run staticcheck on packages
staticcheck:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Running staticcheck...$(NO_COLOR)"
	@$(GO) get -v honnef.co/go/tools/cmd/staticcheck
	@$(ENVFLAGS) staticcheck $(PKGS)

## get-linter: install linter
get-linter:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Getting linter...$(NO_COLOR)"
	@go get -u golang.org/x/lint/golint

## lint: lint package
lint: get-linter
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Running linter...$(NO_COLOR)"
	@$(GO_LINTER) -set_exit_status $(PKGS)

## build: build all
build: 
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Building binary ($(BUILDOS)/$(BUILDARCH)/$(BIN))...$(NO_COLOR)"
	@echo $(ECHOFLAGS) $(ENVFLAGS) $(BUILDENV) $(GO) build $(BUILDFLAGS) -o bin/$(BUILDOS)_$(BUILDARCH)/$(BIN) ./cmd/shortener
	@$(ENVFLAGS) $(BUILDENV) $(GO) build $(BUILDFLAGS) -o bin/$(BUILDOS)_$(BUILDARCH)/$(BIN) ./cmd/shortener

## docker-build: create the docker image
docker-build:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> build container image...$(NO_COLOR)"
	@ROOT_DIR=$(ROOT_DIR) $(DOCKER) build -t shortener .

## integration-tests: run integration tests
integration-tests: docker-build
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> running integration test...$(NO_COLOR)"
	@for f in $(shell ls ${ROOT_DIR}/integration_tests); \
	do echo "$(OK_COLOR) Execution $${f} test... $(NO_COLOR)"; \
	$(DOCKER) run --rm -i shortener < ${ROOT_DIR}/integration_tests/$${f}/input > ./integration_tests/result_test; \
	diff ${ROOT_DIR}/integration_tests/result_test ${ROOT_DIR}/integration_tests/$${f}/output --color; \
	rm ${ROOT_DIR}/integration_tests/result_test; \
	done

## run: runs application	
run: docker-build
	@echo $(ECHOFLAGS) "$(OK_COLOR) ==> running shortener...$(NO_COLOR)"
	@ROOT_DIR=$(ROOT_DIR) $(DOCKER) run --rm -i shortener < $(file)

start:
	@echo $(ECHOFLAGS) "$(OK_COLOR) ==> starting all containers...$(NO_COLOR)"
	@ROOT_DIR=$(ROOT_DIR) $(DOCKER)-compose up --build -d

stop:
	@echo $(ECHOFLAGS) "$(OK_COLOR) ==> stoping all containers...$(NO_COLOR)"
	@ROOT_DIR=$(ROOT_DIR) $(DOCKER)-compose down

migrate:
	@echo $(ECHOFLAGS) "$(OK_COLOR) ==> running migrate...$(NO_COLOR)"
	@ROOT_DIR=$(ROOT_DIR) $(DOCKER) exec -t url-shortener-db mysql -u root -p < migrations/init.sql