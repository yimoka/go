GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
CONFIG_PROTO_FILES=$(shell find config -name *.proto)

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
				--proto_path=./third_party \
				--go_out=paths=source_relative:. \
				${CONFIG_PROTO_FILES}