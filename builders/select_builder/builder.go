package select_builder

type Builder struct {
	tableName string
	aliasName string
}

func New(tableName, aliasName string) *Builder {
	return &Builder{
		tableName: tableName,
		aliasName: aliasName,
	}
}
