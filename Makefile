build:
	@go build -o ./bin/weather-api
run: build
	@./bin/weather-api
test:
	@go test -v ./...