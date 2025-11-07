.DEFAULT_GOAL:=help
-include .makerc

# --- Config -----------------------------------------------------------------

GOMODS=$(shell find . -type f -name go.mod)
# Newline hack for error output
define br


endef

# --- Targets -----------------------------------------------------------------

# This allows us to accept extra arguments
%: .mise .husky go.work
	@:

# Ensure go.work file
go.work:
	@go work init
	@go work use -r .
	@go work sync

.PHONY: .mise
# Install dependencies
.mise: msg := $(br)$(br)Please ensure you have 'mise' installed and activated!$(br)$(br)$$ brew update$(br)$$ brew install mise$(br)$(br)See the documentation: https://mise.jdx.dev/getting-started.html$(br)$(br)
.mise:
ifeq (, $(shell command -v mise))
	$(error ${msg})
endif
	@mise install

.PHONY: .husky
# Configure git hooks for husky
.husky:
	@git config core.hooksPath .husky

### Tasks

.PHONY: check
## Run lint & test
check: tidy examples lint test

.PHONY: tidy
## Run go mod tidy
tidy:
	@echo "〉go mod tidy"
	@$(foreach mod,$(GOMODS), (cd $(dir $(mod)) && echo "📂 $(dir $(mod))" && go mod tidy) &&) true

.PHONY: lint
## Run linter
lint:
	@echo "〉golangci-lint run"
	@$(foreach mod,$(GOMODS), (cd $(dir $(mod)) && echo "📂 $(dir $(mod))" && golangci-lint run) &&) true

.PHONY: lint.fix
## Fix lint violations
lint.fix:
	@echo "〉golangci-lint run fix"
	@$(foreach mod,$(GOMODS), (cd $(dir $(mod)) && echo "📂 $(dir $(mod))" && golangci-lint run --fix) &&) true

.PHONY: test
## Run tests
test:
	@echo "〉go test"
	@$(foreach mod,$(GOMODS), (cd $(dir $(mod)) && echo "📂 $(dir $(mod))" && GO_TEST_TAGS=-skip go test -coverprofile=coverage.out -tags=safe -race ./...) &&) true

.PHONY: outdated
## Show outdated direct dependencies
outdated:
	@echo "〉go mod outdated"
	@go list -u -m -json all | go-mod-outdated -update -direct

.PHONY: build
## Build binary
build:
	@echo "〉go build bin/gotsrpc"
	@rm -f bin/gotsrpc
	@go build -o bin/gotsrpc cmd/gotsrpc/gotsrpc.go

.PHONY: build.debug
## Build binary in debug mode
build.debug:
	@echo "〉go build bin/gotsrpc (debug)"
	@rm -f bin/gotsrpc
	@go build -gcflags "all=-N -l" -o bin/gotsrpc cmd/gotsrpc/gotsrpc.go

.PHONY: install
## Run go install
install:
	@echo "〉installing gotsrpc"
	@go install cmd/gotsrpc/gotsrpc.go

.PHONY: install.debug
## Run go install with debug
install.debug:
	@echo "〉installing gotsrpc (debug)"
	@go install -gcflags "all=-N -l" cmd/gotsrpc/gotsrpc.go

EXAMPLES=context
#EXAMPLES=basic errors monitor nullable union time types context
define examples
.PHONY: example.$(1)
example.$(1):
	@echo "📝  example: ${1}"
	@cd example/${1} && go run ../../cmd/gotsrpc/gotsrpc.go gotsrpc.yml
	@-cd example/${1}/client && ../../node_modules/.bin/tsc --build

.PHONY: example.$(1).run
example.$(1).run: example.${1}
	@cd example/${1} && go run main.go

.PHONY: example.$(1).debug
example.$(1).debug: build.debug
	@cd example/${1} && dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ../../bin/gotsrpc gotsrpc.yml
endef
$(foreach p,$(EXAMPLES),$(eval $(call examples,$(p))))

.PHONY: examples
## Generate examples
examples:
	@echo "〉Generating examples"
	@for name in example/*/; do\
		if [ $$name != "example/node_modules/" ]; then \
			$(MAKE) example.`basename $${name}`;\
		fi \
  done
.PHONY: examples

### Utils

.PHONY: docs
## Open go docs
docs:
	@go doc -http

.PHONY: help
## Show help text
help:
	@echo "gotsrpc\n"
	@echo "Usage:\n  make [task]"
	@awk '{ \
		if($$0 ~ /^### /){ \
			if(help) printf "%-23s %s\n\n", cmd, help; help=""; \
			printf "\n%s:\n", substr($$0,5); \
		} else if($$0 ~ /^[a-zA-Z0-9._-]+:/){ \
			cmd = substr($$0, 1, index($$0, ":")-1); \
			if(help) printf "  %-23s %s\n", cmd, help; help=""; \
		} else if($$0 ~ /^##/){ \
			help = help ? help "\n                        " substr($$0,3) : substr($$0,3); \
		} else if(help){ \
			print "\n                        " help "\n"; help=""; \
		} \
	}' $(MAKEFILE_LIST)
