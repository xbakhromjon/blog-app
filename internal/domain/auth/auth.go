package auth

type SignUpRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type AuthUseCase interface {
	Signup(request SignUpRequest) (string, error)
}
