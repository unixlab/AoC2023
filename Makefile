all: clean fmt test lint cover build
clean:
	rm -f AoC2023
fmt:
	find . -type f -name '*.go' | xargs -L 1 go fmt
test:
	go test -cover ./...
lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.55.2 golangci-lint run ./... -E revive -E govet -E misspell -E goimports -E gofmt -D errcheck --exclude-use-default=false -e "MixedCaps in package name; dayXX should be dayxx"
build:
	CGO_ENABLED=0 \
	go build \
	-o AoC2023
cover:
	go test -coverprofile=/tmp/AoC2023.cover.out ./...
	go tool cover -html=/tmp/AoC2023.cover.out
