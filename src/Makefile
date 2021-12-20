.DEFAULT_GOAL := build

ifeq (src,$(notdir $(CURDIR)))
	BASEDIR = $(realpath ..)
else
	BASEDIR = $(CURDIR)
endif

ifeq ($(OS), Windows_NT)
	OUTPUT_EXT = .exe
else
	OUTPUT_EXT =
endif

SOURCE_FILE := main.go
OUTPUT_DIR := $(BASEDIR)/out
OUTPUT_FILE_BASE := shoppinglist

.PHONY: build lint vet run
lint:
	go fmt ./...

vet: lint
	go vet ./...

build: vet
#	@mkdir -p $(OUTPUT_DIR)
	go build -v -o $(OUTPUT_DIR)/$(OUTPUT_FILE_BASE)$(OUTPUT_EXT)

run: vet
	go run $(SOURCE_FILE)
