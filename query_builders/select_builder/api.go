package select_builder

import (
	"errors"
	"github.com/lowl11/lazy-entity/entity_models"
	"github.com/lowl11/lazy-entity/query_builders/select_builder/join_builder"
	"github.com/lowl11/lazy-entity/query_builders/select_builder/where_builder"
	"github.com/lowl11/lazy-entity/templates"
	"github.com/lowl11/lazy-entity/templates/select_template"
)

func (builder *Builder) Build() (string, error) {
	templateValue := select_template.Main

	templateValue = templates.SetVars(templateValue, entity_models.TemplateVar{
		Key:   "table_name",
		Value: builder.tableName,
	}, entity_models.TemplateVar{
		Key:   "as_table_name",
		Value: select_template.AsName(builder.asName),
	}, entity_models.TemplateVar{
		Key:   "fields",
		Value: select_template.Fields(builder.asName, builder.fields),
	}, entity_models.TemplateVar{
		Key:   "order_by_query",
		Value: "",
	}, entity_models.TemplateVar{
		Key: "group_py_query",
		Value: templates.SetVars(select_template.OrderBy, entity_models.TemplateVar{
			Key:   "fields",
			Value: select_template.OrderByFields(builder.orderByFields),
		}),
	})

	// join query
	var joinQuery string
	if len(builder.joinBuilderList) > 0 {
		for _, joinBuilder := range builder.joinBuilderList {
			tableName := builder.tableName
			if len(builder.asName) > 0 {
				tableName = builder.asName
			}

			joinQuery += joinBuilder.Build(tableName) + "\n\t"
		}
	}

	templateValue = templates.SetVars(templateValue, entity_models.TemplateVar{
		Key:   "join_query",
		Value: joinQuery,
	})

	// where query
	if builder.whereBuilderList != nil {
		whereQuery := "WHERE "
		for index, whereBuilder := range builder.whereBuilderList {
			whereQuery += "(" + whereBuilder.Build() + ")"
			if len(builder.whereBuilderList) > 0 && index < len(builder.whereBuilderList)-1 {
				whereQuery += " AND "
			}
			whereQuery += "\n"
		}
		templateValue = templates.SetVars(templateValue, entity_models.TemplateVar{
			Key:   "condition_query",
			Value: whereQuery,
		})
	}

	templateValue = templates.FilterQuery(templateValue)
	return templateValue, builder.err
}

func (builder *Builder) Table(tableName string) *Builder {
	if tableName == "" {
		builder.err = errors.New("table name is empty")
	}

	builder.tableName = tableName
	return builder
}

func (builder *Builder) As(asName string) *Builder {
	builder.asName = asName
	return builder
}

func (builder *Builder) Fields(fields ...string) *Builder {
	builder.fields = fields
	return builder
}

func (builder *Builder) Join(joinBuilder *join_builder.Builder) *Builder {
	builder.joinBuilderList = append(builder.joinBuilderList, joinBuilder)
	return builder
}

func (builder *Builder) Where(whereBuilder *where_builder.Builder) *Builder {
	builder.whereBuilderList = append(builder.whereBuilderList, whereBuilder)
	return builder
}

func (builder *Builder) OrderBy(fields ...string) *Builder {
	builder.orderByFields = fields
	return builder
}
