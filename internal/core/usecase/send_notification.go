package usecase

import (
	"log"

	"github.com/buemura/health-checker/internal/core/entity"
	"github.com/buemura/health-checker/internal/core/repository"
)

type SendNotification struct {
	repo repository.NotificationRepository
}

func NewSendNotification(repo repository.NotificationRepository) *SendNotification {
	return &SendNotification{repo: repo}
}

func (uc *SendNotification) Execute(in *entity.Notification) error {
	log.Printf("Sending notification: %v", in)
	return nil
}
