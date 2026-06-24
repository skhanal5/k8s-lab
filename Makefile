APP=k8s-lab

.PHONY: build run test fmt lint clean

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