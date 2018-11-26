.PHONY: generate
generate:
	rm -vf demo/output/*
	rm -vf demo/output-commonjs/*
	go run cmd/gotsrpc/gotsrpc.go demo/config.yml
	go run cmd/gotsrpc/gotsrpc.go -skipgotsrpc demo/config-commonjs.yml
	tsc --outFile cmd/demo/demo.js demo/demo.ts

.PHONY: demo
demo: generate
	cd cmd/demo && go run demo.go

.PHONY: install
install:
	GOBIN=/usr/local/bin go install cmd/gotsrpc/gotsrpc.go

.PHONY: build
build: goreleaser
	goreleaser  --skip-publish --skip-validate

.PHONY: release
release: goreleaser
	goreleaser --rm-dist

.PHONY: goreleaser
goreleaser:
	@go get github.com/goreleaser/goreleaser && go install github.com/goreleaser/goreleaser

.PHONY: glide
glide:
	@go get github.com/Masterminds/glide && glide install

.PHONY: test
test:
	go test -v ./...
	
