
.PHONY: demo
demo:
	cd demo && gotsrpc -skipgotsrpc config.yml

.PHONY: install
install:
	GOBIN=/usr/local/bin go install cmd/gotsrpc/gotsrpc.go