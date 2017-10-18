
.PHONY: demo
demo:
	rm -vf demo/output/*
	rm -vf demo/output-commonjs/*
	go run cmd/gotsrpc/gotsrpc.go demo/config.yml
	go run cmd/gotsrpc/gotsrpc.go -skipgotsrpc demo/config-commonjs.yml
	tsc --outFile cmd/demo/demo.js demo/demo.ts 
	cd cmd/demo && go run demo.go

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
