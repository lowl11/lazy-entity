package insert_builder

import "github.com/lowl11/lazy-entity/internal/entity_domain"

type Builder struct {
	tableName string
	pairList  []entity_domain.InsertPair
}

func New(tableName string) *Builder {
	return &Builder{
		tableName: tableName,
		pairList:  make([]entity_domain.InsertPair, 0),
	}
}
