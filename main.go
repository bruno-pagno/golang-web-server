package main

import (
	"database/sql"
	"fmt"
	"golang-web-server/handlers"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres dbname=postgres password=postgres host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	handlers.SetDatabase(db)
	fmt.Println("Connected to the POSTGRESQL database")

	http.HandleFunc("/users", handlers.UsersRouter)
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
