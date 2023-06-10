package update_builder

type Builder struct {
	tableName string

	conditions string
	setValues  []string
}

func New(tableName string) *Builder {
	return &Builder{
		tableName: tableName,
		setValues: make([]string, 0),
	}
}
