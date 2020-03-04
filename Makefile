GO ?= go
GOFMT ?= gofmt "-s"
SWAGGER ?= ~/go/bin/swag
PACKAGES ?= $(shell $(GO) list ./...)
GOFILES := $(shell find . -name "*.go" -type f)
TESTFOLDER := $(shell $(GO) list ./... | grep -v test)

all: swag test build

.PHONY: install
install: deps
	$(GO) install ./cmd/london
	$(GO) install ./cmd/manchester

.PHONY: build
build: deps
	$(GO) build ./cmd/london-server
	$(GO) build ./cmd/manchester-server

.PHONY: swag
swag: deps
	$(GO) get -u github.com/swaggo/swag/cmd/swag
	$(SWAGGER) init -g ../../cmd/london-server/main.go -o docs/london -d pkg/london
	$(SWAGGER) init -g ../../cmd/manchester-server/main.go -o docs/manchester -d pkg/manchester

.PHONY: test
test:
	echo "mode: atomic" > coverage.out
	for d in $(TESTFOLDER); do \
		$(GO) test -v -coverpkg=./... -covermode=atomic -coverprofile=profile.out $$d > tmp.out; \
		cat tmp.out; \
		if grep -q "FAIL" tmp.out; then \
			rm tmp.out; \
			exit 1; \
		elif grep -q "build failed" tmp.out; then \
			rm tmp.out; \
			exit; \
		fi; \
		if [ -f profile.out ]; then \
			cat profile.out | grep -v "mode:" >> coverage.out; \
			rm profile.out; \
		fi; \
	done

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: vet
vet:
	$(GO) vet $(PACKAGES)

.PHONY: lint
lint:
	@hash golint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u golang.org/x/lint/golint; \
	fi
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;

.PHONY: deps
deps:
	@hash go > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		echo "Install Go language before running this!"; \
	fi