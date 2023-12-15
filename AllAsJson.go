package main

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	// "time"
)
// fetchDataFromDBJSON retrieves data from the database and writes it as JSON to the response writer.
func fetchDataFromDBJSON(w http.ResponseWriter, db *sql.DB) {
	// Execute the SQL query
	rows, err := db.Query("SELECT * FROM data1")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
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
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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
