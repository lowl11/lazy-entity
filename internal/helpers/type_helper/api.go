package type_helper

import (
	"encoding/json"
	"fmt"
	"github.com/lowl11/lazy-entity/internal/join_field"
	"reflect"
	"strconv"
)

func ToString(anyValue any) string {
	if _, ok := anyValue.(error); ok {
		return anyValue.(error).Error()
	}

	value := reflect.ValueOf(anyValue)

	switch value.Kind() {
	case reflect.String:
		stringValue := anyValue.(string)

		// if this is counter variable like $1
		if len(stringValue) > 1 && stringValue[0] == '$' && IsNumber(rune(stringValue[1])) {
			return stringValue
		}

		// if this is variable like :id
		if stringValue != "" && stringValue[0] == ':' {
			return stringValue
		}

		// if this is variable for join fields
		if stringValue != "" && stringValue[0] == '$' {
			return stringValue[1:]
		}

		return "'" + stringValue + "'"
	case reflect.Bool:
		return strconv.FormatBool(anyValue.(bool))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32:
		return fmt.Sprintf("%f", value.Float())
	case reflect.Float64:
		return fmt.Sprintf("%g", value.Float())
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		valueInBytes, err := json.Marshal(anyValue)
		if err != nil {
			return ""
		}
		return string(valueInBytes)
	case reflect.Ptr:
		return ToString(value.Elem().Interface())
	default:
		return fmt.Sprintf("%v", value)
	}
}

func IsNumber(value rune) bool {
	switch value {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}

	return false
}

func GetStructFields[T any]() ([]string, []join_field.Field) {
	var object T
	element := reflect.TypeOf(object)
	fieldList := make([]string, 0, element.NumField())
	joinFieldList := make([]join_field.Field, 0, element.NumField())

	for i := 0; i < element.NumField(); i++ {
		fieldValue := element.Field(i).Tag.Get("db")
		joinName, joinFound := element.Field(i).Tag.Lookup("join")

		if fieldValue == "" {
			continue
		}

		if joinFound {
			joinFieldList = append(joinFieldList, join_field.Field{
				Name:     fieldValue,
				Relation: joinName,
			})
			continue
		}

		fieldList = append(fieldList, fieldValue)
	}

	return fieldList, joinFieldList
}

func GetStructFieldsByObject(object any) []string {
	element := reflect.TypeOf(object)
	fieldList := make([]string, 0, element.NumField())
	for i := 0; i < element.NumField(); i++ {
		fieldList = append(fieldList, element.Field(i).Tag.Get("db"))
	}
	return fieldList
}

func GetObjectNonEmptyIndices(object any) []int {
	element := reflect.ValueOf(object).Elem()

	list := make([]int, 0, element.NumField())
	for i := 0; i < element.NumField(); i++ {
		if !element.Field(i).IsZero() {
			list = append(list, i)
		}
	}
	return list
}
