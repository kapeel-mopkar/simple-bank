postgres:
	docker run --name pg12 -p 5454:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it pg12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it pg12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5454/simple_bank?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5454/simple_bank?sslmode=disable" -verbose down

sqlc-gen:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown