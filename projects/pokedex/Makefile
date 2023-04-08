
run:
	go run server.go

local-env:
	docker-compose down -v
	docker-compose up -d

install:
	docker-compose up

unistall:
	docker-compose down -v

openapi-generate:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0