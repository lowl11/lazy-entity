package update_builder

type Builder struct {
	tableName string

	conditions string
	setValues  []string
	setFields  []string
}

func New(tableName string) *Builder {
	return &Builder{
		tableName: tableName,
		setValues: make([]string, 0),
		setFields: make([]string, 0),
	}
}
