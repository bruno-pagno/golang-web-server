package handlers

import (
	"encoding/json"
	"golang-web-server/models"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := []models.User{
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

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

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
}
