package script_repository

func (repo *Repository) Script(folder, script string) string {
	return repo.script.GetScript(folder, script)
}

func (repo *Repository) StartScript(script string) string {
	return repo.script.GetStartScript(script)
}
