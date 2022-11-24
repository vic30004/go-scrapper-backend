.DEFAULT_GOAL :=build

fmt:
	go fmt ./..

vet:
	go vet ./..

build: fmt
	go build github.com/vic30004/go-scrapper-backend

start:
	go run main.go serve