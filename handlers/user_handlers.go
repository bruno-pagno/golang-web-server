package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-web-server/models"
	"net/http"
)

var db *sql.DB

func SetDatabase(database *sql.DB) {
	db = database
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding JSON body", http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "Missing 'name'", http.StatusBadRequest)
		return
	}

	statement := `INSERT INTO users (name) VALUES ($1) RETURNING id`
	err = db.QueryRow(statement, user.Name).Scan(&user.Id)
	if err != nil {
		http.Error(w, "Error inserting new user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	response, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding JSON body", http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "Missing 'name'", http.StatusBadRequest)
		return
	}

	statement := `UPDATE users SET name = $2 WHERE id = $1`
	_, err = db.Exec(statement, userID, user.Name)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	statement := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(statement, userID)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID %s deleted successfully", userID)
}

func UsersRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUsers(w, r)
	case http.MethodPost:
		CreateUser(w, r)
	case http.MethodPut:
		UpdateUser(w, r)
	case http.MethodDelete:
		DeleteUser(w, r)
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}
}
