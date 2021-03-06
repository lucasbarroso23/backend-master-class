package main

import (
	"database/sql"
	"log"

	"github.com/lucasbarroso23/backend-master-class/api"
	db "github.com/lucasbarroso23/backend-master-class/db/sqlc"
	"github.com/lucasbarroso23/backend-master-class/util"

	_ "github.com/lib/pq" // postgres go driver
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
