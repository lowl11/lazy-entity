package where_builder

import (
	"github.com/lowl11/lazy-entity/entity_models"
	"github.com/lowl11/lazy-entity/templates/select_template"
	"strings"
)

const (
	Equals  = "equals"
	More    = "more"
	Less    = "less"
	Like    = "like"
	ILike   = "ilike"
	Between = "between"
)

func (builder *Builder) Build() string {
	if len(builder.conditions) == 0 {
		return ""
	}

	conditionQuery := make([]string, 0, len(builder.conditions))
	for _, condition := range builder.conditions {
		itemConditionQuery := select_template.WhereCondition(&condition)
		conditionQuery = append(conditionQuery, itemConditionQuery)
	}

	conditionSeparator := "AND "
	if builder.or {
		conditionSeparator = "OR "
	}

	return strings.Join(conditionQuery, " "+conditionSeparator)
}

func (builder *Builder) Or() *Builder {
	builder.or = true
	return builder
}

func (builder *Builder) Condition(conditionMode, left string, right ...any) *Builder {
	condition := entity_models.WhereCondition{
		Left:  left,
		Right: right,
	}

	switch conditionMode {
	case Equals:
		condition.Equals = true
		break
	case More:
		condition.More = true
		break
	case Less:
		condition.Less = true
		break
	case Like:
		condition.Like = true
		break
	case ILike:
		condition.ILike = true
		break
	case Between:
		condition.Between = true
		break
	}

	builder.conditions = append(builder.conditions, condition)
	return builder
}
