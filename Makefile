sqlc:
	docker run --rm -v "D:\code\go\go-simple-bank:/src" -w /src kjconroy/sqlc generate
migrateup: 
	migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/simple_bank?sslmode=disable" -verbose down
.PHONY: sqlc