export GOPATH:=$(shell pwd)

default: all

deps:
	go get -d -v eighty/...

build: deps
	go install eighty/cmd/...

clean:
	rm -rf bin pkg
	find src -not -name "eighty" -depth 1 -exec rm -rf {} \;

all: build
