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
		alias := strings.Builder{}
		for _, item := range builder.fieldList {
			alias.WriteString("\n\t")
			alias.WriteString(builder.getFieldItem(item))
			alias.WriteString(",")
		}

		return alias.String()[:alias.Len()-1]
	}

	tab := strings.Builder{}
	for _, item := range builder.fieldList {
		tab.WriteString("\n\t")
		tab.WriteString(builder.getFieldItem(item))
		tab.WriteString(", ")
	}

	return tab.String()[:tab.Len()-2]
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

	joinField := strings.Contains(value, ".")

	// check alias name
	var alias string
	if len(builder.aliasName) > 0 && !joinField {
		alias = builder.aliasName + "."
	}

	// check
	if joinField {
		return sql_helper.AliasName(value)
	}

	return alias + value
}

func (builder *Builder) getTableName() string {
	tableName := strings.Builder{}
	if len(builder.aliasName) > 0 {
		tableName.WriteString(builder.tableName)
		tableName.WriteString(" AS ")
		tableName.WriteString(builder.aliasName)
	} else {
		tableName.WriteString(builder.tableName)
	}

	return tableName.String()
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
