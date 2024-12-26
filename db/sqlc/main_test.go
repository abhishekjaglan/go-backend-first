package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) { // TestMain by convention is the entry point of all tests inside a package
	var err error
	testDB, err = sql.Open(dbDriver, dbSource) // connects to the database (sql.DB connection)
	if err != nil {
		log.Fatal("cannot connect tp db:", err)
	}

	testQueries = New(testDB) // returns DBTX that holds all the transaction/queries to run

	os.Exit(m.Run()) // runs all the tests
}
