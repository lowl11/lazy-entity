package delete_builder

type Builder struct {
	tableName string

	conditions string
}

func New(tableName string) *Builder {
	return &Builder{
		tableName: tableName,
	}
}
