package script_repository

func (repo *Repository) Script(folder, script string) string {
	return repo.script.Script(folder, script)
}

func (repo *Repository) StartScript(script string) string {
	return repo.script.StartScript(script)
}
