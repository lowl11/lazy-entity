package repo_config

type Crud struct {
	// AliasName name which gives as table name.
	// Example: SELECT * FROM questions AS question. So, "question" is an alias
	AliasName string

	// IdName is name of entity primary key (id).
	// Usually, it names "id", but if it is not, use this field
	IdName string

	// PageSize is how many items contain one page.
	// By default page size is 10
	PageSize int

	// Joins is list of joins which can be used in queries.
	// Even if not at all queries need all joins, you need to include it
	Joins []Join

	// ThreadSafe is a flag turning on/off Thread Safe mode.
	// This mode need to protect Repository from concurrency
	ThreadSafe bool

	// Debug is a flag turning on/off Debug mode.
	// Debug mode prints all (fmt.Println()) calling queries/methods
	Debug bool
}
