package join_service

import (
	"github.com/lowl11/lazy-entity/internal/services/condition_service"
	"github.com/lowl11/lazy-entity/internal/services/template_service"
	"github.com/lowl11/lazy-entity/predicates"
	"strings"
)

func (service *Service) Get() string {
	if len(service.joinList) == 0 {
		return ""
	}

	queries := make([]string, 0, len(service.joinList))
	for _, item := range service.joinList {
		queries = append(queries, "\t"+getType(item.Type)+" "+template_service.
			New(template).
			Var("TABLE_NAME", item.TableName).
			Var("ALIAS_NAME", getAlias(item.AliasName)).
			Var("CONDITION_LIST", condition_service.
				New(predicates.And, item.ConditionList).
				NoWhere().
				Alias(service.mainTableName).
				Get()).
			Get())
	}
	return strings.Join(queries, ",\n")
}
