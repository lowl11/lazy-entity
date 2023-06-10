package crud_repository

const (
	defaultIdName = "id"
)

func (repo *CrudRepository[T, ID]) fieldListWithoutID() []string {
	if len(repo.fieldList) == 0 {
		return repo.fieldList
	}

	return repo.fieldList[1:]
}
