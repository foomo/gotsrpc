.PHONY: generate
generate:
	rm -vf demo/output/*
	rm -vf demo/output-commonjs/*
	rm -vf demo/output-commonjs-async/*
	go run cmd/gotsrpc/gotsrpc.go demo/config.yml
	go run cmd/gotsrpc/gotsrpc.go -skipgotsrpc demo/config-commonjs.yml
	go run cmd/gotsrpc/gotsrpc.go -skipgotsrpc demo/config-commonjs-async.yml
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

