package main

import (
	"log/slog"
	"net/http"

	"github.com/dilithaw123/go-example-project/internal/user"
)

func main() {
	logger := slog.Default()
	userHandlerEnv := user.NewUserHandlerEnv(user.NewMongoUserService(), logger)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /user", userHandlerEnv.GetUserByID)
	mux.Handle("GET /", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", mux)
}
