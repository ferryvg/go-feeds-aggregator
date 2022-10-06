# 'help' target by default
.PHONY: default
default: help

# enable silent mode
ifndef VERBOSE
.SILENT:
endif

NAME="go-feeds-aggregator"
USER_ID=$(shell id -u)
GROUP_ID=$(shell id -g)

## Compile daemon into dist/
compile:
	mkdir -p ./dist
	rm -f ./dist/${NAME}d
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"  -a -o ./dist/${NAME}d ./cmd/aggregator/main.go

## Run tests
test:
	go test -tags tests -v ./...