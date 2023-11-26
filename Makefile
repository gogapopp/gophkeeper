BINARY_SERVER=./cmd/server/server.exe
BINARY_CLIENT=./cmd/client/client.exe
VERSION=`git describe`
COMMIT=`git rev-parse HEAD`

all: test build

build:
	go build -ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT)" -o $(BINARY_SERVER) ./cmd/server/
	go build -ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT)" -o $(BINARY_CLIENT) ./cmd/client/

run: build
	$(BINARY_SERVER)
	$(BINARY_CLIENT)

test:
	go test -v ./...