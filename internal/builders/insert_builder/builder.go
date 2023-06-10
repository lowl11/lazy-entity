package insert_builder

type Builder struct {
	tableName string

	fieldList       []string
	variableList    []any
	variableMode    bool
	skipConflict    bool
	onConflict      string
	returningFields []string
}

func New(tableName string) *Builder {
	return &Builder{
		tableName:       tableName,
		fieldList:       make([]string, 0),
		variableList:    make([]any, 0),
		returningFields: make([]string, 0),
	}
}
