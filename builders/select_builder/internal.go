package select_builder

import (
	"github.com/lowl11/lazy-entity/internal/helpers/sql_helper"
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
	"regexp"
	"strings"
)

func (builder *Builder) getFields(query *strings.Builder) {
	if len(builder.fieldList) == 0 {
		query.WriteString("\n\t*")
		return
	}

	if len(builder.aliasName) > 0 {
		for index, item := range builder.fieldList {
			query.WriteString("\n\t")
			query.WriteString(builder.getFieldItem(item))

			if index < len(builder.fieldList)-1 {
				query.WriteString(",")
			}
		}

		return
	}

	for index, item := range builder.fieldList {
		query.WriteString("\n\t")
		query.WriteString(builder.getFieldItem(item))

		if index < len(builder.fieldList)-1 {
			query.WriteString(", ")
		}
	}
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

func (builder *Builder) getTableName(query *strings.Builder) {
	if len(builder.aliasName) > 0 {
		query.WriteString(builder.tableName)
		query.WriteString(" AS ")
		query.WriteString(builder.aliasName)
		return
	}

	query.WriteString(builder.tableName)
}

func (builder *Builder) getOffset(query *strings.Builder) {
	if builder.offset < 0 {
		return
	}

	query.WriteString("OFFSET ")
	query.WriteString(type_helper.ToString(builder.offset))
	query.WriteString("\n")
}

func (builder *Builder) getLimit(query *strings.Builder) {
	if builder.limit < 0 {
		return
	}

	query.WriteString("LIMIT ")
	query.WriteString(type_helper.ToString(builder.limit))
	query.WriteString("\n")
}
