// package db

// import (
// 	"database/sql"
// 	"log"
// 	"os"
// 	"testing"

// 	_ "github.com/lib/pq"
// )

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://postgres_nobel:postgres_nobel@localhost:5432/simple_bank?sslmode=disable"
// )

// // var testQueries *Queries
// // // var testDB *sql.DB

// // func TestMain(m *testing.M) {
// // 	conn, err := sql.Open(dbDriver, dbSource)
// // 	if err != nil {
// // 		log.Fatal("cannot connect to the DB:", err)
// // 	}
// // 	// code := m.Run()
// // 	// // testDB.Close()
// // 	testQueries = New(conn)
// // 	os.Exit(m.Run())

// // }

package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const dbSource = "postgresql://postgres_nobel:postgres_nobel@localhost:5432/simple_bank?sslmode=disable"

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error

	// Use context to establish connection
	testDB, err = pgxpool.New(context.Background(), dbSource)
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
