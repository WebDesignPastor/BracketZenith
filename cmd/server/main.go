package main

import (
	"BracketZenith/internal/database"
	"BracketZenith/internal/handlers"
	"BracketZenith/internal/repositories"
	"BracketZenith/internal/routes"
	"BracketZenith/internal/services"
	"log"
	"net/http"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)

	userService := services.NewUserService(userRepo)

	userHandler := handlers.NewUserHandler(userService)

	router := routes.NewRouter(userHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
