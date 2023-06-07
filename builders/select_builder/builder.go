package select_builder

type Builder struct {
	tableName string
	aliasName string
}

func New(tableName string) *Builder {
	return &Builder{
		tableName: tableName,
	}
}
