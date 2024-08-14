package handlers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID	string `json:"id"`
	Name	string `json:"name"`
	Email string `json:"email"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := User{ID: "1", Name: "John Doe", Email: "john@example.com"}
	json.NewEncoder(w).Encode(user)
}