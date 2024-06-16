package user

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type UserHandlerEnv struct {
	UserService UserService
	Logger      *slog.Logger
}

func NewUserHandlerEnv(userService UserService, logger *slog.Logger) *UserHandlerEnv {
	return &UserHandlerEnv{
		UserService: userService,
		Logger:      logger,
	}
}

func (env *UserHandlerEnv) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	user, err := env.UserService.GetUserByID(r.Context(), id)
	if err != nil {
		env.Logger.Error("failed to get user by id: %v", err)
		http.Error(w, "failed to get user by id", http.StatusInternalServerError)
		return
	}
	j, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func ApplyUserRoutes(mux *http.ServeMux, env *UserHandlerEnv) {
	mux.HandleFunc("GET /user", env.GetUserByID)
}
