run: build
	@./bin/swiftclip-backend

build:
	@go build -o bin/swiftclip-backend cmd/api/main.go
