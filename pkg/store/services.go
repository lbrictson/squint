package store

import (
	"context"
	"fmt"
	"time"

	"github.com/lib/pq"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lbrictson/squint/pkg/api"
)

type ServicesStore struct {
	conn *pgxpool.Pool // conn holds the already initialized db connection
}

func (s ServicesStore) Create(ctx context.Context, input CreateServiceInput) (*api.Service, error) {
	id := uuid.New().String()
	status := "Operational"
	_, err := s.conn.Exec(ctx, "INSERT INTO services(id, name, description, status, group_id) VALUES($1, $2, $3, $4, $5)",
		id, input.Name, input.Description, status, input.Group)
	if err != nil {
		return nil, err
	}
	return s.GetByID(ctx, id)
}

func (s ServicesStore) Delete(ctx context.Context, id string) error {
	_, err := s.conn.Exec(ctx, `DELETE FROM services WHERE id = $1`, id)
	return err
}

func (s ServicesStore) Update(ctx context.Context, service api.Service) (*api.Service, error) {
	if service.Pages == nil {
		service.Pages = []string{}
	}
	_, err := s.conn.Exec(ctx, `UPDATE services SET (name, description, status, status_since, pages, updated_at, group_id) = ($1, $2, $3, $4, $5, $6, $7) WHERE id = $8`,
		service.Name, service.Description, service.Status, service.StatusSince, pq.Array(service.Pages), time.Now(), service.Group, service.ID)
	if err != nil {
		return nil, err
	}
	return s.GetByID(ctx, service.ID)
}

func (s ServicesStore) List(ctx context.Context, options ServiceListOptions) ([]api.Service, error) {
	whereClause := ``
	if options.GroupEqual != nil {
		whereClause = whereClause + `WHERE group_id = '` + *options.GroupEqual + `'`
	}
	if options.StatusEqual != nil {
		if whereClause == `` {
			whereClause = whereClause + `WHERE status = '` + *options.StatusEqual + `'`
		} else {
			whereClause = whereClause + `AND status = '` + *options.StatusEqual + `'`
		}
	}
	if options.PageEqual != nil {
		if whereClause == `` {
			whereClause = whereClause + `WHERE '` + *options.PageEqual + `' = ANY (pages)`
		} else {
			whereClause = whereClause + `AND '` + *options.PageEqual + `' = ANY (pages)`
		}
	}
	if options.OrderBy != nil {
		whereClause = whereClause + ` ORDER BY ` + *options.OrderBy
	}
	if options.Limit != nil {
		whereClause = whereClause + ` LIMIT ` + fmt.Sprintf("%v", *options.Limit)
	}
	if options.OffSet != nil {
		whereClause = whereClause + ` OFFSET ` + fmt.Sprintf("%v", *options.OffSet)
	}
	services := []api.Service{}
	rows, err := s.conn.Query(ctx, `SELECT id,name,description,status,status_since,group_id,pages,created_at,updated_at FROM services `+whereClause)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		service := api.Service{}
		err := rows.Scan(&service.ID, &service.Name, &service.Description, &service.Status, &service.StatusSince, &service.Group, pq.Array(&service.Pages), &service.CreatedAt, &service.UpdatedAt)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func (s ServicesStore) GetByID(ctx context.Context, id string) (*api.Service, error) {
	row := s.conn.QueryRow(ctx, "select id,name,description,status,status_since,group_id,pages,created_at,updated_at from services WHERE id = $1", id)
	service := api.Service{}
	err := row.Scan(&service.ID, &service.Name, &service.Description, &service.Status, &service.StatusSince, &service.Group, pq.Array(&service.Pages), &service.CreatedAt, &service.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &service, nil
}
