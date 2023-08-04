# go-simple-bank

D:\Gia\codes\go\go-simple-bank

docker run --rm -v "D:\Gia\codes\go\go-simple-bank:/src" -w /src kjconroy/sqlc generate
because sqlc not support Windows => user docker to run sqlc command

    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_USER: postgres
          POSTGRES_PASSWORD: admin
          POSTGRES_DB: simple_bank
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 5432:5432