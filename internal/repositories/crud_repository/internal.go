package crud_repository

const (
	defaultIdName = "id"
)

func (repo *CrudRepository[T, ID]) getFieldList() []string {
	fieldList := make([]string, 0, len(repo.fieldList))
	for _, item := range repo.fieldList {
		if item == repo.idName {
			continue
		}

		fieldList = append(fieldList, item)
	}
	return fieldList
}

func (repo *CrudRepository[T, ID]) getNonEmptyFields(okayIndices []int) []string {
	fieldList := make([]string, 0, len(okayIndices))
	for _, index := range okayIndices {
		fieldList = append(fieldList, repo.fieldList[index])
	}
	return fieldList
}
