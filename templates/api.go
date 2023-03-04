package templates

import (
	"github.com/lowl11/lazy-entity/entity_models"
	"strings"
)

func SetVars(templateValue string, vars ...entity_models.TemplateVar) string {
	for _, variable := range vars {
		key := "<" + variable.Key + ">"
		templateValue = strings.ReplaceAll(templateValue, key, variable.Value)
	}
	return templateValue
}

func FilterQuery(query string) string {
	return strings.TrimSpace(query)
}
