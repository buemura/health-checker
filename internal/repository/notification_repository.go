package repository

import "github.com/buemura/health-checker/internal/entity"

type NotificationRepository interface {
	FindAll() ([]*entity.Notification, error)
	Create(*entity.Notification) (*entity.Notification, error)
}
