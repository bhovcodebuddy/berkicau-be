package services

import (
	"ngobar/berkicau/helper"
	"ngobar/berkicau/middlewares"
	"ngobar/berkicau/models"
	"ngobar/berkicau/repositories"
	"ngobar/berkicau/requests"
	"os"
)

type UserService interface {
	RegisterNewUser(requests.RegistrationRequest) (valid bool, message string)
	GetLoginInformation(requests.LoginRequest) (valid bool, message string, token *string)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService() UserService {
	return &userService{userRepo: repositories.NewUserRepository(helper.DB)}
}

func (s *userService) RegisterNewUser(request requests.RegistrationRequest) (valid bool, message string) {
	err := request.Validate()
	if err != nil {
		return false, err.Error()
	}

	user, err := s.userRepo.GetDataByEmail(request.Email)
	if err != nil {
		return false, err.Error()
	}

	if user.Email != "" {
		return false, "Email is already registered"
	}

	user, err = s.userRepo.GetDataByUsername(request.Username)
	if err != nil {
		return false, err.Error()
	}

	if user.Username != "" {
		return false, "Username is already used"
	}

	if request.Password != request.PasswordConfirmation {
		return false, "Password does not match"
	}

	user = models.User{
		Email: request.Email,
		Username: request.Username,
		FullName: request.FullName,
		Password: request.Password,
	}

	_, err = s.userRepo.Create(user)
	if err != nil {
		return false, err.Error()
	}

	return true, "Registration success"
}

func (s *userService) GetLoginInformation(request requests.LoginRequest) (valid bool, message string, token *string) {
	err := request.Validate()
	if err != nil {
		return false, err.Error(), nil
	}

	user, err := s.userRepo.GetDataByUsername(request.Username)
	if err != nil {
		return false, err.Error(), nil
	}

	if request.Password != user.Password {
		return false, "Password does not match", nil
	}

	result, err := middlewares.Authenticate(user.ID, os.Getenv("JWT_SECRET"))
	if err != nil {
		return false, err.Error(), nil
	}

	return true, "Login success", &result
}