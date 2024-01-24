up:
	docker-compose up -d --build

down:
	docker-compose down --remove-orphans --volumes

migrate-status:
	goose -dir ./database/migrations postgres "user=postgres password=changeme host=0.0.0.0 dbname=gortener sslmode=disable" status

migrate-up:
	goose -dir ./database/migrations postgres "user=postgres password=changeme host=0.0.0.0 dbname=gortener sslmode=disable" up

database-connect:
	docker exec -it database psql -U postgres

database-dump:
	docker exec -it database pg_dump -U postgres  --schema-only gortener > database/schema.sql

database-generate:
	sqlc generate

check-dependencies:
	go mod tidy

start:
	go run main.go

start-server:
	go run main.go server