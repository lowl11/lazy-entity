package join_service

import "github.com/lowl11/lazy-entity/internal/entity_domain"

type Service struct {
	mainTableName string
	joinList      []entity_domain.JoinPair
}

func New(mainTableName string, joinList []entity_domain.JoinPair) *Service {
	return &Service{
		mainTableName: mainTableName,
		joinList:      joinList,
	}
}
