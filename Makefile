
.PHONY: demo
demo:
	cd demo && gotsrpc -skipgotsrpc config.yml

.PHONY: install
install:
	GOBIN=/usr/local/bin go install cmd/gotsrpc/gotsrpc.go

release: goreleaser glide
	goreleaser --rm-dist

goreleaser:
	@go get github.com/goreleaser/goreleaser && go install github.com/goreleaser/goreleaser

glide:
	@go get github.com/Masterminds/glide && glide install

test: demo
	go test $(glide nv)
