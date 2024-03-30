BINARY_NAME=myapp

build:
	go build -o ${BINARY_NAME}

clean-build:
	go clean

run: build
ifneq ("$(WORDS)", "")
	./${BINARY_NAME} -s "$(WORDS)"
endif

compile:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows

clean-compile:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows