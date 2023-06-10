package script_service

import "strings"

// StartScript get script from folder /resources/scripts/start/<script_file>.sql
func (event *Service) StartScript(script string) string {
	return event.startScripts[script+".sql"]
}

// Script get script from folder /resources/script/<folder>/<script_file>.sql
func (event *Service) Script(folder, script string) string {
	// remove .sql
	script = strings.ReplaceAll(script, ".sql", "")

	// if there is no folder
	if _, ok := event.scripts[folder].(map[string]string); !ok {
		return ""
	}

	// success case
	return event.scripts[folder].(map[string]string)[script+".sql"]
}
