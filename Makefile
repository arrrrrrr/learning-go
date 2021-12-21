# Include the go makefile
include go.Makefile

PLATFORM = local
DOCKER_DIR = $(BASEDIR)/docker

.PHONY: build_base build_run run-docker build-docker

build_base:
	@docker build . \
		-f $(DOCKER_DIR)/Dockerfile \
		--tag shoppinglist-base:latest \
		--target base \
		--platform ${PLATFORM}

build_run:
	@docker build . \
		-f $(DOCKER_DIR)/Dockerfile \
		--tag shoppinglist-run:latest \
		--target run \
		--platform ${PLATFORM}

run-docker: build_run
	@docker run \
		--rm \
		-it \
		--name "shoppinglist" \
		shoppinglist-run \
		"make run TARGET=shoppinglist"

build-docker: build_base
	@docker build . \
		-f $(DOCKER_DIR)/Dockerfile \
		--tag shoppinglist-bin:latest \
		--target bin \
		--output out \
		--platform ${PLATFORM}

