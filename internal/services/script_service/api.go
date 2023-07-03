package script_service

import (
	"strings"
)

func (service *Service) Read() error {
	service.scripts = make(map[string]any)
	service.startScripts = make(map[string]string)

	if err := service.readStartScripts(); err != nil {
		return err
	}

	if err := service.readScripts(); err != nil {
		return err
	}

	return nil
}

// GetStartScript get script from folder /resources/scripts/start/<script_file>.sql
func (service *Service) GetStartScript(script string) string {
	if service.startScripts == nil || len(service.startScripts) == 0 {
		return ""
	}

	return service.startScripts[script+".sql"]
}

// GetScript get script from folder /resources/script/<folder>/<script_file>.sql
func (service *Service) GetScript(folder, script string) string {
	if service.scripts == nil || len(service.scripts) == 0 {
		return ""
	}

	// remove .sql
	script = strings.ReplaceAll(script, ".sql", "")

	// if there is no folder
	if _, ok := service.scripts[folder].(map[string]string); !ok {
		return ""
	}

	// success case
	return service.scripts[folder].(map[string]string)[script+".sql"]
}
