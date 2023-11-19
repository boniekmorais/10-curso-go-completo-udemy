package userService

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"example.com/crud-basico/database"
	"github.com/gorilla/mux"
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

	parameters := mux.Vars(r)
	ID, error := strconv.ParseUint(parameters["id"], 10, 32) // Nome do parametro, base decimal, tamanho em bits

	if error != nil {
		w.Write([]byte("Error on converting parameter to integer"))
	}

	db, error := database.Connect()

	if error != nil {
		w.Write([]byte("Error on connecting to database"))
		return
	}

	defer db.Close()

	result, error := db.Query("SELECT * FROM usuarios WHERE id = ?", ID)

	if error != nil {
		w.Write([]byte("Error on selecting users"))
		return
	}

	defer result.Close()

	var user user

	if result.Next() {
		if error := result.Scan(&user.ID, &user.Nome, &user.Email); error != nil {
			w.Write([]byte("Error on getting user data"))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if error := json.NewEncoder(w).Encode(user); error != nil {
		w.Write([]byte("Error on converting user to JSON"))
		return
	}

}

// Update User
func UpdateUserById(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)

	ID, error := strconv.ParseUint(parameters["id"], 10, 32)

	if error != nil {
		w.Write([]byte("Error on converting parameter to integer"))
	}

	requestBody, error := io.ReadAll(r.Body)

	if error != nil {
		w.Write([]byte("Error on reading the request body"))
		return
	}

	var userToUpdate user

	if error = json.Unmarshal(requestBody, &userToUpdate); error != nil {
		w.Write([]byte("Error on parsing the request body"))
		return
	}

	db, error := database.Connect()

	if error != nil {
		w.Write([]byte("Error on connecting to database"))
		return
	}

	defer db.Close()

	statement, error := db.Prepare("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?")

	if error != nil {
		w.Write([]byte("Error on creating statement"))
		return
	}

	defer statement.Close()

	if _, error := statement.Exec(userToUpdate.Nome, userToUpdate.Email, ID); error != nil {
		w.Write([]byte("Error on updating user"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}

// Delete User
func DeleteUserById(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)

	ID, error := strconv.ParseUint(parameters["id"], 10, 32)

	if error != nil {
		w.Write([]byte("Error on converting parameter to integer"))
	}

	db, error := database.Connect()

	if error != nil {
		w.Write([]byte("Error on connecting to database"))
		return
	}

	defer db.Close()

	statement, error := db.Prepare("DELETE FROM usuarios WHERE id = ?")

	if error != nil {
		w.Write([]byte("Error on creating statement"))
		return
	}

	defer statement.Close()

	if _, error := statement.Exec(ID); error != nil {
		w.Write([]byte("Error on deleting user"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}
