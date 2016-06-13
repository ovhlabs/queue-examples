all: build-go

build-go:
	cd golang && ./package.sh
