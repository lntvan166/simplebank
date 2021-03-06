package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/lntvan166/simplebank/api"
	db "github.com/lntvan166/simplebank/db/sqlc"
	"github.com/lntvan166/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load config", err)
	}

	conn, err := sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.SERVER_ADDRESS)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
