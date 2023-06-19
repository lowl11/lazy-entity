package script_repository

import (
	"github.com/lowl11/lazy-entity/internal/repositories/base_repository"
	"github.com/lowl11/lazy-entity/services/script_service"
)

type Repository struct {
	base_repository.Repository
	scriptService *script_service.Service
}

func New(scriptService *script_service.Service) *Repository {
	return &Repository{
		Repository:    base_repository.Repository{},
		scriptService: scriptService,
	}
}
