package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
)
func fetchDataFromDBJSON(w http.ResponseWriter, db *sql.DB) {
	// Execute the SQL query
	log.Println("Fetching JSON data from database")
	rows, err := db.Query("SELECT * FROM data1")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error JSON", http.StatusInternalServerError)
		return
	}
	log.Println("Fetch JSON data from database successfully")
	defer rows.Close()

	// Create a slice to hold the retrieved data
	var data []map[string]interface{}

	// Process the query results
	var id int
	var skill string

	for rows.Next() {
		err := rows.Scan(&id, &skill)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error API", http.StatusInternalServerError)
			return
		}

		// Create a map for each row
		rowData := map[string]interface{}{
			"ID":   id,
			"skill": skill,
		}

		// Append the map to the data slice
		data = append(data, rowData)
	}

	// Encode the data slice as JSON and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// apiHandlerJSON is an HTTP handler function for the API that returns data as JSON.
func apiHandlerJSON(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fetchDataFromDBJSON(w, db)
}

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
	var skill string

	for rows.Next() {
		err := rows.Scan(&id, &skill)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// You can process the retrieved data here
		fmt.Fprintf(w, "ID: %d, Name: %s\n", id, skill)
	}
}

// dbConnect connects to the database and returns a *sql.DB instance.
func dbConnect() (*sql.DB, error) {
	// Open a database connection
	log.Println("Connecting to database..")
	db, err := sql.Open(DBDriver, fmt.Sprintf("%s:%s@%s", DBUser, DBPassword, DBName))
	if err != nil {
		return nil, err
	}
	log.Println("Connected to database.")
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
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		apiHandlerJSON(w, r, db)
	})
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
