SRC=$(shell find . -name "*.go")

.PHONY: fmt vet test install_deps clean

default: all

all: build fmt vet test

build:
	$(info ******************** building ********************)
	go build

fmt:
	$(info ******************** checking formatting ********************)
	@test -z $(shell gofmt -l $(SRC)) || (gofmt -d $(SRC); exit 1)

test: install_deps vet
	$(info ******************** running tests ********************)
	go test -v ./...

install_deps:
	$(info ******************** downloading dependencies ********************)
	go get -v ./...

vet:
	$(info ******************** vetting ********************)
	go vet ./...

clean:
	rm -rf $(BIN)