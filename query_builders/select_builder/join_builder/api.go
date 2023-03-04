package join_builder

import (
	"github.com/lowl11/lazy-entity/entity_models"
	"github.com/lowl11/lazy-entity/templates"
	"github.com/lowl11/lazy-entity/templates/select_template"
	"strings"
)

func (builder *Builder) Build(mainTable string) string {
	if builder.tableName == "" {
		return ""
	}

	templateValue := select_template.InnerJoin
	if builder.left {
		templateValue = select_template.LeftJoin
	}

	templateValue = templates.SetVars(templateValue, entity_models.TemplateVar{
		Key:   "join_table",
		Value: builder.tableName,
	}, entity_models.TemplateVar{
		Key:   "as_join_table",
		Value: select_template.AsName(builder.asName),
	})

	conditionTemplateList := make([]string, 0, len(builder.conditions))
	for _, condition := range builder.conditions {
		conditionTemplate := select_template.JoinCondition(mainTable, builder.tableName, &condition)
		conditionTemplateList = append(conditionTemplateList, conditionTemplate)
	}
	templateValue = templates.SetVars(templateValue, entity_models.TemplateVar{
		Key:   "join_conditions",
		Value: strings.Join(conditionTemplateList, " AND "),
	})

	return templateValue
}

func (builder *Builder) As(asName string) *Builder {
	builder.asName = asName
	return builder
}

func (builder *Builder) Conditions(conditions ...entity_models.JoinCondition) *Builder {
	builder.conditions = conditions
	return builder
}

func (builder *Builder) Left() *Builder {
	builder.left = true
	return builder
}
