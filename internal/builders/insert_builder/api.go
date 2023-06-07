package insert_builder

import (
	"github.com/lowl11/lazy-entity/internal/entity_domain"
	"github.com/lowl11/lazy-entity/internal/template_service"
	"strings"
)

func (builder *Builder) Build() string {
	templateList := make([]string, 0, 1)

	// field list
	fieldList := make([]string, 0, len(builder.pairList))
	for _, item := range builder.pairList {
		fieldList = append(fieldList, item.Field)
	}

	// value list
	valueList := make([]string, 0, len(builder.pairList))
	for _, item := range builder.pairList {
		valueList = append(valueList, getValue(item.ValueType, item.Value, item.Field))
	}

	// main template
	main := template_service.
		New(mainTemplate).
		Var("TABLE_NAME", builder.tableName).
		Var("FIELD_LIST", strings.Join(fieldList, ", ")).
		Get()

	// values template
	value := template_service.
		New(valueTemplate).
		Var("VALUE_LIST", strings.Join(valueList, ", ")).
		Get()

	templateList = append(templateList, main, value)

	return strings.Join(templateList, "\n")
}

func (builder *Builder) Pair(pairs ...entity_domain.InsertPair) *Builder {
	builder.pairList = append(builder.pairList, pairs...)
	return builder
}
