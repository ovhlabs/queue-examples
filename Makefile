
build-go:
	cd go && \
	./package.sh

build-go-in-docker:
	cd go && \
	docker run --rm \
		-e GOPATH=/go -e GOBIN=/go/bin/ -e CGO_ENABLED=0 \
		-v $$(pwd):/go/src/github.com/runabove/queue-examples/go \
		-w /go/src/github.com/runabove/queue-examples/go \
			golang:1.7.3 \
				go build -o kafka-client

node-install-deps:
	cd nodejs && \
		npm install

python-install-deps:
	cd python && \
		pip install --upgrade kafka-python==1.3.1

## Build Docker images

build-docker: build-docker-go build-docker-nodejs build-docker-python

build-docker-go:
	cd go && \
		make build-go-in-docker && \
		docker build -t qaas/go-client-example .

build-docker-python:
	cd python && \
		docker build -t qaas/py-client-example .

build-docker-nodejs:
	cd nodejs && \
		docker build -t qaas/node-client-example .

# Run Docker containers

run-go-consumer:
	docker run --rm qaas/go-client-example \
		consume --broker ${HOST}:9093 \
			--username ${SASL_USERNAME} --password ${SASL_PASSWORD} \
			--topic ${TOPIC} --consumer-group ${USER}.go-g1

run-python-consumer:
	docker run --rm qaas/py-client-example \
		consume --broker ${HOST}:9093 \
			--username ${SASL_USERNAME} --password ${SASL_PASSWORD} \
			--topic ${TOPIC} --consumer-group ${USER}.py-g1

run-nodejs-consumer:
	docker run --rm qaas/node-client-example \
		consume --zk ${HOST}:2181 \
			--username ${SASL_USERNAME} --password ${SASL_PASSWORD} \
			--topic ${TOPIC} --consumer-group ${USER}.node-g1
