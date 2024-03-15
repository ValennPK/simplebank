postgres:
	docker run --name postgres16 -p 8080:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=asd123 -e POSTGRES_DB=mydatabase -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres16 dropdb simple_bank

migrateup:
	migrate -path  db/migration -database "postgresql://root:asd123@localhost:8080/simple_bank?sslmode=disable" -verbose up 

migratedown:
	migrate -path  db/migration -database "postgresql://root:asd123@localhost:8080/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...


server:
	go run main.go

.PHONY: postgres  createdb dropdb migrateup migratedown sqlc test server