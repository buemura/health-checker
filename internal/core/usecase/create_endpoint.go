package usecase

import (
	"crypto/rand"
	"time"

	"github.com/buemura/health-checker/internal/core/dto"
	"github.com/buemura/health-checker/internal/core/entity"
	"github.com/buemura/health-checker/internal/core/repository"
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
		Status:         "UP",
		CheckFrequency: in.CheckFrequency,
		LastChecked:    time.Now(),
		NotifyTo:       in.NotifyTo,
	})
	if err != nil {
		return nil, err
	}

	return edp, nil
}
