package grow_select_service

type Service struct {
	grow int
}

func New() *Service {
	return &Service{}
}
