package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Sokorinjo/goapi/internal/db"
)

type User struct {
	ID      int64
	Name    string
	Pass    string
	Created string
}

// Create user
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("INSERT INTO users (user_name, user_pass) VALUES(?, ?)", user.Name, user.Pass)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting user: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User successfuly created."))
}

// Get specific user
func getUser(w http.ResponseWriter, r *http.Request) {
	//Get userID from params
	userId := chi.URLParam(r, "userId")

	var user User

	//Make a query to DB
	err := db.DB.QueryRow("SELECT user_id, user_name, user_pass, created_at FROM users WHERE user_id = ?", userId).Scan(&user.ID, &user.Name, &user.Pass, &user.Created)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "User not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// Delete user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	_, err := db.DB.Exec("DELETE FROM users WHERE user_id=?", userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("User " + userId + " deleted."))
}

// Update user credentials
func updateUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("UPDATE users SET user_name=?, user_pass=? WHERE user_id=?", user.Name, user.Pass, userId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to update user: %s", err.Error()), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated"))
}

// Get all users
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT * FROM users;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Pass, &user.Created); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
