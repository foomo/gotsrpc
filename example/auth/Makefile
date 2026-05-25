.PHONY: build-client run-server run clean

# Build the TypeScript client to dist directory
build-client:
	cd client && npx tsc

# Run the server
run-server:
	go run main.go

# Build client and run server
run: build-client run-server

# Generate gotsrpc code
generate:
	rm -f client/src/*_gen.ts
	rm -f service/*_gen.go
	gotsrpc gotsrpc.yml

# Clean compiled JavaScript files
clean:
	rm -rf client/dist/*.js

# Full build: generate + build client + run server
dev: generate build-client run-server
