export GOPATH:=$(shell pwd)

default: all

deps:
	go get -d -v eighty/...

build: deps
	go install eighty/cmd/...

clean:
	rm -rf bin pkg src/github.com src/gopkg.in

all: build
