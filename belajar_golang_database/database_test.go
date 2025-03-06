package belajargolangdatabase

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestDatabase(t *testing.T) {

}

func TestDatabaseConnection(t *testing.T) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=CoRei7%4dKontol123!@#*?*?*?321 dbname=testit sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	t.Log("Successfully connected to database!")
}
