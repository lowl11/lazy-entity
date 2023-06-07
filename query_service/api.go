package query_service

import (
	"github.com/lowl11/lazy-entity/builders/delete_builder"
	"github.com/lowl11/lazy-entity/builders/insert_builder"
	"github.com/lowl11/lazy-entity/builders/select_builder"
	"github.com/lowl11/lazy-entity/builders/update_builder"
)

func (service *Service) Select() *select_builder.Builder {
	return select_builder.
		New(service.tableName, service.aliasName)
}

func (service *Service) Insert() *insert_builder.Builder {
	return insert_builder.
		New(service.tableName)
}

func (service *Service) Update() *update_builder.Builder {
	return update_builder.
		New(service.tableName)
}

func (service *Service) Delete() *delete_builder.Builder {
	return delete_builder.
		New(service.tableName)
}
