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

// Get All Users
func ListUsers(w http.ResponseWriter, r *http.Request) {

	db, error := database.Connect()

	if error != nil {
		w.Write([]byte("Error on connecting to database"))
		return
	}

	defer db.Close()

	result, error := db.Query("SELECT * FROM usuarios")

	if error != nil {
		w.Write([]byte("Error on selecting users"))
		return
	}

	defer result.Close()

	var userList []user

	for result.Next() {

		var user user

		if error := result.Scan(&user.ID, &user.Nome, &user.Email); error != nil {
			w.Write([]byte("Error on getting user data"))
			return
		}

		userList = append(userList, user)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if error := json.NewEncoder(w).Encode(userList); error != nil {
		w.Write([]byte("Error on converting users list to JSON"))
		return
	}

	// enc := json.NewEncoder(w)
	// enc.SetIndent("", "    ")

	// if error := enc.Encode(userList); error != nil {
	// 	w.Write([]byte("Error on converting users list to JSON"))
	// 	return
	// }

}

// Get One User By Id
func GetUserById(w http.ResponseWriter, r *http.Request) {

}
