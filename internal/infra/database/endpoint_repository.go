package database

import (
	"context"
	"log"

	"github.com/buemura/health-checker/internal/core/entity"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EndpointRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewEndpointRepositoryImpl(db *pgxpool.Pool) *EndpointRepositoryImpl {
	return &EndpointRepositoryImpl{db: db}
}

func (r *EndpointRepositoryImpl) Create(e *entity.Endpoint) (*entity.Endpoint, error) {
	_, err := r.db.Exec(context.Background(), "INSERT INTO endpoints (id, name, url, check_frequency, notify_to) VALUES ($1, $2, $3, $4, $5)",
		e.ID, e.Name, e.Url, e.CheckFrequency, e.NotifyTo)
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

	endpoints, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[entity.Endpoint])
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return endpoints, nil
}
