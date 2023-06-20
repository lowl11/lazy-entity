package grow_select_service

func (service *Service) Get() int {
	return service.grow + additionalSpace
}

func (service *Service) Fields(count int) *Service {
	service.grow = service.grow + count*avgFieldLen
	return service
}

func (service *Service) Join(tableName, aliasName, conditions string) *Service {
	service.grow += innerJoinKeyword + len(tableName) + len(aliasName) + len(conditions)
	return service
}

func (service *Service) LeftJoin(tableName, aliasName, conditions string) *Service {
	service.grow += leftJoinKeyword + len(tableName) + len(aliasName) + len(conditions)
	return service
}

func (service *Service) RightJoin(tableName, aliasName, conditions string) *Service {
	service.grow += rightJoinKeyword + len(tableName) + len(aliasName) + len(conditions)
	return service
}

func (service *Service) Where(conditions string) *Service {
	service.grow += whereKeyword + len(conditions)
	return service
}

func (service *Service) Table(tableName *string) *Service {
	service.grow += selectKeyword + fromKeyword + len(*tableName)
	return service
}

func (service *Service) Alias(alias *string) *Service {
	service.grow += asKeyword + len(*alias)
	return service
}

func (service *Service) GroupBy(count int) *Service {
	service.grow += groupByKeyword + avgFieldLen*count
	return service
}

func (service *Service) OrderBy(count int) *Service {
	service.grow += orderByKeyword + avgFieldLen*count
	return service
}

func (service *Service) Having() *Service {
	service.grow += havingKeyword + avgFieldLen
	return service
}

func (service *Service) Offset() *Service {
	service.grow += offsetKeyword + avgNumLen
	return service
}

func (service *Service) Limit() *Service {
	service.grow += limitKeyword + avgNumLen
	return service
}
