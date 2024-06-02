package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"ngobar/berkicau/models"
)

type UserRepository interface {
	GetDataByEmail(email string) (models.User, error)
	GetDataByUsername(username string) (models.User, error)
	Create(user models.User) (uint, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetDataByEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT id, email, username, full_name FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Username, &user.FullName)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetDataByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow("SELECT id, email, username, full_name FROM users WHERE username = ?", username).Scan(&user.ID, &user.Email, &user.Username, &user.FullName)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) Create(user models.User) (uint, error) {
	query := `INSERT INTO users(email, username, full_name, password)
		VALUES (?, ?, ?, ?)`

	result, err := r.db.Exec(query, user.Email, user.Username, user.FullName, user.Password)
	if err != nil {
		message := "error occurred while executing query"
		detailedErrorMessage := fmt.Sprintf("%s: %v", message, err.Error())
		log.Printf("Detailed error message: %v", detailedErrorMessage)
		return 0, fmt.Errorf("failed to execute query '%s': %w", query, err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		message := "error occurred while retrieving results"
		detailedErrorMessage := fmt.Sprintf("%s: %v", message, err.Error())
		log.Printf("Detailed error message: %v", detailedErrorMessage)
		return 0, fmt.Errorf("failed to retriever results '%s': %w", query, err)
	}

	return uint(affectedRows), nil
}