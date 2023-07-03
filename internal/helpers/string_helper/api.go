package string_helper

import "strings"

func Concat(args ...string) string {
	if len(args) == 0 {
		return ""
	}

	builder := strings.Builder{}

	var grow int
	for _, arg := range args {
		grow += len(arg)
	}
	builder.Grow(grow)

	for _, arg := range args {
		builder.WriteString(arg)
	}

	return builder.String()
}
