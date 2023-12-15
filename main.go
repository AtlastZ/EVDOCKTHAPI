package main

import (
	// "database/sql"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	// "time"
)

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
