.PHONY: build

VERSION=1.0.0d
ENVIRONMENT=develop
BINARY=admin
RELEASE=`date +%Y%m%d%H%M%S`
COMMIT=`git rev-parse --short HEAD`

build: armv7l aarch64 x86_64

armv7l:
	@mkdir -p ./release/armv7l
	@rm -rf  ./release/armv7l/*
	@env GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w -X main.version=v${VERSION} -X main.release=#0~${ENVIRONMENT}${RELEASE}-${COMMIT}" -o ./release/armv7l/${BINARY} ../cmd/*.go
	@cp ./config.json ./release/armv7l

aarch64:
	@mkdir -p ./release/aarch64
	@rm -rf  ./release/aarch64/*
	@env GOOS=linux GOARCH=arm64 go build -ldflags="-s -w -X main.version=v${VERSION} -X main.release=#0~${ENVIRONMENT}${RELEASE}-${COMMIT}" -o ./release/aarch64/${BINARY} ../cmd/*.go
	@cp ./config.json ./release/aarch64

x86_64:
	@mkdir -p ./release/x86_64
	@rm -rf  ./release/x86_64/*
	@env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X main.version=v${VERSION} -X main.release=#0~${ENVIRONMENT}${RELEASE}-${COMMIT}" -o ./release/x86_64/${BINARY} ../cmd/*.go
	@cp ./config.json ./release/x86_64
