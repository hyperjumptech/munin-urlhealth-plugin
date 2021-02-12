GOPATH=$(shell go env GOPATH)
GO111MODULE=on

.PHONY: all test clean build docker

format:
	go fmt ./...

build-windows: format
	mkdir -p build/windows
	env GOOS=windows GOARCH=amd64 go build -o build/windows ./...

build-linux: format
	mkdir -p build/linux
	env GOOS=linux GOARCH=amd64 go build -o build/linux ./...

build-macos: format
	mkdir -p build/macos
	env GOOS=darwin GOARCH=amd64 go build -o build/macos ./...

build: build-windows build-linux build-macos
	echo "Done"