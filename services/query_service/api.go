package query_service

import "github.com/lowl11/lazy-entity/query_builders/select_builder"

func Select() *select_builder.Builder {
	return select_builder.New()
}
