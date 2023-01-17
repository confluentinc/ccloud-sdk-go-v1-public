ifeq ($(shell uname),Darwin)
SED := gsed
else
SED := sed
endif

CLEAN_TARGETS += protoc-clean prototool-clean
GO_EXTRA_DEPS := install-protoc protoc-binaries install-mocker install-hub install-prototool install-protoc-gen-validate
TEST_TARGETS += mods-vet
GO_TEST_TARGET := lint-go run-go-tests
BUILD_GO_OVERRIDE := generate-mocks compile
GO111MODULE := on
GO_USE_VENDOR ?= -mod=vendor
export GO111MODULE

include ./mk-include/cc-begin.mk
include ./mk-include/cc-protoc.mk
include ./mk-include/cc-semver.mk
include ./mk-include/cc-semaphore.mk
include ./mk-include/cc-go.mk
include ./mk-include/cc-vault.mk
include ./mk-include/cc-maven.mk
include ./mk-include/cc-end.mk

GOPATH := $(shell go env GOPATH)
export GOPATH

GIT_STATUS=$(shell (git status --porcelain | grep -q .) && echo dirty || echo clean)

ifeq ($(CI),true)
PATH := $(GOPATH)/bin:$(PATH)
export PATH
endif

# override maven path
ifeq ($(CI), true)
MVN := $(SEMAPHORE_GIT_DIR)/mvnw $(MAVEN_ARGS)
endif

# Exclude linting and formatting generated protobuf and k8s client code
ALL_SRC = $(shell find . -type d -path ./vendor -prune -o -type d -path ./.gomodcache -prune -o -type d -path ./.semaphore-cache -prune -o -type d -path ./client -prune -o -name \*.go -not -name bindata.go -not -name \*.pb.go -print)

PROTOC_VERSION := 3.9.0

PROTOTOOL := $(BIN_PATH)/prototool
PROTOTOOL_VERSION ?= v1.9.0
PROTOTOOL_INSTALLED_VERSION := $(shell $(PROTOTOOL) version --json | jq .version -r | xargs printf "v%s")

PROTO_GEN_VALIDATE_VERSION ?= dab397f346bc74e424ea52606a1c1d3dd7b2d2d6

GO_WHICH := goenv which
ifeq ($(shell which goenv >/dev/null; echo $$?), 1)
GO_WHICH := which
endif

UNAME := $(shell uname)

.PHONY: install-hub
install-hub:
ifeq ($(UNAME),Darwin)
	mkdir -p /tmp/hub && \
		curl -L -o /tmp/hub/hub.tgz https://github.com/github/hub/releases/download/v2.14.2/hub-darwin-amd64-2.14.2.tgz && \
		tar -C /tmp/hub -xzvf /tmp/hub/hub.tgz && cp /tmp/hub/hub-darwin-amd64-2.14.2/bin/hub $(BIN_PATH)/hub
else ifeq ($(UNAME),Linux)
	mkdir -p /tmp/hub && \
		curl -L -o /tmp/hub/hub.tgz https://github.com/github/hub/releases/download/v2.14.2/hub-linux-amd64-2.14.2.tgz && \
		tar -C /tmp/hub -xzvf /tmp/hub/hub.tgz && cp /tmp/hub/hub-linux-amd64-2.14.2/bin/hub $(BIN_PATH)/hub
else
	echo 'Please install hub yourself by following https://github.com/github/hub/tree/v2.14.2#installation.'
endif
	hub --version


.PHONY: install-mocker
install-mocker:
	go install github.com/travisjeffery/mocker/cmd/mocker@v1.1.0

.PHONY: protoc-binaries
protoc-binaries:
	go install github.com/gogo/protobuf/protoc-gen-gogo@v1.3.2
	go install github.com/gogo/googleapis/protoc-gen-gogogoogleapis@v1.4.1
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install github.com/ckaznocha/protoc-gen-lint@v0.2.4
	go install github.com/travisjeffery/proto-go-sql/protoc-gen-sql@v0.0.0-20190911121832-39ff47280e87
	go install github.com/confluentinc/proto-go-setter/protoc-gen-setter@v0.3.0
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.9.5
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v1.9.5
	# Install protoc-gen-structs, since we already have it just install it from local
	(cd cmd/protoc-gen-structs && go install .)

