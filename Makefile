.DEFAULT_GOAL:=help
-include .makerc

# --- Targets -----------------------------------------------------------------

# This allows us to accept extra arguments
%: .husky
	@:

.PHONY: .husky
# Configure git hooks for husky
.husky:
	@if ! command -v husky &> /dev/null; then \
		echo "ERROR: missing executeable 'husky', please run:"; \
		echo "\n$ go install github.com/go-courier/husky/cmd/husky@latest\n"; \
	fi
	@git config core.hooksPath .husky

## === Tasks ===

.PHONY: check
## Run lint & test
check: tidy examples lint test

.PHONY: test
## Run go test
test:
	@GO_TEST_TAGS=-skip go test -coverprofile=coverage.out --tags=safe -race ./...

.PHONY: install
## Run go install
install:
	@go install cmd/gotsrpc/gotsrpc.go

.PHONY: install.debug
## Run go install with debug
install.debug:
	@go install -gcflags "all=-N -l" cmd/gotsrpc/gotsrpc.go

.PHONY: outdated
## Show outdated direct dependencies
outdated:
	@go list -u -m -json all | go-mod-outdated -update -direct

.PHONY: build.debug
## Build binary in debug mode
build.debug:
	@rm -f bin/gotsrpc
	@go build -gcflags "all=-N -l" -o bin/gotsrpc cmd/gotsrpc/gotsrpc.go

## === Tools ===

EXAMPLES=basic errors monitor nullable union time
define examples
.PHONY: example.$(1)
example.$(1):
	cd example/${1} && go run ../../cmd/gotsrpc/gotsrpc.go gotsrpc.yml
	cd example/${1}/client && ../../node_modules/.bin/tsc --build

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

.PHONY: lint
## Run linter
lint:
	@golangci-lint run

.PHONY: lint.fix
## Run linter and fix
lint.fix:
	@golangci-lint run --fix

.PHONY: tidy
## Run go mod tidy
tidy:
	@go mod tidy

## === Examples ===

.PHONY: examples
## Build examples
examples:
	@for name in example/*/; do\
		if [ $$name != "example/node_modules/" ]; then \
			echo "-------- $${name} ------------";\
			$(MAKE) example.`basename $${name}`;\
		fi \
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
