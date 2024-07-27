package repository

import "github.com/buemura/health-checker/internal/core/entity"

type EndpointRepository interface {
	FindAll() ([]*entity.Endpoint, error)
	Create(*entity.Endpoint) (*entity.Endpoint, error)
}
