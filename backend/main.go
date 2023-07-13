package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Isaiah-peter/lostandfound/api"
	db "github.com/Isaiah-peter/lostandfound/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres12:secret@localhost:5342/lostandfound?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("fail to connect to database: ", err.Error())
	}
	store := db.NewStore(conn)

	server, err := api.NewServer(store)

	if err != nil {
		fmt.Println("cannot create server", err.Error())
	}

	err = server.Start()

	if err != nil {
		log.Fatal("fail to start server: ", err.Error())
	}
}