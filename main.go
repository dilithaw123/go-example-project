package main

import (
	"html/template"
	"log/slog"
	"os"

	"github.com/dilithaw123/go-example-project/internal/http"
	"github.com/dilithaw123/go-example-project/internal/user"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	templates, err := template.ParseGlob("templates/*.html")
	if err != nil {
		logger.Error("Error parsing templates", "error", err)
		return
	}
	handlerEnv := http.HandlerEnv{
		UserService: user.NewMongoUserService(logger),
		Logger:      logger,
		Templates:   templates,
	}
	port := ":8080"
	handlerEnv.StartServer(port)
}
