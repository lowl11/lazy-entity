package select_builder

import (
	"github.com/lowl11/lazy-collection/array"
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"strings"
)

func (builder *Builder) getFields() string {
	if len(builder.fieldList) == 0 {
		return "*"
	}

	if len(builder.aliasName) > 0 {
		aliasedFields := make([]string, 0, len(builder.fieldList))
		array.NewWithList[string](builder.fieldList...).Each(func(item string) {
			if strings.Contains(item, ".") {
				aliasedFields = append(aliasedFields, item)
				return
			}

			aliasedFields = append(aliasedFields, builder.aliasName+"."+item)
		})

		return strings.Join(aliasedFields, ", ")
	}

	return strings.Join(builder.fieldList, ", ")
}

func (builder *Builder) getFieldItem(value string) string {
	var alias string
	if len(builder.aliasName) > 0 && !strings.Contains(value, ".") {
		alias = builder.aliasName + "."
	}

	return alias + value
}

func (builder *Builder) getTableName() string {
	var alias string
	if len(builder.aliasName) > 0 {
		alias = " AS " + builder.aliasName
	}
	return builder.tableName + alias
}

func (builder *Builder) getOffset() string {
	if builder.offset < 0 {
		return ""
	}

	return "OFFSET " + type_helper.ToString(builder.offset)
}

func (builder *Builder) getLimit() string {
	if builder.limit < 0 {
		return ""
	}

	return "LIMIT " + type_helper.ToString(builder.limit)
}
