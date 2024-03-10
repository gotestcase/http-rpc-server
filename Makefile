build:
	@go build -o bin/rpc server.go

run: build
	@./bin/rpc

client:
	@go build -o bin/client main.go
	@./bin/client

test:
	@go test rpc_test.go -v