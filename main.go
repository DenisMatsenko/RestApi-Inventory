package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "Dm2016dM"
	dbname = "first_db"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=Dm2016dM dbname=first_db sslmode=disable")
	CheckError(err)
	defer db.Close()

	// script := `INSERT INTO "Employee" ("Name", "LastName", "Empid") VALUES ('Denis', 'Matsenko', 1)`
	// script := `DELETE FROM "Employee" WHERE "Empid" = 1`
	// script := `UPDATE "Employee" SET "Name" = 'Michal', "LastName" = 'Vanis' WHERE "Empid" = 1`
	script := `SELECT "Name" AS name, "LastName" AS lastName, "Empid" AS id FROM "Employee"`
	rows, err := db.Query(script)
	CheckError(err)

	for rows.Next() {
		var Name string
		var LastName string
		var Empid int

		// Scan the row values into variables
		err = rows.Scan(&LastName, &Name, &Empid)
		CheckError(err)

		// Print the row values
		fmt.Println(Name, LastName, Empid)
	}

	fmt.Println(err)
	CheckError(err)

	fmt.Println("Successfull")
}