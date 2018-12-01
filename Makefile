SHELL:=/bin/bash

build: mkartifactsdir build-linux-amd64 build-linux-386 build-darwin-amd64 build-darwin-386 build-windows-amd64 build-windows-386

mkartifactsdir:
	(test ! -d artifacts && mkdir artifacts) || true

build-linux-amd64:
	GO111MODULE="on" GOOS="linux" GOARCH="amd64" go build -ldflags "-s -w" -o artifacts/nocors-server-linux-amd64.bin

build-linux-386:
	GO111MODULE="on" GOOS="linux" GOARCH="386" go build -ldflags "-s -w" -o artifacts/nocors-server-linux-i386.bin

build-darwin-amd64:
	GO111MODULE="on" GOOS="darwin" GOARCH="amd64" go build -ldflags "-s -w" -o artifacts/nocors-server-darwin-amd64.bin

build-darwin-386:
	GO111MODULE="on" GOOS="darwin" GOARCH="386" go build -ldflags "-s -w" -o artifacts/nocors-server-darwin-i386.bin

build-windows-amd64:
	GO111MODULE="on" GOOS="windows" GOARCH="amd64" go build -ldflags "-s -w" -o artifacts/nocors-server-windows-amd64.exe

build-windows-386:
	GO111MODULE="on" GOOS="windows" GOARCH="386" go build -ldflags "-s -w" -o artifacts/nocors-server-windows-i386.exe

clean:
	rm -rf artifacts

.PHONY: build