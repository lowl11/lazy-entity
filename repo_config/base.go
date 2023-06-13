package repo_config

import "github.com/lowl11/lazy-entity/builders/select_builder"

type Join struct {
	TableName string
	AliasName string
	Condition func(builder *select_builder.Builder) string
}
