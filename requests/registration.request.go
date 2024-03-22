package requests

type RegistrationRequest struct {
	Email                string `json:"email"`
	Username             string `json:"username"`
	FullName             string `json:"fullName"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}
