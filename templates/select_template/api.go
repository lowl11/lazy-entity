package select_template

import (
	"fmt"
	"github.com/lowl11/lazy-entity/entity_models"
	"github.com/lowl11/lazy-entity/services/type_service"
	"strings"
)

const (
	Main = `
SELECT
	<fields>
FROM <table_name> <as_table_name>
	<join_query>
<condition_query>
<order_by_query>
<group_py_query>`

	InnerJoin = `INNER JOIN <join_table> <as_join_table> ON (<join_conditions>)`
	LeftJoin  = `LEFT JOIN <join_table> <as_join_table> ON (<join_conditions>)`

	Where = `WHERE <conditions>`

	OrderBy = "ORDER BY <fields>"

	GroupGy = "GROUP BY"
)

func Fields(asName string, values []string) string {
	if len(values) == 0 {
		return "*"
	}

	if len(asName) > 0 {
		for i := 0; i < len(values); i++ {
			if strings.Contains(values[i], ".") {
				continue
			}
			values[i] = asName + "." + values[i]
		}
	}

	return strings.Join(values, ", ")
}

func OrderByFields(values []string) string {
	if len(values) == 0 {
		return ""
	}

	return strings.Join(values, ", ")
}

func AsName(asTableName string) string {
	if asTableName == "" {
		return ""
	}
	return "AS " + asTableName
}

func JoinCondition(mainTable, joinTable string, condition *entity_models.JoinCondition) string {
	if condition == nil {
		return ""
	}

	return mainTable + "." + condition.MainKey + " = " + joinTable + "." + condition.JoinKey
}

func WhereCondition(condition *entity_models.WhereCondition) string {
	if len(condition.Right) == 0 {
		return ""
	}

	sign := "="
	if condition.More {
		sign = ">"
	} else if condition.Less {
		sign = "<"
	} else if condition.Like {
		sign = "LIKE"
	} else if condition.ILike {
		sign = "ILIKE"
	}

	rightValue := type_service.GetString(condition.Right[0])
	if condition.Between && len(condition.Right) >= 2 {
		rightValue = fmt.Sprintf("%s BETWEEN %s", condition.Right[0], condition.Right[1])
		return condition.Left + rightValue
	}

	if condition.Like || condition.ILike {
		rightValue = "'" + rightValue + "'"
	}

	if condition.Equals {
		if _, ok := condition.Right[0].(string); ok {
			rightValue = "'" + rightValue + "'"
		}
	}

	return condition.Left + " " + sign + " " + rightValue
}
