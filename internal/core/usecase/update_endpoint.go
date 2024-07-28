package usecase

import (
	"github.com/buemura/health-checker/internal/core/entity"
	"github.com/buemura/health-checker/internal/core/repository"
)

type UpdateEndpoint struct {
	repo repository.EndpointRepository
}

func NewUpdateEndpoint(repo repository.EndpointRepository) *UpdateEndpoint {
	return &UpdateEndpoint{repo: repo}
}

func (uc *UpdateEndpoint) Execute(in *entity.Endpoint) (*entity.Endpoint, error) {
	edp, err := uc.repo.Update(in)
	if err != nil {
		return nil, err
	}
	return edp, nil
}
