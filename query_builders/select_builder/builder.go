package select_builder

import (
	"github.com/lowl11/lazy-entity/query_builders/select_builder/join_builder"
	"github.com/lowl11/lazy-entity/query_builders/select_builder/where_builder"
)

type Builder struct {
	tableName string
	asName    string

	fields        []string
	orderByFields []string

	joinBuilderList  []*join_builder.Builder
	whereBuilderList []*where_builder.Builder

	err error
}

func New() *Builder {
	return &Builder{
		fields:          []string{},
		orderByFields:   []string{},
		joinBuilderList: make([]*join_builder.Builder, 0),
	}
}
