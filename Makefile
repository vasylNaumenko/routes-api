.SILENT:
.EXPORT_ALL_VARIABLES:
.PHONY: all test help lint build run run-build run-compose clean rebuild

name = routes-api
all: run

help:
	bash help/help.sh

lint:
	golangci-lint run

test:
	go clean -testcache ./...
	go test ./...

build: test
	go build -o ./cmd/${name} ./cmd/.

run:
	cd cmd; go run .

run-build: build
	cd cmd/{name}; ./${name}

run-compose:
	docker-compose up

clean:
	go clean ./...
	rm -f ./cmd/${name}

rebuild: clean build


