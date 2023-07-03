package script_service

import (
	"github.com/lowl11/lazyfile/data/interfaces"
	"github.com/lowl11/lazyfile/fmanager"
)

type Service struct {
	manager interfaces.IManager

	startScripts map[string]string
	scripts      map[string]any
}

func New() *Service {
	return &Service{
		manager: fmanager.New(defaultResourcesPath),
	}
}
