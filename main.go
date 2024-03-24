package main

import (
	"log"
	"net/http"
    "fmt"
    "encoding/json"
)

type User struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

func main() {
	getUsers := func(w http.ResponseWriter, _ *http.Request) {
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
	}

	http.HandleFunc("/users", getUsers)

    fmt.Printf("Starting server at port 8080\n")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
