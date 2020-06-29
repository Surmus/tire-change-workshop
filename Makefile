GO ?= $(shell which go)
GOFMT := $(shell which gofmt) "-s"
SWAGGER ?= ${GOPATH}/bin/swag
GOLINT ?= ${GOPATH}/bin/golint
PACKAGES ?= $(shell $(GO) list ./...)
GOFILES := $(shell find . -name "*.go" -type f)
TESTFOLDER := $(shell $(GO) list ./... | grep -v test)

all: swag test build_linux build_windows

.PHONY: install
install: deps
	$(GO) install ./cmd/london
	$(GO) install ./cmd/manchester

.PHONY: build_linux
build_linux: deps
	env GOOS=linux GOARCH=amd64 $(GO) build -o ./build/linux64/london-server --tags "linux" -a -tags netgo -ldflags '-w -extldflags "-static"' ./cmd/london-server
	env GOOS=linux GOARCH=amd64 $(GO) build -o ./build/linux64/manchester-server --tags "linux" -a -tags netgo -ldflags '-w -extldflags "-static"' ./cmd/manchester-server

.PHONY: build_windows
build_windows: deps
	@hash gcc-multilib > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		echo "Install gcc-multilib before running this!"; \
	fi
	@hash gcc-mingw-w64 > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		echo "Install gcc-mingw-w64 before running this!"; \
	fi
	env CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 $(GO) build -o ./build/win64/london-server.exe ./cmd/london-server
	env CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 $(GO) build -o ./build/win64/manchester-server.exe ./cmd/manchester-server

.PHONY: swag
swag: deps
	$(GO) get -u github.com/swaggo/swag/cmd/swag
	$(SWAGGER) init -g ../../cmd/london-server/main.go -o api/london -d internal/london
	$(SWAGGER) init -g ../../cmd/manchester-server/main.go -o api/manchester -d internal/manchester

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
	for PKG in $(PACKAGES); do $(GOLINT) -set_exit_status $$PKG || exit 1; done;

.PHONY: deps
deps:
	@hash go > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		echo "Install Go language before running this!"; \
	fi