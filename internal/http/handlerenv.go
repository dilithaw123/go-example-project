package http

import (
	"html/template"
	"log/slog"

	"github.com/dilithaw123/go-example-project/internal/user"
)

type HandlerEnv struct {
	UserService user.UserService
	Logger      *slog.Logger
	Templates   *template.Template
}
