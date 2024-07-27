package usecase

import (
	"crypto/rand"
	"time"

	"github.com/buemura/health-checker/internal/core/dto"
	"github.com/buemura/health-checker/internal/core/entity"
	"github.com/buemura/health-checker/internal/core/repository"
	"github.com/lucsky/cuid"
)

type SaveNotification struct {
	repo repository.NotificationRepository
}

func NewSaveNotification(repo repository.NotificationRepository) *SaveNotification {
	return &SaveNotification{repo: repo}
}

func (uc *SaveNotification) Execute(in *dto.CreateNotificationIn) (*entity.Notification, error) {
	cuid, err := cuid.NewCrypto(rand.Reader)
	if err != nil {
		return nil, err
	}

	notification, err := uc.repo.Create(&entity.Notification{
		ID:          cuid,
		EndpointID:  in.EndpointID,
		Destination: in.Destination,
		CreatedAt:   time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return notification, nil
}
