package repository

import (
	"github.com/lowl11/lazy-entity/internal/repositories/script_repository"
	"github.com/lowl11/lazy-entity/services/script_service"
)

func NewScript(scriptService *script_service.Service) *script_repository.Repository {
	return script_repository.New(scriptService)
}
