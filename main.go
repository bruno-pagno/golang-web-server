package main

import (
	"fmt"
	"golang-web-server/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetUsers(w, r)
		case "POST":
			handlers.CreateUser(w, r)
		default:
			http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		}
	})

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
