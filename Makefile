APP=k8s-lab
GOLANGCI_LINT_VERSION=latest

.PHONY: build run test fmt fmt-check lint clean

build:
	go build -o bin/$(APP) ./cmd/k8s-lab

run:
	go run ./cmd/k8s-lab

test:
	go test ./...

fmt:
	gofmt -w .

fmt-check:
	test -z "$$(gofmt -l .)"

lint:
	golangci-lint run

clean:
	rm -rf bin