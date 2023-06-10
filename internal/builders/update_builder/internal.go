package update_builder

import (
	"github.com/lowl11/lazy-entity/internal/helpers/sql_helper"
	"regexp"
	"strings"
)

func (builder *Builder) getFieldItem(value string) string {
	// check aggregate function
	if sql_helper.ContainsAggregate(value) {
		reg, _ := regexp.Compile(".*?\\((.*?)\\)")
		match := reg.FindAllStringSubmatch(value, -1)

		if len(match) > 0 && len(match[0]) > 1 {
			return strings.ReplaceAll(value, match[0][1], builder.getFieldItem(match[0][1]))
		}
	}

	return value
}
