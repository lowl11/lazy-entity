package condition_builder

import (
	"github.com/lowl11/lazy-entity/internal/helpers/string_helper"
	"github.com/lowl11/lazy-entity/internal/helpers/type_helper"
)

func (builder *Builder) statement(field, statement string, value any) string {
	return string_helper.Concat(builder.getFieldItem(field), statement, type_helper.ToString(value))
}
