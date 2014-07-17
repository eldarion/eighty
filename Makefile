export GOPATH:=$(shell pwd)

default: all

build:
	go install eighty/cmd/...

all: build
