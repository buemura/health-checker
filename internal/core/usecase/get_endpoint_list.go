package usecase

import (
	"github.com/buemura/health-checker/internal/core/entity"
	"github.com/buemura/health-checker/internal/core/repository"
)

type GetEndpointList struct {
	repo repository.EndpointRepository
}

func NewGetEndpointList(repo repository.EndpointRepository) *GetEndpointList {
	return &GetEndpointList{repo: repo}
}

func (uc *GetEndpointList) Execute() ([]*entity.Endpoint, error) {
	edp, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return edp, nil
}
