all: build-go

build-go:
	cd golang && ./package.sh

build-node:
	npm install

build-python:
	pip install kafka-python
