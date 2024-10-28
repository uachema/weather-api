build:
	@go build -o ./bin/weather-api ./cmd

run: build
	@ ./bin/weather-api

test:
	@go test -v ./...
