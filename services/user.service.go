package services

import (
	"ngobar/berkicau/requests"
)

type UserService interface {
	RegisterNewUser(requests.RegistrationRequest) (valid bool, message string)
	// CheckLoginInformation(requests.RegistrationRequest) (valid bool, message string)
}
