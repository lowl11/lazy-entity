package repo_config

type Crud struct {
	AliasName string
	IdName    string
	Joins     []Join

	ThreadSafe bool
	Debug      bool
}
