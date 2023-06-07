package join_service

import "github.com/lowl11/lazy-entity/join_types"

const (
	template = "JOIN {{TABLE_NAME}}{{ALIAS_NAME}} on ({{CONDITION_LIST}})"
)

func getAlias(aliasName string) string {
	if aliasName == "" {
		return ""
	}

	return " AS " + aliasName
}

func getType(joinType string) string {
	if joinType == "" {
		return join_types.Inner
	}

	return joinType
}
