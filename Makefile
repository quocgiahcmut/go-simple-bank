sqlc:
	docker run --rm -v "D:\Gia\codes\go\go-simple-bank:/src" -w /src kjconroy/sqlc generate

.PHONY: sqlc