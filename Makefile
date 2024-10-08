BINARY_NAME := todolist-web

all: run

cross-build:
	GOARCH=amd64 GOOS=linux   go build -o ./bin/$(BINARY_NAME)-linux ./cmd/main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/$(BINARY_NAME)-win.exe ./cmd/main.go

build:
	go build -o ./bin/$(BINARY_NAME) ./cmd/main.go

run:
	go run ./cmd/todolist-web/main.go

clean:
	rm ./bin/*
