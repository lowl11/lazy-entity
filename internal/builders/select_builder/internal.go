package select_builder

import (
	"github.com/lowl11/lazy-collection/array"
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
	if len(builder.aliasName) > 0 {
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
