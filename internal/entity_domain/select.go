package entity_domain

type JoinPair struct {
	TableName     string
	AliasName     string
	Type          string
	ConditionList []ConditionPair
}
