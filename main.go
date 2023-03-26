package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	// "github.com/rs/cors"
)

type item struct {
	Name 			string 		`json:"name"`
	Description 	string 		`json:"description"`
	Quality 		int 		`json:"quality"`
	Quantity 		int 		`json:"quantity"`
	ItemType 		string 		`json:"itemType"`
	Price 			int 		`json:"price"`
	Id 				string 		`json:"id"`
	Color 			string 		`json:"color"`
	Img 			string 		`json:"img"`
}

const dbconnection = "host=localhost port=5432 user=postgres password=Dm2016dM dbname=EShopInventory sslmode=disable"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// func main() {
// 	// ? Open the connection
// 	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=Dm2016dM dbname=first_db sslmode=disable")
// 	CheckError(err)
// 	defer db.Close()

// 	// script := `INSERT INTO "Employee" ("Name", "LastName", "Empid") VALUES ('Denis', 'Matsenko', 1)`
// 	// script := `DELETE FROM "Employee" WHERE "Empid" = 1`
// 	// script := `UPDATE "Employee" SET "Name" = 'Michal', "LastName" = 'Vanis' WHERE "Empid" = 1`
// 	script := `SELECT "Empid", "Name", "LastName" FROM "Employee"`
	
// 	// ? Execute the script
// 	rows, err := db.Query(script)
// 	CheckError(err)

// 	// ? Iterate over the rows
// 	for rows.Next() {
// 		var student Student

// 		// ?  Scan the row values into variables
// 		err = rows.Scan(&student.Empid, &student.LastName, &student.Name)
// 		CheckError(err)

// 		// ? Print the row values
// 		fmt.Println(student.Empid, student.Name)
// 	}
// 	rows.Close()
	
// 	fmt.Println("Done")
// }

func addItem(w http.ResponseWriter, r *http.Request) {
	// ? Decode the request body into a new `Item` instance
	var item item
	json.NewDecoder(r.Body).Decode(&item)

	// ? Open the connection to database
	db, err := sql.Open("postgres", dbconnection)
	CheckError(err)
	defer db.Close()

	// ? Create and execute the SQL script
	script := fmt.Sprintf(`INSERT INTO "Inventory" ("Name", "Description", "Quality", "Quantity", "ItemType", "Price", "Color", "Img") VALUES ('%s', '%s', '%d', '%d', '%s', '%d', '%s', '%s')`, item.Name, item.Description, item.Quality, item.Quantity, item.ItemType, item.Price, item.Color, item.Img)
	_, sqlerr := db.Exec(script)
	CheckError(sqlerr)

	// ? Return
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	var id string = mux.Vars(r)["id"]

	// ? Open the connection to database
	db, err := sql.Open("postgres", dbconnection)
	CheckError(err)
	defer db.Close()

	// ? Create and execute the SQL script
	script := fmt.Sprintf(`DELETE FROM "Inventory" WHERE "Id" = '%s'`, id)

	// ? Execute the script
	_, sqlerr := db.Exec(script)
	CheckError(sqlerr)

	// ? Return
	// ! header key and value
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func deleteAllItems(w http.ResponseWriter, r *http.Request) {
	// ? Open the connection to database
	db, err := sql.Open("postgres", dbconnection)
	CheckError(err)
	defer db.Close()

	// ? Create and execute the SQL script
	script := `TRUNCATE TABLE "Inventory"`
	_, sqlerr := db.Exec(script)
	CheckError(sqlerr)

	// ? Return
	// ! header key and value
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// func addItemToDB(item Item) {
// 	fmt.Println("Adding item to DB")
// 	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=Dm2016dM dbname=EShopInventory sslmode=disable")
// 	CheckError(err)
// 	defer db.Close()

// 	script := `INSERT INTO "Inventory" ("Name", "Description", "Quality", "Quantity", "ItemType", "Price", "ID", "Color", "Img") VALUES ('` + item.name + `', '` + item.description + `', ` + string(item.quality) + `, ` + string(item.quantity) + `, '` + item.itemType + `', ` + string(item.price) + `, '` + item.id + `', '` + item.color + `', '` + item.img + `')`

// 	db.Exec(script)
// }

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/additem", addItem).Methods("POST")
	mux.HandleFunc("/deleteall", deleteAllItems).Methods("DELETE")
	mux.HandleFunc("/delete/id={id}", deleteItem).Methods("DELETE")
	fmt.Println("Server running on port 8080")
	http.ListenAndServe("localhost:8080", mux)
}