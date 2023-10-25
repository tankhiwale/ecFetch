

build:
	@go build -o bin/fetch

run: build
	@ ./bin/fetch
