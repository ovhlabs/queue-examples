#!/bin/sh
OS="darwin linux windows"
ARCH="amd64"

for GOOS in $OS; do
    for GOARCH in $ARCH; do
        architecture="${GOOS}-${GOARCH}"
        echo "Building ${architecture}"
        export GOOS=$GOOS
        export GOARCH=$GOARCH
        go get
        go build -o bin/qaas-client-${architecture}
    done
done
