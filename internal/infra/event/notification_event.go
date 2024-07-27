package event

import (
	"github.com/buemura/health-checker/internal/core/dto"
	"github.com/buemura/health-checker/internal/core/entity"
	"github.com/buemura/health-checker/internal/core/usecase"
	"github.com/buemura/health-checker/internal/infra/database"
)

type NotificationEvent struct {
}

func NewNotificationEvent() *NotificationEvent {
	return &NotificationEvent{}
}

func (e *NotificationEvent) SendNotification(in *dto.CreateNotificationIn) (*entity.Notification, error) {
	nr := database.NewNotificationRepositoryImpl(database.DB)
	saveNotificationUC := usecase.NewSaveNotification(nr)
	sendNotificationUC := usecase.NewSendNotification(nr)

	notif, err := saveNotificationUC.Execute(in)
	if err != nil {
		return nil, err
	}

	if err := sendNotificationUC.Execute(notif); err != nil {
		return nil, err
	}

	return notif, nil
}
