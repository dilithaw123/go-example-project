package http

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (env *HandlerEnv) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idNum, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	user, err := env.UserService.GetUserByID(r.Context(), idNum)
	if err != nil {
		env.Logger.Error("Error getting user by id: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (env *HandlerEnv) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := env.UserService.GetUsers(r.Context())
	if err != nil {
		env.Logger.Error("Error getting all users: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (env *HandlerEnv) GetUsersHTML(w http.ResponseWriter, r *http.Request) {
	users, err := env.UserService.GetUsers(r.Context())
	if err != nil {
		env.Logger.Error("Error getting all users: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err = env.Templates.ExecuteTemplate(w, "users.html", users)
	if err != nil {
		env.Logger.Error("Error executing template", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
