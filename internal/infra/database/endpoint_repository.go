package database

import (
	"context"
	"time"

	"github.com/buemura/health-checker/internal/core/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EndpointRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewEndpointRepositoryImpl(db *pgxpool.Pool) *EndpointRepositoryImpl {
	return &EndpointRepositoryImpl{db: db}
}

func (r *EndpointRepositoryImpl) Create(e *entity.Endpoint) (*entity.Endpoint, error) {
	_, err := r.db.Exec(context.Background(), "INSERT INTO endpoints (id, name, url, status, check_frequency, last_checked, notify_to) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		e.ID, e.Name, e.Url, e.Status, e.CheckFrequency, e.LastChecked, e.NotifyTo)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *EndpointRepositoryImpl) FindAll() ([]*entity.Endpoint, error) {
	rows, err := r.db.Query(context.Background(), "SELECT id, name, url, status, check_frequency, last_checked, notify_to FROM endpoints")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	endpoints := []*entity.Endpoint{}
	for rows.Next() {
		var id, name, url, status, notifyTo string
		var checkFrequency int
		var lastChecked time.Time

		if err := rows.Scan(&id, &name, &url, &status, &checkFrequency, &lastChecked, &notifyTo); err != nil {
			return nil, err
		}
		endpoints = append(endpoints, &entity.Endpoint{
			ID:             id,
			Name:           name,
			Url:            url,
			Status:         status,
			CheckFrequency: checkFrequency,
			LastChecked:    &lastChecked,
			NotifyTo:       notifyTo,
		})
	}
	return endpoints, nil
}
