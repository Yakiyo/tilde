alias b := build
alias u := update
alias r := run

default:
    @just --list

update:
    go mod tidy
    go get

build:
    go build -o tilde main.go

@run *arg:
    go run main.go $@

fmt:
	go fmt ./...