.PHONY: build run \
	build_docker

build:
	@go build -ldflags="-s -w" -o rest-api main.go

run:
	@go run main.go

build_docker:
	@docker build -t xm-rest-api:latest -f Dockerfile .
