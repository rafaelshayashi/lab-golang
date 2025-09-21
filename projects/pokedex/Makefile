run:
	go run server.go

up:
	docker-compose up -d --build

down:
	docker-compose down --remove-orphans --volumes

database-connect:
	docker exec -it pokedex-database psql -U postgres

openapi-generate:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0