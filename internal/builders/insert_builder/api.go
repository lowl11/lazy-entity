package insert_builder

import "strings"

/*
insert into guarantees (
        id,
        sum,
        status,

        initiator_iin,
        initiator_colvir_id,
        initiator_phone,

        principial_bin,
        principial_company_type,
        principial_colvir_id
)
values (
        :id,
        :sum,
        :status,

        :initiator_iin,
        :initiator_colvir_id,
        :initiator_phone,

        :principial_bin,
        :principial_company_type,
        :principial_colvir_id
)
*/

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
