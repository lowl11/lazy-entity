package delete_builder

import "github.com/lowl11/lazy-entity/template_service"

func (builder *Builder) Build() string {
	builder.query = template_service.
		New(mainTemplate).
		Var("TABLE_NAME", builder.tableName).
		Get()
	return builder.query
}
