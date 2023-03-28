package main

import (
	// "bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	// "io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	// "github.com/rs/cors"
)

type item struct {
	Name 			string 		`json:"name"`
	Manufacturer	string 		`json:"manufacturer"`
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

// # Add item
func addItem(w http.ResponseWriter, r *http.Request) {
	// ? Decode the request body for isArray check
	respBody, _ := ioutil.ReadAll(r.Body)
	var body []map[string]any

	json.Unmarshal(respBody, &body)
	fmt.Printf("%v", body)



	// x := bytes.TrimLeft(r.Body, " \t\r\n")

	// isArray := len(x) > 0 && x[0] == '['
	// isObject := len(x) > 0 && x[0] == '{'


	// // ? Open the connection to database
	// db, err := sql.Open("postgres", dbconnection)
	// CheckError(err)
	// defer db.Close()

	// var script string

	// fmt.Println("everything is ok")
	
	// if len(body) == 0 {
	// 	// * if it is only one object  

	// 	// ? Decode the request body into a new `Item` instance
	// 	var item item
	// 	err := json.NewDecoder(r.Body).Decode(&item)
	// 	CheckError(err)
	// 	// ? Create and execute the SQL script
	// 	script = fmt.Sprintf(`INSERT INTO "Inventory" ("Name", "Manufacturer", "Description", "Quality", "Quantity", "ItemType", "Price", "Color", "Img") VALUES ('%s', '%s', '%s', '%d', '%d', '%s', '%d', '%s', '%s')`, item.Name, item.Manufacturer, item.Description, item.Quality, item.Quantity, item.ItemType, item.Price, item.Color, item.Img)
	// 	_, sqlerr := db.Exec(script)
	// 	CheckError(sqlerr)

	// } else  {
	// 	// * if it is an array of objects
	// 	CheckError(err)

	// 	for _, item := range itemArr {
	// 		// ? Create and execute the SQL script
	// 		script = fmt.Sprintf(`INSERT INTO "Inventory" ("Name", "Manufacturer", "Description", "Quality", "Quantity", "ItemType", "Price", "Color", "Img") VALUES ('%s', '%s', '%s', '%d', '%d', '%s', '%d', '%s', '%s')`, item.Name, item.Manufacturer, item.Description, item.Quality, item.Quantity, item.ItemType, item.Price, item.Color, item.Img)
	// 		_, sqlerr := db.Exec(script)
	// 		CheckError(sqlerr)
	// 	}
	// }


	// ? Return
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// # Delete item by id
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

// # Delete all items
func deleteAllItems(w http.ResponseWriter, r *http.Request) {
	if mux.Vars(r)["id"] == "" {
		println("No id provided")
	}

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

func test(w http.ResponseWriter, r *http.Request) {
	// ? Decode the request body into a new `Item` instance
	var items []item
	json.NewDecoder(r.Body).Decode(&items)

	println(items[0].Name)
	println(items[1].Name)
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
	mux.HandleFunc("/test", test).Methods("POST")
	mux.HandleFunc("/additem", addItem).Methods("POST")
	mux.HandleFunc("/deleteall", deleteAllItems).Methods("DELETE")
	mux.HandleFunc("/delete/id={id}", deleteItem).Methods("DELETE")
	fmt.Println("Server running on port 8080")
	http.ListenAndServe("localhost:8080", mux)
}