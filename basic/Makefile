BIN_DIR := ./bin
GO_BUILD_OPTS := -o $(BIN_DIR)/$@

build: $(patsubst cmd/%/,%,$(wildcard cmd/*/))
%: $(wildcard cmd/$@/*.go) | install
	go build $(GO_BUILD_OPTS) ./cmd/$@

install:
	@mkdir -p $(BIN_DIR)

clean:
	rm -rf $(BIN_DIR)/*

test:
	go test -v ./...

cover:
	go test -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
	open cover.html

.PHONY: build install clean test cover
