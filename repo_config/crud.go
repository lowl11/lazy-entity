package repo_config

type Crud struct {
	AliasName string
	IdName    string
	PageSize  int
	Joins     []Join

	ThreadSafe bool
	Debug      bool
}
