.PHONY: generate
generate:
	rm -vf demo/output/*
	go run cmd/gotsrpc/gotsrpc.go demo/config.yml
	tsc --outFile cmd/demo/demo.js demo/demo.ts

.PHONY: demo
demo: generate
	cd cmd/demo && go run demo.go

.PHONY: install
install:
	go install cmd/gotsrpc/gotsrpc.go

install.debug:
	go install -gcflags "all=-N -l" cmd/gotsrpc/gotsrpc.go

.PHONY: test
test:
	go test -v ./...

.PHONY: outdated
outdated:
	go list -u -m -json all | go-mod-outdated -update -direct

build:
	goreleaser build --snapshot --rm-dist

build.debug:
	rm -f bin/gotsrpc
	go build -gcflags "all=-N -l" -o bin/gotsrpc cmd/gotsrpc/gotsrpc.go

debug.demo: build.debug
	 dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec bin/gotsrpc demo/config.yml