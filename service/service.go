package service

import (
	"github.com/sauravgsh16/graphql-go/domain"
)

// Service interface
type Service interface {
	Get(int) (domain.Object, error)
	Create(domain.Object) (domain.Object, error)
	Delete(int) error
	GetAll() ([]domain.Object, error)
	GetAllByKey(int) (interface{}, error)
}
