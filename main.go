package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/lntvan166/simplebank/api"
	db "github.com/lntvan166/simplebank/db/sqlc"
)

const (
	dbDriver     = "postgres"
	dbSource     = "postgresql://postgres:handsome2022@localhost:2345/simple_bank?sslmode=disable"
	sererAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(sererAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
