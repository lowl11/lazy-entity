package repository

import (
	"github.com/lowl11/lazy-entity/internal/repositories/script_repository"
	"github.com/lowl11/lazy-entity/internal/services/script_service"
)

// IScriptRepository read .sql scripts from "resources" folder
type IScriptRepository interface {
	IRepository

	// Read get all scripts' content to the memory
	Read() error

	// Get returns content of chosen script.
	// If folder or script name will be wrong, method will return empty string
	Get(folder, name string) string

	// Start returns script from "start" folder.
	// If script name will be wrong, method will return empty string
	Start(name string) string
}

func NewScript() IScriptRepository {
	return script_repository.New(script_service.New())
}
