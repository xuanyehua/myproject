
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker


build:

	go build -o user_web-web main.go plugin.go

test:
	go test -v ./... -cover

docker:
	docker build . -t user_web-web:latest
