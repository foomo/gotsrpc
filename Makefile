
.PHONY: demo
demo:
	cd demo && gotsrpc -skipgotsrpc config.yml

.PHONY: install
install:
	GOBIN=/usr/local/bin go install cmd/gotsrpc/gotsrpc.go

build: goreleaser
	goreleaser


goreleaser:
	@go get github.com/goreleaser/goreleaser && go install github.com/goreleaser/goreleaser

glide:
	@go get github.com/Masterminds/glide && glide install