.DEFAULT_GOAL:=help

.PHONY: test
## Run go test
test:
	go test -v ./...

.PHONY: install
## Run go install
install:
	go install cmd/gotsrpc/gotsrpc.go

## Run go install with debug
install.debug:
	go install -gcflags "all=-N -l" cmd/gotsrpc/gotsrpc.go

.PHONY: outdated
## Show outdated direct dependencies
outdated:
	go list -u -m -json all | go-mod-outdated -update -direct

## Run goreleaser
build:
	goreleaser build --snapshot --rm-dist

## run go build with debug
build.debug:
	rm -f bin/gotsrpc
	go build -gcflags "all=-N -l" -o bin/gotsrpc cmd/gotsrpc/gotsrpc.go


## === Tools ===

EXAMPLES=basic errors nullable union
define examples
.PHONY: example.$(1)
example.$(1):
	cd example/${1} && go run ../../cmd/gotsrpc/gotsrpc.go gotsrpc.yml
	cd example/${1}/client && tsc --build

.PHONY: example.$(1).run
example.$(1).run: example.${1}
	cd example/${1} && go run main.go

.PHONY: example.$(1).debug
example.$(1).debug: build.debug
	cd example/${1} && dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ../../bin/gotsrpc gotsrpc.yml

.PHONY: example.$(1).lint
example.$(1).lint:
	cd example/${1} && golangci-lint run
endef
$(foreach p,$(EXAMPLES),$(eval $(call examples,$(p))))

## Run go mod tidy recursive
.PHONY: lint
lint:
	# @golangci-lint run
	@for name in example/*/; do\
		echo "-------- $${name} ------------";\
		sh -c "cd $$(pwd)/$${name} && golangci-lint run";\
  done

## Run go mod tidy recursive
.PHONY: gomod
gomod:
	@go mod tidy
	@for name in example/*/; do\
		echo "-------- $${name} ------------";\
		sh -c "cd $$(pwd)/$${name} && go mod tidy";\
  done

## === Examples ===

.PHONY: examples
## Build examples
examples:
	@for name in example/*/; do\
		echo "-------- $${name} ------------";\
		$(MAKE) example.`basename $${name}`;\
  done
.PHONY: examples

## === Utils ===

## Show help text
help:
	@awk '{ \
			if ($$0 ~ /^.PHONY: [a-zA-Z\-\_0-9]+$$/) { \
				helpCommand = substr($$0, index($$0, ":") + 2); \
				if (helpMessage) { \
					printf "\033[36m%-23s\033[0m %s\n", \
						helpCommand, helpMessage; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^[a-zA-Z\-\_0-9.]+:/) { \
				helpCommand = substr($$0, 0, index($$0, ":")); \
				if (helpMessage) { \
					printf "\033[36m%-23s\033[0m %s\n", \
						helpCommand, helpMessage"\n"; \
					helpMessage = ""; \
				} \
			} else if ($$0 ~ /^##/) { \
				if (helpMessage) { \
					helpMessage = helpMessage"\n                        "substr($$0, 3); \
				} else { \
					helpMessage = substr($$0, 3); \
				} \
			} else { \
				if (helpMessage) { \
					print "\n                        "helpMessage"\n" \
				} \
				helpMessage = ""; \
			} \
		}' \
		$(MAKEFILE_LIST)