.PHONY: install-prototool
install-prototool:
ifneq ($(PROTOTOOL_VERSION),$(PROTOTOOL_INSTALLED_VERSION))
ifeq ($(CI),true)
# Install barebones prototool
	curl -sSL \
		https://github.com/uber/prototool/releases/download/${PROTOTOOL_VERSION}/prototool-$$(uname -s)-$$(uname -m) \
		-o $(PROTOTOOL) && \
		chmod +x $(PROTOTOOL)
else
# Install prototool with bash completions and man pages.
	curl -sSL \
		https://github.com/uber/prototool/releases/download/${PROTOTOOL_VERSION}/prototool-$$(uname -s)-$$(uname -m).tar.gz | \
		tar -C $(BIN_PATH)/.. --strip-components 1 -xz
endif
endif

# Installed differently because we can't use modules, because they require go 1.14
.PHONY: install-protoc-gen-validate
install-protoc-gen-validate:
	mkdir -p $(GOPATH)/src/github.com && \
		cd $(GOPATH)/src/github.com && \
		rm -rf envoyproxy && mkdir envoyproxy && \
		cd envoyproxy && \
		git clone https://github.com/confluentinc/protoc-gen-validate.git && \
		cd protoc-gen-validate && \
		git checkout $(PROTO_GEN_VALIDATE_VERSION) && \
		GOBIN=$(GOPATH)/bin make build

.PHONY: prototool-clean
prototool-clean:
	rm -rf $(PROTOTOOL)

.PHONY: protoc-lint
protoc-lint:
	@./scripts/lint $(PROTOC)

.PHONY: prototool-breaks
prototool-breaks:
	@./scripts/break_check.sh

.PHONY: generate
generate: generate-mocks generate-modules

.PHONY: generate-mocks
generate-mocks:
	@./scripts/generate-mocks

.PHONY: generate-modules
generate-modules:
	@./scripts/mod

.PHONY: compile
compile:
	@./scripts/compile $(PROTOC)
	@make code-generator
	@echo "compiled ok"
	@make fmt
	@echo "format ok"

.PHONY: code-generator
code-generator:
	@test -d code-generator && rm -rf ./code-generator || true
	@git clone https://github.com/kubernetes/code-generator.git
	@git -C code-generator checkout kubernetes-1.15.7
	@./code-generator/generate-groups.sh all \
	  github.com/confluentinc/cc-structs/client \
	  github.com/confluentinc/cc-structs \
	  "operator:v1" \
	  --go-header-file $(PWD)/header.txt
	@rm -rf ./code-generator

.PHONY: protoc-clean
protoc-clean:
	rm -f `which protoc-gen-gogo`
	rm -f `which protoc-gen-lint`
	rm -f `which protoc-gen-structs`
	rm -f `which protoc-gen-sql`
	rm -f `which protoc-gen-setter`
	rm -f `which modvendor`
	rm -f `which protoc-gen-grpc-gateway`
	rm -f `which protoc-gen-swagger`
	rm -f `which protoc-gen-gogogoogleapis`
	rm -rf ./vendor

.PHONY: symlink-protos
symlink-protos:
	@./scripts/symlink-protos-for-maven

.PHONY: mvn-install
mvn-install: symlink-protos

.PHONY: mvn-set-bumped-version
mvn-set-bumped-version:
		$(MVN) versions:set \
			-DnewVersion=$(BUMPED_CLEAN_VERSION) \
			-DgenerateBackupPoms=false
		git add --verbose pom.xml

.PHONY: tag-release
tag-release:
	# Delete tag if it already exists
	git tag -d $(BUMPED_VERSION) || true
	git push $(GIT_REMOTE_NAME) :$(BUMPED_VERSION) || true
	@./scripts/tagmods $(BUMPED_VERSION)
	git push $(GIT_REMOTE_NAME) $(RELEASE_BRANCH) --tags

.PHONY: push-changes
push-changes:
ifeq ($(RELEASE_BRANCH),$(_empty))
ifeq ($(GIT_STATUS),dirty)
	# commit and push and compilation changes, mainly for the CI to make compiles consistent
	git add --all
	git commit -m "chore: compile"
	git push -u $(GIT_REMOTE_NAME) $(SEMAPHORE_GIT_PR_BRANCH)
endif
endif

.PHONY: mods-vet
mods-vet:
	@./scripts/vet.sh

.PHONY: run-go-tests
run-go-tests:
	@./scripts/test.sh
