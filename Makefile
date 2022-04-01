all: test

test:
	go test -covermode atomic -v ./...
