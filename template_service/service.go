package template_service

type Service struct {
	template string

	pairList []pair
}

func New(template string) *Service {
	return &Service{
		template: template,
		pairList: make([]pair, 0),
	}
}
