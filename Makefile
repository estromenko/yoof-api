
BIN=main
BUILDFLAGS=-o $(BIN)

.all: build
.PHONY: build clean docs

build: docs
	go build $(BUILDFLAGS) cmd/app/main.go

docs:
	cmd/swag init -g cmd/app/main.go

clean:
	rm -rf $(BIN)
