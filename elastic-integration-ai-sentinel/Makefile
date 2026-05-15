ELASTIC_PACKAGE ?= elastic-package

.PHONY: lint format build test-pipeline test-asset test clean

lint:
	$(ELASTIC_PACKAGE) lint

format:
	$(ELASTIC_PACKAGE) format

build:
	$(ELASTIC_PACKAGE) build

test-pipeline:
	$(ELASTIC_PACKAGE) test pipeline

test-asset:
	$(ELASTIC_PACKAGE) test asset

test: lint build test-pipeline test-asset

clean:
	rm -rf build .elastic-package
