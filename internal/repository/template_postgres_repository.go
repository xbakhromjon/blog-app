package repository

import (
	"database/sql"
	"golang-project-template/internal/domain"
)

// templatePostgresRepository postgresql implementation of TemplateRepository
type templatePostgresRepository struct {
	db *sql.DB
}

func NewTemplatePostgresRepository(db *sql.DB) domain.TemplateRepository {
	return &templatePostgresRepository{db: db}
}
