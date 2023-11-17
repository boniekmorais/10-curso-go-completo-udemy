package userService

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	fmt.Println(userToCreate)
}
