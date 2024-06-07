package services

import (
	"ngobar/berkicau/helper"
	"ngobar/berkicau/models"
	"ngobar/berkicau/repositories"
	"ngobar/berkicau/requests"
)

type UserService interface {
	RegisterNewUser(requests.RegistrationRequest) (valid bool, message string)
	CheckLoginInformation(requests.LoginRequest) (valid bool, message string)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService() UserService {
	return &userService{userRepo: repositories.NewUserRepository(helper.DB)}
}

func (s *userService) RegisterNewUser(request requests.RegistrationRequest) (valid bool, message string) {
	err := requests.Validate(request)
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

func (s *userService) CheckLoginInformation(request requests.LoginRequest) (valid bool, message string) {
	return true, "Login success"
}