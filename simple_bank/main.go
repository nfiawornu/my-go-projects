package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nfiawornu/my-go-projects/simple_bank/api"
	db "github.com/nfiawornu/my-go-projects/simple_bank/db/sqlc"
	"github.com/nfiawornu/my-go-projects/simple_bank/util"
)

// var testQueries *Queries
var testDB *pgxpool.Pool

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Use context to establish connection
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}

}
