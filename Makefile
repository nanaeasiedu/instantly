fmt:
	@go fmt ./...

install:
	@glide installl

build:
	@go build -o instantly

run:
	./instantly

.PHONY: install build run
