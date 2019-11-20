SHELL := /bin/bash


help:
	@echo "- build"
	@echo "- infra-up"
	@echo "- infra-down"
	@echo "- test"

build:
	./scripts/build.sh

test:
	./scripts/test.sh


infra-up:
	./scripts/infra.sh up

infra-down:
	./scripts/infra.sh down