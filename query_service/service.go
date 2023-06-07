package query_service

type Service struct {
	tableName string
}

func New(tableName string) *Service {
	return &Service{
		tableName: tableName,
	}
}
