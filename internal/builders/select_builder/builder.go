package select_builder

import "github.com/lowl11/lazy-entity/order_types"

type Builder struct {
	fieldList []string

	tableName        string
	aliasName        string
	joinList         []joinModel
	conditions       string
	orderFields      []string
	orderType        string
	havingExpression string
	offset           int
	limit            int
}

func New(fields ...string) *Builder {
	return &Builder{
		fieldList:   fields,
		joinList:    make([]joinModel, 0),
		orderFields: make([]string, 0),
		orderType:   order_types.Asc,
		offset:      -1,
		limit:       -1,
	}
}
