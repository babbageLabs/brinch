BINARY_NAME=main.exe

build:
	@go build -o ${BINARY_NAME} main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test -v ./...

coverage:
	go test -v ./... -cover