default: dev

pro:
	@go build -trimpath -ldflags '-w -s' -v -o ./bin/fiber-api .

dev: doc
	@go build -v -o ./bin/fiber-api .

doc:
	@swag init