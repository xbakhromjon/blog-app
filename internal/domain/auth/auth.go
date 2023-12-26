package auth

type AuthorizeRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type AuthUseCase interface {
	Authorize(request AuthorizeRequest) (string, error)
}
