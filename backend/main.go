package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/Isaiah-peter/lostandfound/api"
	db "github.com/Isaiah-peter/lostandfound/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres12:secret@localhost:5342/lostandfound?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("fail to connect to database: ", err.Error())
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("fail to start server: ", err.Error())
	}
}