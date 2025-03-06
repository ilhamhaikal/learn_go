package belajargolangdatabase

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

/*
untuk mengamankan dari sql injection kita bisa menggunakan parameterized query
parameterized query adalah query yang menggunakan parameter untuk menghindari sql injection
format query yang benar adalah menggunakan $1, $2, $3 untuk postgresql
format query yang benar adalah menggunakan ?, ?, ? untuk mysql

bisa menggunakan queryContext atau execContext
queryContext digunakan untuk query yang mengembalikan data
execContext digunakan untuk query yang tidak mengembalikan data
*/
func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin' ; #"
	password := "ilham123"

	// Test login
	query := `SELECT username FROM "user" WHERE username = $1 AND password = $2 LIMIT 1`
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully logged in as:", username)
	} else {
		fmt.Println("Login failed")
	}
}

func TestSqlInjection2(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin' ; #"
	password := "admin123"

	// Test login
	query := `INSERT INTO "user" (username, password) VALUES ($1, $2)`
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		fmt.Println("successfully login:", username)
	} else {
		fmt.Println("error login", err)
	}
}

// auto increment pada golang menggunakan RETURNING id untuk database postgresql
func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	email := "ilham@test.com"
	comment := "ini adalah comment"

	// Test login
	query := `INSERT INTO comment (email, comment) VALUES ($1, $2) RETURNING id`
	var insertId int64
	err := db.QueryRowContext(ctx, query, email, comment).Scan(&insertId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Last inserted ID: %d\n", insertId)
}

// prepare statement digunakan untuk mengeksekusi query yang sama berkali-kali
func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comment (email, comment) VALUES ($1, $2) RETURNING id"
	statment, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer statment.Close()

	for i := 0; i < 10; i++ {
		email := "ilham" + strconv.Itoa(1) + "@test.com"
		comment := "ini adalah comment"

		var id int64
		err = statment.QueryRowContext(ctx, email, comment).Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Last inserted ID: %d\n", id)
	}
}
