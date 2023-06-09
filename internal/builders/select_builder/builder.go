package select_builder

type Builder struct {
	fieldList []string

	tableName  string
	aliasName  string
	joinList   []joinModel
	conditions string
	offset     int
	limit      int
}

func New(fields ...string) *Builder {
	return &Builder{
		fieldList: fields,
		joinList:  make([]joinModel, 0),
		offset:    -1,
		limit:     -1,
	}
}
