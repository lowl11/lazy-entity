package insert_builder

type Builder struct {
	tableName string

	fieldList    []string
	variableMode bool
}

func New(tableName string) *Builder {
	return &Builder{
		tableName: tableName,
		fieldList: make([]string, 0),
	}
}
