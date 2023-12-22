package usecase

import "golang-project-template/internal/domain"

// templateUseCase interface used for usecase because we can use mock usecase for API contract and easy testing
type templateUseCase struct {
	repository domain.TemplateRepository
}

func NewTemplateUseCase(repository domain.TemplateRepository) domain.TemplateUseCase {

	return &templateUseCase{repository: repository}
}

// implement need methods
