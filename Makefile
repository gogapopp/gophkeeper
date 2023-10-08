BINARY_SERVER=server.exe
BINARY_CLIENT=./cmd/client/client.exe
VERSION=`git describe`
COMMIT=`git rev-parse HEAD`

all: test build

build:
    go build -ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT)" -o $(BINARY_SERVER) ./cmd/server/main.go
    go build -ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT)" -o $(BINARY_CLIENT) ./cmd/client/main.go

run: build
    ./cmd/server/main.go
    ./cmd/client/main.go
test:
	go test -v ./...