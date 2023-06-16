package grow_select_service

/*
	field = len(aliasName) + 1 (dot) + len(fieldName) + 8 (spaces) + 1 (запятая)
	fieldList = fieldCount * field

	join = 8 (spaces) + len(join type) + 6 (keyword JOIN with spaces) + len(table name) + 4 (keyword AS with spaces) + len(aliasName) + 4 (keyword ON with spaces) + len(condition) + 2 (скобки)
	joinList = joinCount * join
	condition = 8 (spaces) + len(aliasName) + 1 (dot) + len(fieldName) + len(condition_sign) + 2 spaces + len(value) + AND/OR (3/2 len)
	conditionList = conditionCount * condition
	where = 5 (keyword WHERE) + conditionList
	order_by = 9 (keyword order by + space) + len(aliasName) + 1 (dot) + len(fieldName) ASC/DESC(3/4 len)
	group_by = 9 (keyword group by + space) + len(aliasName) + 1 (dot) + len(fieldName)
	having = 8 (keyword having + space) + len(condition)
	offset = 7 (keyword offset + space) + len(value)
	limit = 7 (keyword limit + space) + len(value)
	+ 100 (запас)
*/

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
