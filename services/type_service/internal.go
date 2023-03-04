package type_service

func isInt(value any) bool {
	_, ok := value.(int)
	return ok
}

func isString(value any) bool {
	_, ok := value.(string)
	return ok
}

func isBool(value any) bool {
	_, ok := value.(bool)
	return ok
}
