BINARY_NAME=tcp-server

build:
	go build -o $(BINARY_NAME)
test:
	go test
clean:
	go clean
	rm -rf $(BINARY_NAME)
run:
	go build -o $(BINARY_NAME)
	./$(BINARY_NAME)
deps:
	go get github.com/stretchr/testify