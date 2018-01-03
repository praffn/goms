build-cli:
	GOOS=linux GOARCH=amd64 go build

build-docker:
	docker build -t consignment-cli .

build: build-cli build-docker

run:
	docker run -e MICRO_REGISTRY=mdns consignment-cli