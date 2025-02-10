.PHONY: dev clear build start prepare
	
dev:
	@air -- | pino-pretty -c -t

clear:
	@rm -rf tmp server

build:
	@go build -o server ./cmd/server.go

start:
	./server

test-integration:
	@go test -v ./cmd