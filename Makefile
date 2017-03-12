TARGETS = linux-amd64 darwin-amd64 windows-386 windows-amd64
COMMAND_NAME = instantly
PACKAGE_NAME = github.com/ngenerio/$(COMMAND_NAME)
LDFLAGS = -ldflags=-X=main.version=$(VERSION)
OBJECTS = $(patsubst $(COMMAND_NAME)-windows-amd64%,$(COMMAND_NAME)-windows-amd64%.exe, $(patsubst $(COMMAND_NAME)-windows-386%,$(COMMAND_NAME)-windows-386%.exe, $(patsubst %,$(COMMAND_NAME)-%-v$(VERSION), $(TARGETS))))

fmt:
	@go fmt ./...

install:
	@glide install

clean: check-env
	rm -fr ./bin

release: clean check-env $(OBJECTS)

run:
	./instantly

$(OBJECTS): $(wildcard *.go)
	env GOOS=`echo $@ | cut -d'-' -f2` GOARCH=`echo $@ | cut -d'-' -f3 | cut -d'.' -f 1` go build -o ./bin/$@ $(LDFLAGS) $(PACKAGE_NAME)

check-env:
ifndef VERSION
	$(error VERSION is undefined)
endif

.PHONY: install clean release run
