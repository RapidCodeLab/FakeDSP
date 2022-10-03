include .env
SERVER_BINARY_NAME=server


build:
	go build -o .bin/${SERVER_BINARY_NAME} cmd/server/main.go
run:
	go build -o .bin/${SERVER_BINARY_NAME} cmd/server/main.go
	chmod +x .bin/${SERVER_BINARY_NAME}
	./.bin/${SERVER_BINARY_NAME}
clean:
	go clean
	rm .bin/${SERVER_BINARY_NAME}    
