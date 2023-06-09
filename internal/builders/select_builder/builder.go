package select_builder

type Builder struct {
	fieldList []string

	tableName  string
	aliasName  string
	joinList   []joinModel
	conditions string
}

func New(fields ...string) *Builder {
	return &Builder{
		fieldList: fields,
		joinList:  make([]joinModel, 0),
	}
}
