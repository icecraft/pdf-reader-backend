all: clean test build image

check:
	go mod tidy
	go fmt ./...
	go vet ./...

test: check
	go test -p 1 -gcflags=-l ./...

coverage: check
	go-acc -o coverage.cov ./... -- -timeout 900s -race 
	go tool cover -func=coverage.cov

clean:
	rm -rf build

.PHONY: check clean test
