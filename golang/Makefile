
build-go-in-docker:
	docker run --rm \
		-e GOPATH=/go \
		-e GOBIN=/go/bin/ \
		-e CGO_ENABLED=0 \
		-v $$(pwd):/go/src/github.com/runabove/queue-examples/golang \
		-w /go/src/github.com/runabove/queue-examples/golang \
			golang:1.6.2 \
				go build -o qaas-client