package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/quocgiahcmut/simple-bank/api"
	db "github.com/quocgiahcmut/simple-bank/db/sqlc"
	"github.com/quocgiahcmut/simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cant load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cant connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannt create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cant start server: ", err)
	}
}
