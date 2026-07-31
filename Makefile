.DEFAULT_GOAL:=help
-include .makerc

# --- Config -----------------------------------------------------------------

GOMODS=$(shell find . -type f -name go.mod)
# Newline hack for error output
define br


endef

# --- Targets -----------------------------------------------------------------

# This allows us to accept extra arguments
%: .mise .lefthook go.work
	@:

# Ensure go.work file
go.work:
	@go work init
	@go work use -r .
	@go work sync

.PHONY: .mise
# Install dependencies
.mise:
ifeq (, $(shell command -v mise))
	$(error $(br)$(br)Please ensure you have 'mise' installed and activated!$(br)$(br)  $$ brew update$(br)  $$ brew install mise$(br)$(br)See the documentation: https://mise.jdx.dev/getting-started.html)
endif
	@mise install

.PHONY: .lefthook
# Configure git hooks for lefthook
.lefthook:
	@lefthook install --reset-hooks-path

### Tasks

.PHONY: check
## Run lint & tests
check: tidy examples generate lint.fix test.race audit

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
	@GO_TEST_TAGS=-skip go test -coverprofile=coverage.out -tags=safe work

.PHONY: test.race
## Run tests with -race
test.race:
	@echo "〉go test -race"
	@GO_TEST_TAGS=-skip go test -coverprofile=coverage.out -tags=safe -race work

.PHONY: test.nocache
## Run tests with -count=1
test.nocache:
	@echo "〉go test -count=1"
	@GO_TEST_TAGS=-skip go test -coverprofile=coverage.out -tags=safe -count=1 work

.PHONY: test.bench
## Run tests with -bench
test.bench:
	@echo "〉go test -bench"
	@GO_TEST_TAGS=-skip go test -tags=safe -bench=. -benchmem . -run ^$ | tee benchmarks.out

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

EXAMPLES=basic monitor
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

.PHONY: generate
## Run go generate
generate:
	@echo "〉go generate"
	@go generate work

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

### Security

.PHONY: audit
## Run security audit
audit:
	@echo "〉security audit"
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@$(foreach mod,$(GOMODS), (cd $(dir $(mod)) && echo "📂 $(dir $(mod))" && govulncheck ./...) &&) true

### Dependencies

.PHONY: tidy
## Run go mod tidy
tidy:
	@echo "〉go mod tidy"
	@$(foreach mod,$(GOMODS), (cd $(dir $(mod)) && echo "📂 $(dir $(mod))" && go mod tidy) &&) true
	@go work sync

.PHONY: outdated
## Show outdated direct dependencies
outdated:
	@echo "〉go mod outdated"
	@$(foreach mod,$(GOMODS), (cd $(dir $(mod)) && echo "📂 $(dir $(mod))" && go mod tidy && go list -u -m -json all | go-mod-outdated -update -direct) &&) true

.PHONY: upgrade
## Upgrade direct dependencies in all go.mod files
upgrade:
	@echo "〉go mod upgrade"
	@rm -f go.work go.work.sum
	@$(foreach mod,$(GOMODS), (cd $(dir $(mod)) && echo "📂 $(dir $(mod))" && go mod tidy && deps=$$(go list -u -m -f '{{if and (not .Main) (not .Indirect) .Update}}{{.Path}}{{end}}' all); [ -z "$$deps" ] || for dep in $$deps; do go get "$$dep@latest"; done; go mod tidy) &&) true

### Documentation

.PHONY: docs
## Open docs
docs:
	@echo "〉starting docs"
	@cd docs && bun install && bun run dev

.PHONY: docs.build
## Open docs
docs.build:
	@echo "〉building docs"
	@cd docs && bun install && bun run build

.PHONY: godocs
## Open go docs
godocs:
	@echo "〉starting go docs"
	@go doc -http

### Utils

.PHONY: help
# https://patorjk.com/software/taag/#p=display&f=Tmplr&t=gotsrpc&x=none&v=4&h=4&w=80&we=false
## Show help text
help: g=\033[0;32m
help: b=\033[0;34m
help: w=\033[0;90m
help: e=\033[0m
help:
	@echo "$(g)"
	@echo ""
	@echo "┏┓┏┓╋┏┏┓┏┓┏"
	@echo "┗┫┗┛┗┛┛ ┣┛┗"
	@echo " ┛      ┛"
	@echo "with ❤ foomo by bestbytes"
	@echo "$(e)"
	@echo "$(b)Usage:$(e)\n  make [task]"
	@awk '{ \
		if($$0 ~ /^### /){ \
			if(help) printf "  %-21s $(w)%s$(e)\n\n", cmd, help; help=""; \
			printf "$(b)\n%s:$(e)\n", substr($$0,5); \
		} else if($$0 ~ /^[a-zA-Z0-9._-]+:/){ \
			cmd = substr($$0, 1, index($$0, ":")-1); \
			if(help) printf "  %-21s $(w)%s$(e)\n", cmd, help; help=""; \
		} else if($$0 ~ /^##/){ \
			help = help ? help "\n                        " substr($$0,3) : substr($$0,3); \
		} else if(help){ \
			print "\n                        $(w)" help "$(e)\n"; help=""; \
		} \
	}' $(MAKEFILE_LIST)
	@echo ""

