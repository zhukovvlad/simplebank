// Package db provides functions for interacting with the database.
package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	// dbDriver defines the database driver to be used.
	dbDriver = "postgres"

	// dbSource defines the connection string for the database.
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

// testQueries holds an instance of the queries to be used in tests.
var testQueries *Queries
var testDB *sql.DB

// TestMain initializes the database for testing and runs all tests in the package.
func TestMain(m *testing.M) {
	var err error
	// Opens a connection to the database.
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Creates a new instance of queries for testing.
	testQueries = New(testDB)

	// Runs all tests in the package.
	os.Exit(m.Run())
}