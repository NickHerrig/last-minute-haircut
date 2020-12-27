SHELL := /bin/bash

.PHONY: default
default: run

.PHONY: deps 
deps:
	python3 -m venv venv/
	./venv/bin/pip install -Uq wheel
	./venv/bin/pip install -Uq -r requirements.txt

.PHONY: run 
run: deps 
	./venv/bin/python3 haircut.py 

.PHONY: clean
clean: 
	rm -rf ./venv 

.PHONY: help
help:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
