ELASTIC_PACKAGE ?= elastic-package
ELASTIC_PACKAGE_STACK_ARGS ?=

.PHONY: lint format build test-pipeline test-asset stack-up stack-down test clean

lint:
	$(ELASTIC_PACKAGE) lint

format:
	$(ELASTIC_PACKAGE) format

build:
	$(ELASTIC_PACKAGE) build

test-pipeline:
	$(ELASTIC_PACKAGE) stack up -d --services=elasticsearch
	$(ELASTIC_PACKAGE) test pipeline
	$(ELASTIC_PACKAGE) stack down

test-asset:
	$(ELASTIC_PACKAGE) test asset

stack-up:
	$(ELASTIC_PACKAGE) stack up $(ELASTIC_PACKAGE_STACK_ARGS)

stack-down:
	$(ELASTIC_PACKAGE) stack down

test: lint build test-pipeline test-asset

clean:
	rm -rf build .elastic-package
