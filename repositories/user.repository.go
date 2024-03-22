package repositories

import "ngobar/berkicau/models"

type UserRepository interface {
	GetDataByEmail(email string) (models.User, error)
	GetDataByUsername(username string) (models.User, error)
	Create(user models.User) (models.User, error)
}
