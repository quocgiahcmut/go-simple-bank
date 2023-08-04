sqlc-lap:
	docker run --rm -v "D:\Gia\codes\go\go-simple-bank:/src" -w /src kjconroy/sqlc generate
sqlc-pc:
	docker run --rm -v "D:\code\go\go-simple-bank:/src" -w /src kjconroy/sqlc generate
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it postgres dropdb simple_bank
migrateup: 
	migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/simple_bank?sslmode=disable" -verbose down
test: 
	go test -v -cover ./...
server:
	go run main.go
.PHONY: sqlc