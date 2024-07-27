package database

import (
	"database/sql"
	"time"

	"github.com/buemura/health-checker/internal/entity"
)

type NotificationRepositoryImpl struct {
	db *sql.DB
}

func NewNotificationRepositoryImpl(db *sql.DB) *NotificationRepositoryImpl {
	return &NotificationRepositoryImpl{db: db}
}

func (r *NotificationRepositoryImpl) Create(n *entity.Notification) (*entity.Notification, error) {
	_, err := r.db.Exec("INSERT INTO notifications (id, endpoint_id, destination, created_at) VALUES ($1, $2, $3, $4)",
		n.ID, n.EndpointID, n.Destination, n.CreatedAt)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (r *NotificationRepositoryImpl) FindAll() ([]*entity.Notification, error) {
	rows, err := r.db.Query("SELECT id, endpoint_id, destination, created_at FROM notifications")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	notifications := []*entity.Notification{}
	for rows.Next() {
		var id, endpointId, destination string
		var createdAt time.Time

		if err := rows.Scan(&id, &endpointId, &destination, &createdAt); err != nil {
			return nil, err
		}
		notifications = append(notifications, &entity.Notification{
			ID:          id,
			EndpointID:  endpointId,
			Destination: destination,
			CreatedAt:   createdAt,
		})
	}
	return notifications, nil
}
