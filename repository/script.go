package repository

import (
	"github.com/lowl11/lazy-entity/internal/repositories/script_repository"
	"github.com/lowl11/lazy-entity/services/script_service"
)

// IScriptRepository read .sql scripts from "resources" folder
type IScriptRepository interface {
	IRepository

	Read() error

	Get(folder, name string) string
	Start(name string) string

	ScriptPath(path string)
	StartPath(path string)
}

func NewScript() IScriptRepository {
	return script_repository.New(script_service.New())
}
