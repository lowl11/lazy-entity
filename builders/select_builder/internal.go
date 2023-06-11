package select_builder

import (
	"github.com/lowl11/lazy-entity/internal/helpers/sql_helper"
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"regexp"
	"strings"
)

func (builder *Builder) getFields() string {
	if len(builder.fieldList) == 0 {
		return "\n\t*"
	}

	if len(builder.aliasName) > 0 {
		aliasedFields := make([]string, 0, len(builder.fieldList))
		for _, item := range builder.fieldList {
			aliasedFields = append(aliasedFields, "\n\t"+builder.getFieldItem(item))
		}

		return strings.Join(aliasedFields, ", ")
	}

	tabFieldList := make([]string, 0, len(builder.fieldList))
	for _, item := range builder.fieldList {
		tabFieldList = append(tabFieldList, "\n\t"+item)
	}
	return strings.Join(tabFieldList, ", ")
}

func (builder *Builder) getFieldItem(value string) string {
	// check aggregate function
	if sql_helper.ContainsAggregate(value) {
		reg, _ := regexp.Compile(".*?\\((.*?)\\)")
		match := reg.FindAllStringSubmatch(value, -1)

		if len(match) > 0 && len(match[0]) > 1 {
			return strings.ReplaceAll(value, match[0][1], builder.getFieldItem(match[0][1]))
		}
	}

	// check alias name
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
