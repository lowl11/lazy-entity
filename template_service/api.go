package template_service

import "strings"

func (service *Service) Var(key, value string) *Service {
	service.pairList = append(service.pairList, pair{
		Key:   key,
		Value: value,
	})
	return service
}

func (service *Service) Get() string {
	result := service.template
	for _, pair := range service.pairList {
		result = strings.ReplaceAll(result, "{{"+pair.Key+"}}", pair.Value)
	}
	return result
}
