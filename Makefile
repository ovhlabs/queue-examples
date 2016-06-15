
build-go:
	cd golang && \
		./package.sh

node-install-deps:
	cd nodejs && \
		npm install

python-install-deps:
	cd python && \
		pip install kafka-python

## Build Docker images

build-docker: build-docker-golang build-docker-nodejs build-docker-python

build-docker-golang:
	cd golang && \
		make build-go-in-docker && \
		docker build -t qaas/golang-client-example .

build-docker-nodejs:
	cd nodejs && \
		docker build -t qaas/nodejs-client-example .

build-docker-python:
	cd python && \
		docker build -t qaas/python-client-example .

# Run Docker containers

run-golang-consumer:
	docker run --rm qaas/golang-client-example \
		consume --kafka ${HOST}:9092 --key ${KEY} --topic ${PREFIX}.${TOPIC} --group ${PREFIX}.golang-${GROUP}

run-nodejs-consumer:
	docker run --rm qaas/nodejs-client-example \
		consume --zk ${HOST}:2181 --key ${KEY} --topic ${PREFIX}.${TOPIC} --group ${PREFIX}.nodejs-${GROUP}

run-python-consumer:
	docker run --rm qaas/python-client-example \
		consume --kafka ${HOST}:9092 --key ${KEY} --topic ${PREFIX}.${TOPIC} --group ${PREFIX}.python-${GROUP}
