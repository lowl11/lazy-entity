package update_builder

type Builder struct {
	tableName string
}

func New(tableName string) *Builder {
	return &Builder{
		tableName: tableName,
	}
}
