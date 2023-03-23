package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// const (
// 	host = "localhost"
// 	port = 5432
// 	user = "postgres"
// 	password = "Dm2016dM"
// 	dbname = "first_db"
// )

type Student struct {
	Name string
	LastName string
	Empid int
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// ? Open the connection
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=Dm2016dM dbname=first_db sslmode=disable")
	CheckError(err)
	defer db.Close()

	// script := `INSERT INTO "Employee" ("Name", "LastName", "Empid") VALUES ('Denis', 'Matsenko', 1)`
	// script := `DELETE FROM "Employee" WHERE "Empid" = 1`
	// script := `UPDATE "Employee" SET "Name" = 'Michal', "LastName" = 'Vanis' WHERE "Empid" = 1`
	script := `SELECT "Empid", "Name", "LastName" FROM "Employee"`
	
	// ? Execute the script
	rows, err := db.Query(script)
	CheckError(err)

	// ? Iterate over the rows
	for rows.Next() {
		var student Student

		// ?  Scan the row values into variables
		err = rows.Scan(&student.Empid, &student.LastName, &student.Name)
		CheckError(err)

		// ? Print the row values
		fmt.Println(student.Empid, student.Name)
	}
	rows.Close()
	
	fmt.Println("Done")
}