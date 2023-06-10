package script_repository

import (
	"github.com/lowl11/lazy-entity/internal/repositories/base_repository"
	"github.com/lowl11/lazy-entity/services/script_service"
)

type Repository struct {
	base_repository.Repository
	script *script_service.Service
}

func New(script *script_service.Service) *Repository {
	return &Repository{
		script: script,
	}
}
