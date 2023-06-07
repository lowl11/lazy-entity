package query_service

import (
	"github.com/lowl11/lazy-entity/builders/delete_builder"
)

func (service *Service) Delete() string {
	return delete_builder.New(service.tableName).Build()
}
