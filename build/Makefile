.PHONY: build

VERSION=1.0.0d
ENVIRONMENT=production
BINARY=admin
RELEASE=`date +%Y%m%d%H%M%S`
COMMIT=`git rev-parse --short HEAD`

build: armv7l aarch64 x86_64

armv7l:
	@mkdir -p ./dist/armv7l
	@rm -rf  ./dist/armv7l/*
	@env GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w -X main.version=v${VERSION} -X main.release=#0~${ENVIRONMENT}${RELEASE}-${COMMIT}" -o ./dist/armv7l/${BINARY} ../cmd/*.go
	@cp ./config.json ./dist/armv7l

aarch64:
	@mkdir -p ./dist/aarch64
	@rm -rf  ./dist/aarch64/*
	@env GOOS=linux GOARCH=arm64 go build -ldflags="-s -w -X main.version=v${VERSION} -X main.release=#0~${ENVIRONMENT}${RELEASE}-${COMMIT}" -o ./dist/aarch64/${BINARY} ../cmd/*.go
	@cp ./config.json ./dist/aarch64

x86_64:
	@mkdir -p ./dist/x86_64
	@rm -rf  ./dist/x86_64/*
	@env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X main.version=v${VERSION} -X main.release=#0~${ENVIRONMENT}${RELEASE}-${COMMIT}" -o ./dist/x86_64/${BINARY} ../cmd/*.go
	@cp ./config.json ./dist/x86_64
