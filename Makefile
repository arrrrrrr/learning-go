.DEFAULT_GOAL := bin
PLATFORM=local

.PHONY: bin build_base build_run run
build_base:
	@docker build . \
		--tag shoppinglist-base:latest \
		--target base \
		--platform ${PLATFORM}

build_run: build_base
	@docker build . \
		--tag shoppinglist-run:latest \
		--target run \
		--platform ${PLATFORM}

run: build_run
	@docker run \
		--rm \
		-it \
		--name "shoppinglist" \
		shoppinglist-run:latest \
		"make run"

bin: build_base
	docker build . \
		--tag shoppinglist-bin:latest \
		--target bin \
		--output out \
		--platform ${PLATFORM}
