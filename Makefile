export GOPATH:=$(shell pwd)

default: all

deps-assets:
	go get github.com/jteeuwen/go-bindata
	GOOS="" GOARCH="" go install github.com/jteeuwen/go-bindata/go-bindata
	bin/go-bindata -o src/eighty/assets/bindata.go -pkg="assets" -prefix=assets -nocompress assets/...

deps: deps-assets
	go get -d -v eighty/...

build: deps
	go install eighty/cmd/...

clean:
	rm -rf bin pkg src/github.com src/gopkg.in src/eighty/assets

all: build
