proto:
	protoc	--go_out=./internal/pb --go_opt=paths=source_relative \
    		--go-grpc_out=./internal/pb --go-grpc_opt=paths=source_relative \
    		./gismeteo.proto

lint:
	golangci-lint run --deadline=5m -v --tests=false

build:
	go build -o server github.com/daniilty/clock-ws-hub/cmd/server
