package select_builder

import (
	"github.com/lowl11/lazy-entity/builders/condition_builder"
	"github.com/lowl11/lazy-entity/order_types"
	"strings"
)

type Builder struct {
	condition_builder.Builder

	fieldList []string
	tableName string
	aliasName string
	joinList  []joinModel
	//conditions       string
	conditions       *strings.Builder
	orderFields      []string
	orderType        string
	havingExpression string
	groupByFields    []string
	offset           int
	limit            int
}

func New(fields ...string) *Builder {
	builder := &Builder{
		fieldList:     fields,
		conditions:    &strings.Builder{},
		joinList:      make([]joinModel, 0),
		orderFields:   make([]string, 0),
		orderType:     order_types.Asc,
		groupByFields: make([]string, 0),
		offset:        -1,
		limit:         -1,
	}

	builder.Builder = condition_builder.New(builder.getFieldItem)
	return builder
}
