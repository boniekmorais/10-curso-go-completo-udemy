package userService

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"example.com/crud-basico/database"
)

type user struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// Create User
func CreateUser(w http.ResponseWriter, r *http.Request) {

	requestBody, error := io.ReadAll(r.Body)

	if error != nil {
		w.Write([]byte("Error on reading the request body"))
		return
	}

	var userToCreate user

	if error = json.Unmarshal(requestBody, &userToCreate); error != nil {
		w.Write([]byte("Error on parsing the request body"))
		return
	}

	db, error := database.Connect()

	if error != nil {
		w.Write([]byte("Error on connecting to database"))
		return
	}

	defer db.Close()

	statement, error := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")

	if error != nil {
		w.Write([]byte("Error on creating statement"))
		return
	}

	defer statement.Close()

	result, error := statement.Exec(userToCreate.Nome, userToCreate.Email)

	if error != nil {
		w.Write([]byte("Error on saving user"))
		return
	}

	idSaved, error := result.LastInsertId()

	if error != nil {
		w.Write([]byte("Error on getting ID"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User saved with successful. ID: %d", idSaved)))
}
