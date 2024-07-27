package usecase

import (
	"crypto/rand"

	"github.com/buemura/health-checker/internal/dto"
	"github.com/buemura/health-checker/internal/entity"
	"github.com/buemura/health-checker/internal/repository"
	"github.com/lucsky/cuid"
)

type CreateEndpoint struct {
	repo repository.EndpointRepository
}

func NewCreateEndpoint(repo repository.EndpointRepository) *CreateEndpoint {
	return &CreateEndpoint{repo: repo}
}

func (uc *CreateEndpoint) Execute(in *dto.CreateEndpointIn) (*entity.Endpoint, error) {
	cuid, err := cuid.NewCrypto(rand.Reader)
	if err != nil {
		return nil, err
	}

	edp, err := uc.repo.Create(&entity.Endpoint{
		ID:             cuid,
		Name:           in.Name,
		Url:            in.Url,
		Status:         in.Status,
		CheckFrequency: in.CheckFrequency,
		NotifyTo:       in.NotifyTo,
	})
	if err != nil {
		return nil, err
	}

	return edp, nil
}
