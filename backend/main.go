package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Isaiah-peter/lostandfound/api"
	db "github.com/Isaiah-peter/lostandfound/db/sqlc"
	"github.com/Isaiah-peter/lostandfound/util"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres12:secret@localhost:5342/lostandfound?sslmode=disable"
	address = "localhost:8080"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config", err)
	}

	fmt.Println(config)
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("fail to connect to database: ", err.Error())
	}
	store := db.NewStore(conn)

	server, err := api.NewServer(store)

	if err != nil {
		fmt.Println("cannot create server", err.Error())
	}

	err = server.Start(address)

	if err != nil {
		log.Fatal("fail to start server: ", err.Error())
	}
}