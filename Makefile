default: dev

pro:
	@go build -trimpath -ldflags '-w -s' -v -o ./bin/fiber-api .

dev:
	@go build -v -o ./bin/fiber-api .

doc:
	@swag init