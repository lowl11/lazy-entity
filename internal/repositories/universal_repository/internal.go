package universal_repository

import (
	"strings"
)

const (
	defaultIdName = "id"
)

func (repo *Repository[T, ID]) getFieldList(withID bool) []string {
	fieldList := make([]string, 0, len(repo.fieldList))
	for _, item := range repo.fieldList {
		if !withID && item == repo.idName {
			continue
		}

		fieldName := item

		// join field
		if strings.Contains(fieldName, "_") {
			fieldName = strings.Replace(fieldName, "_", ".", 1) + " " + fieldName
		}

		fieldList = append(fieldList, fieldName)
	}
	return fieldList
}

func (repo *Repository[T, ID]) getNonEmptyFields(okayIndices []int) []string {
	fieldList := make([]string, 0, len(okayIndices))
	for _, index := range okayIndices {
		fieldList = append(fieldList, repo.fieldList[index])
	}
	return fieldList
}
