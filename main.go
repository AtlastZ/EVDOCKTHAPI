package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)

// fetchDataFromDB retrieves data from the database and writes it to the response writer.
func fetchDataFromDB(w http.ResponseWriter, db *sql.DB) {
	// Execute the SQL query
	rows, err := db.Query("SELECT * FROM data1")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Process the query results
	var id int
	var name string

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// You can process the retrieved data here
		fmt.Fprintf(w, "ID: %d, Name: %s\n", id, name)
	}
}

// dbConnect connects to the database and returns a *sql.DB instance.
func dbConnect() (*sql.DB, error) {
	// Open a database connection
	db, err := sql.Open(DBDriver, fmt.Sprintf("%s:%s@/%s", DBUser, DBPassword, DBName))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// apiHandler is an HTTP handler function for the API.
func apiHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fetchDataFromDB(w, db)
}

func greet(w http.ResponseWriter, r *http.Request) {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	fmt.Fprintf(w, "Hello World! %s", (time.Now().In(location)))
}

func main() {

	db, err := dbConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new HTTP handler function for the API
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		apiHandler(w, r, db)
	})
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
