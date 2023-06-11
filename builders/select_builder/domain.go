package select_builder

type joinModel struct {
	TableName  string
	AliasName  string
	Conditions string

	joinType string
}
