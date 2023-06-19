package script_repository

func (repo *Repository) Read() error {
	return repo.scriptService.Read()
}

func (repo *Repository) Get(folder, name string) string {
	return repo.scriptService.GetScript(folder, name)
}

func (repo *Repository) Start(name string) string {
	return repo.scriptService.GetStartScript(name)
}

func (repo *Repository) StartPath(path string) {
	repo.scriptService.StartPath(path)
}

func (repo *Repository) ScriptPath(path string) {
	repo.scriptService.ScriptPath(path)
}
