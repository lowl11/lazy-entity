package universal_repository

import (
	"strings"
)

const (
	defaultIdName = "id"
)

func (repo *Repository[T, ID]) getFields(withID bool) []string {
	fieldList := make([]string, 0, len(repo.fieldList))
	for _, item := range repo.fieldList {
		if !withID && item == repo.idName {
			continue
		}

		fieldList = append(fieldList, item)
	}
	return fieldList
}

func (repo *Repository[T, ID]) getFieldsWithJoin(withID bool) []string {
	return append(repo.getFields(withID), repo.getJoinFields()...)
}

func (repo *Repository[T, ID]) getJoinFields() []string {
	fields := make([]string, 0, len(repo.joinFieldList))
	for _, item := range repo.joinList {
		for _, field := range repo.joinFieldList {
			if item.TableName == field.Relation {
				originalFieldName := strings.Replace(field.Name, "_", ".", 1)
				fields = append(fields, originalFieldName+" "+field.Name)
			}
		}
	}
	return fields
}

func (repo *Repository[T, ID]) getNonEmptyFields(okayIndices []int) []string {
	fieldList := make([]string, 0, len(okayIndices))
	fieldsLen := len(repo.fieldList)

	for _, index := range okayIndices {
		if index > fieldsLen {
			continue
		}

		fieldList = append(fieldList, repo.fieldList[index])
	}

	return fieldList
}

func (repo *Repository[T, ID]) lock() {
	if !repo.threadSafe {
		return
	}

	repo.mutex.Lock()
}

func (repo *Repository[T, ID]) unlock() {
	if !repo.threadSafe {
		return
	}

	repo.mutex.Unlock()
}
