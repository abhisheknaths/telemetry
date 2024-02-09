SOURCES := $(wildcard *.go)
VERSION=$(shell git describe --tags --long --dirty 2>/dev/null)

## we must have tagged the repo at least for the version to work
ifeq ($(VERSION),)
	VERSION = UNKNOWN
endif

build: $(sources) build/Dockerfile
	docker build -t telemetry:latest . -f build/Dockerfile --build-arg VERSION=$(VERSION)