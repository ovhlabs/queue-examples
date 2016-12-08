
package:
	./package.sh

build:
	go build -o kafka-client

consume:
	./kafka-client consume \
		--broker ${HOST}:9093 \
		--username ${SASL_USERNAME} --password ${SASL_PASSWORD} \
		--topic ${TOPIC} --consumer-group ${SASL_USERNAME}.group-go

produce:
	./kafka-client produce \
		--broker ${HOST}:9093 \
		--username ${SASL_USERNAME} --password ${SASL_PASSWORD} \
		--topic ${TOPIC}

##

build-binary-in-docker:
	docker run --rm \
		-e GOPATH=/go -e GOBIN=/go/bin/ -e CGO_ENABLED=0 \
		-v $$(pwd):/go/src/github.com/runabove/queue-examples/go \
		-w /go/src/github.com/runabove/queue-examples/go \
			golang:1.7.3 \
				go build -o kafka-client

build-docker: build-binary-in-docker
	docker build -t ovhqueue/go-client-example .

consume-in-docker:
	docker run --rm -ti ovhqueue/go-client-example \
		consume --broker ${HOST}:9093 \
			--username ${SASL_USERNAME} --password ${SASL_PASSWORD} \
			--topic ${TOPIC} --consumer-group ${SASL_USERNAME}.group-dgo
