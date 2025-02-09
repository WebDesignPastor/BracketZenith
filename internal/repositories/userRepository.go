package repositories

import (
	"BracketZenith/internal/models"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetAllUsers fetches all users from the database
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, first_name, last_name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(user models.User) (int64, error) {
	result, err := r.db.Exec(
		"INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)",
		user.FirstName, user.LastName, user.Email, user.Password,
	)
	if err != nil {
		return 0, err
	}

	// Return the ID of the newly created user
	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil
}
