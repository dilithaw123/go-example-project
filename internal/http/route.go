package http

import "net/http"

func (env *HandlerEnv) StartServer(port string) {
	env.Logger.Info("Starting server", "port", port)
	mux := http.NewServeMux()
	env.Route(mux)
	http.ListenAndServe(port, mux)
}

func (env *HandlerEnv) Route(mux *http.ServeMux) {
	mux.Handle("GET /api/users", env.IPMiddleware(http.HandlerFunc(env.GetUsers)))
	mux.Handle("GET /api/user", env.IPMiddleware(http.HandlerFunc(env.GetUserByID)))
	mux.Handle("GET /users", env.IPMiddleware(http.HandlerFunc(env.GetUsersHTML)))
	mux.Handle("GET /", env.IPMiddleware(http.FileServer(http.Dir("./static"))))
}
