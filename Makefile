build:
	@go build -o bin/userAuth

run: build
	@./bin/userAuth
