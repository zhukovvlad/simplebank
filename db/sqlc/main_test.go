// Package db provides functions for interacting with the database.
package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/zhukovvlad/simplebank/util"
)

// testQueries holds an instance of the queries to be used in tests.
var testQueries *Queries
var testDB *sql.DB

// TestMain initializes the database for testing and runs all tests in the package.
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Opens a connection to the database.
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Creates a new instance of queries for testing.
	testQueries = New(testDB)

	// Runs all tests in the package.
	os.Exit(m.Run())
}