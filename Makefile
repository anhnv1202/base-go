APP_NAME= server

run:
	go run ./cmd/$(APP_NAME)

build:
	go build -o ./cmd/$(APP_NAME) main.go
