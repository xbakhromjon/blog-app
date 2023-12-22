package domain

// Template domain struct
type Template struct {
	Id   int64
	Name string
}

// TemplateRequest create or update request to business logic
type TemplateRequest struct {
	Name string `json:"name"`
}

type TemplateResponse struct {
	Id   int64
	Name string
}

type TemplateUseCase interface {
	// needed methods
}

type TemplateRepository interface {
	// needed methods
}
