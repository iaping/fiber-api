default: dev

pro:
	@go build -trimpath -ldflags '-w -s' -v -o ./bin/fiber-api .

dev: docs
	@go build -v -o ./bin/fiber-api .

docs:
	@swag init