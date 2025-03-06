package belajargolangdatabase

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// cara membuat connection pooling
func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=CoRei7%4dKontol123!@#*?*?*?321 dbname=testit sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
