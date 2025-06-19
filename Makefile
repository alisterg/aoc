test:
	go test -v ./...

format:
	gofmt -w .

.PHONY: test
