package requests

import "errors"

type RegistrationRequest struct {
	Email                string `json:"email"`
	Username             string `json:"username"`
	FullName             string `json:"fullName"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

func (r *RegistrationRequest) Validate() error {
	if r.Email == "" {
		return errors.New("email is required")
	}

	if r.Username == "" {
		return errors.New("username is required")
	}

	if r.FullName == "" {
		return errors.New("full name is required")
	}

	if r.Password == "" {
		return errors.New("password is required")
	}

	if r.PasswordConfirmation != r.Password {
		return errors.New("password does not match")
	}

	return nil
}