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
mock:
	mockgen -package mockdb -destination db\mock\store.go github.com/quocgiahcmut/simple-bank/db/sqlc Store
remove-pb:
	rm .\pb\*.go
proto-gen:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto
.PHONY: sqlc