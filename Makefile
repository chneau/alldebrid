.SILENT:
.ONESHELL:
.NOTPARALLEL:
.EXPORT_ALL_VARIABLES:
.PHONY: run deps build clean exec

name=$(shell basename $(CURDIR))

run: build exec clean

exec:
	./bin/${name}

build:
	CGO_ENABLED=0 go build -o bin/${name} -ldflags '-s -w -extldflags "-static"'

clean:
	rm -rf bin

deps:
	govendor init
	govendor add +e
	govendor update +v

dev:
	go get -u -v github.com/kardianos/govendor

