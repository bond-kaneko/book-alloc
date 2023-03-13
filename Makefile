.PHONY: deps gen lint error build dev test

GOOS = linux
GOARCH = amd64
GOPATH = ${shell go env GOPATH}
COMMIT_HASH = ${shell git rev-parse HEAD}

error:
	exit 1

gen:
	go generate ./...

dev:
	bash ./scripts/initdb.sh
	mkdir -p ./tmp
	air -c ./cmd/book-alloc/.air.toml
