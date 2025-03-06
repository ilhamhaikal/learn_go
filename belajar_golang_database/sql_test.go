package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

/*
alasan menggunakan execcontext adalah untuk menghandle context yang lebih baik
agar tidak terjadi memory leak setiap connection yang dibuat menggunakan defear
harus di close
jika tidak di close maka akan terjadi memory leak
jika terjadi memory leak maka akan terjadi penumpukan memory yang tidak terpakai
*/

func TestSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES('Ilham', 'Ilham')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new user")
}

func TestSQLQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", id, "Name:", name)
	}
}

/*
dalam query complex kita menggunakan sql.NullString, sql.NullTime, dan bool
karena kita tidak tahu apakah data tersebut ada atau tidak
jika data tersebut tidak ada maka akan mengembalikan nilai default dan data tidak akan di tampilkan
jika data tersebut ada maka akan menampilkan data tersebut
*/
func TestQueryComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, created_at, birth_date, married FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &createdAt, &birthDate, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Created At:", createdAt)
		if birthDate.Valid {
			fmt.Println("Birth Date:", birthDate)
		}
		fmt.Println("Married:", married)
		fmt.Println("------------------------") // Adding a separator between records
	}
}
