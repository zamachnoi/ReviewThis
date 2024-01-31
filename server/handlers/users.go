// handlers/user.go
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/zamachnoi/viewthis/data"
	"github.com/zamachnoi/viewthis/models"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        getUserHandler(w, r)
    case "POST":
        createUserHandler(w, r)
    default:
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    user, err := data.GetUserByID(uint(id))
    if err != nil {
        http.Error(w, "Error retrieving user", http.StatusInternalServerError)
        return
    }

    if user == nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(user)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var newUser models.User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    createdUser, err := data.CreateUser(newUser)
    if err != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(createdUser)
}

