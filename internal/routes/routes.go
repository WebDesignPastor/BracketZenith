package routes

import (
	"BracketZenith/internal/handlers"

	"github.com/gorilla/mux"
)

func NewRouter(userHandler *handlers.UserHandler) *mux.Router {
	router := mux.NewRouter()

	// Route to create a user
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")

	return router
}
