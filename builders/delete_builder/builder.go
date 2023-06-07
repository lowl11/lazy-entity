package delete_builder

type Builder struct {
	query string

	tableName string
}

func New(tableName string) *Builder {
	return &Builder{
		tableName: tableName,
	}
}
