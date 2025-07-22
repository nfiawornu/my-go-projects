package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nfiawornu/my-go-projects/simple_bank/api"
	db "github.com/nfiawornu/my-go-projects/simple_bank/db/sqlc"
)

const (
	serverAddress = "0.0.0.0:8082"
	dbSource      = "postgresql://postgres_nobel:postgres_nobel@localhost:5432/simple_bank?sslmode=disable"
)

// var testQueries *Queries
var testDB *pgxpool.Pool

func main() {

	// Use context to establish connection
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to the DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}

}
