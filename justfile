default:
    echo 'Hello, world!'

update:
    go mod tidy
    go get

build:
    go build -o tilde main.go

@run *arg:
    go run main.go $@

fmt:
	go fmt ./...