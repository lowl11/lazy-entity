package insert_builder

import "strings"

func (builder *Builder) Build() string {
	queries := make([]string, 0, 3)

	// main template
	main := "INSERT INTO " + builder.tableName + builder.getFields()
	queries = append(queries, main)

	// values template
	if builder.variableMode {
		queries = append(queries, "VALUES ("+builder.getVariableFields()+")")
	} else {
		queries = append(queries, "VALUES ("+builder.getVariables()+")")
	}

	// on conflict template
	if builder.skipConflict {
		queries = append(queries, "ON CONFLICT DO NOTHING")
	} else if builder.onConflict != "" {
		queries = append(queries, "ON CONFLICT DO "+builder.onConflict)
	}

	// returning template
	if len(builder.returningFields) > 0 {
		queries = append(queries, "RETURNING "+strings.Join(builder.returningFields, ", "))
	}

	return strings.Join(queries, "\n")
}

func (builder *Builder) Fields(fields ...string) *Builder {
	builder.fieldList = append(builder.fieldList, fields...)
	return builder
}

func (builder *Builder) Variables(variableList ...any) *Builder {
	builder.variableList = append(builder.variableList, variableList...)
	return builder
}

func (builder *Builder) VariableMode() *Builder {
	builder.variableMode = true
	return builder
}

func (builder *Builder) OnConflict(query string) *Builder {
	builder.onConflict = query
	return builder
}

func (builder *Builder) SkipConflict() *Builder {
	builder.skipConflict = true
	return builder
}

func (builder *Builder) Returning(fields ...string) *Builder {
	builder.returningFields = append(builder.returningFields, fields...)
	return builder
}
