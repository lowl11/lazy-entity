package crud_repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-entity/internal/repositories"
	"github.com/lowl11/lazy-entity/internal/repositories/universal_repository"
)

type Repository[T any, ID repositories.IComparableID] struct {
	universal_repository.Repository[T, ID]
}

func New[T any, ID repositories.IComparableID](connection *sqlx.DB, tableName string) *Repository[T, ID] {
	return &Repository[T, ID]{
		Repository: universal_repository.New[T, ID](connection, tableName),
	}
}
