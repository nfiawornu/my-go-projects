package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nfiawornu/my-go-projects/simple_bank/util"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Use context to establish connection
	testDB, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the DB:", err)
	}

	// Pass pgxpool.Pool (implements DBTX) to sqlc's New()
	testQueries = New(testDB)

	// Run all tests
	code := m.Run()

	// Cleanup
	testDB.Close()
	os.Exit(code)
}
