SHELL := /bin/bash

.DEFAULT_GOAL := help

.PHONY: build
install:
	python3 -m venv .venv
	( \
		source .venv/bin/activate; \
		pip install -r requirements.txt; \
	)

.PHONY: clean
clean: 
	rm -rf ./.venv 

.PHONY: help
help:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
