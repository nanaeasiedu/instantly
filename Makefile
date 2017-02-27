fmt:
	@go fmt ./...

install:
	@godep save ./...

run:
	./instantly

.PHONY: install run