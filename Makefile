GO_VERSION=1.17

deps:
	go mod tidy -compat=$(GO_VERSION)

test:
	go test ./...

