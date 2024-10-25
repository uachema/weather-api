build:
	@go build -o ./bin/weather-api

run: build
	@PORT=3000 ./bin/weather-api

test:
	@go test -v ./...
