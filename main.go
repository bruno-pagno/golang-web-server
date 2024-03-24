package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	getUsers := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			users := []User{
				{Id: 1, Name: "Foo"},
				{Id: 2, Name: "Bar"},
			}

			usersJSON, err := json.Marshal(users)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(usersJSON)

		case "POST":
			var user User

			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				http.Error(w, "Error decoding JSON body", http.StatusBadRequest)
				return
			}

			if user.Id <= 0 || user.Name == "" {
				http.Error(w, "Missing or invalid 'Id' and/or 'name'", http.StatusBadRequest)
				return
			}

			response, err := json.Marshal(user)
			if err != nil {
				http.Error(w, "Error encoding response JSON", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		default:
			http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		}
	}

	http.HandleFunc("/users", getUsers)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
