package sql_helper

import (
	"strings"
)

func ContainsAggregate(value string) bool {
	upper := strings.ToUpper(value)

	return strings.Contains(upper, aggregateCount) ||
		strings.Contains(upper, aggregateMin) ||
		strings.Contains(upper, aggregateMax)
}
