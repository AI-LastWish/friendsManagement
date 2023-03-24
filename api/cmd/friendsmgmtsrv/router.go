package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"backend/api/internal/handler/rest/public/user"
)

// UserRouter: User Router
type UserRouter struct {
	handler user.UserHandler
	router  *chi.Mux
}

// NewUserRouter: create new user Router
func NewUserRouter(r user.UserHandler) *UserRouter {
	router := chi.NewRouter()
	return &UserRouter{
		handler: r,
		router:  router,
	}
}

// Routes: Router of users
func (r UserRouter) routes() http.Handler {
	r.router.Get("/users", r.handler.List())
	r.router.Post("/user", r.handler.Get())
	r.router.Post("/invite", r.handler.CreateFriendship())
	r.router.Post("/friends", r.handler.GetFriendList())
	r.router.Post("/common", r.handler.GetCommonFriends())
	r.router.Post("/subscribe", r.handler.CreateSubscribe())
	r.router.Post("/blocks", r.handler.CreateBlock())
	r.router.Post("/retrieve", r.handler.GetRetrieveUpdates())

	return r.router
}
