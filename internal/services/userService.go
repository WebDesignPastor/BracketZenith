package services

import (
	"BracketZenith/internal/models"
	"BracketZenith/internal/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// GetAllUsers fetches all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

// CreateUser handles the business logic for creating a user
func (s *UserService) CreateUser(user models.User) (int64, error) {
	return s.userRepo.CreateUser(user)
}
