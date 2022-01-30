package store

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lbrictson/squint/pkg/api"
)

type GroupsStore struct {
	conn *pgxpool.Pool // conn holds the already initialized db connection
}

func (g GroupsStore) Create(ctx context.Context, input CreateGroupInput) (*api.Group, error) {
	id := uuid.New().String()
	_, err := g.conn.Exec(ctx, "INSERT INTO groups(id, name, description, expanded) VALUES($1, $2, $3, $4)",
		id, input.Name, input.Description, input.Expanded)
	if err != nil {
		return nil, err
	}
	return g.GetByID(ctx, id)
}

func (g GroupsStore) Delete(ctx context.Context, id string) error {
	_, err := g.conn.Exec(ctx, `DELETE FROM groups WHERE id = $1`, id)
	return err
}

func (g GroupsStore) Update(ctx context.Context, group api.Group) (*api.Group, error) {
	_, err := g.conn.Exec(ctx, `UPDATE groups SET (name, description, expanded, updated_at) = ($1, $2, $3, $4) WHERE id = $5`,
		group.Name, group.Description, group.Expanded, time.Now(), group.ID)
	if err != nil {
		return nil, err
	}
	return g.GetByID(ctx, group.ID)
}

func (g GroupsStore) List(ctx context.Context, options GroupListOptions) ([]api.Group, error) {
	whereClause := ``
	if options.OrderBy != nil {
		whereClause = whereClause + ` ORDER BY ` + *options.OrderBy
	}
	if options.Limit != nil {
		whereClause = whereClause + ` LIMIT ` + fmt.Sprintf("%v", *options.Limit)
	}
	if options.OffSet != nil {
		whereClause = whereClause + ` OFFSET ` + fmt.Sprintf("%v", *options.OffSet)
	}
	groups := []api.Group{}
	rows, err := g.conn.Query(ctx, `SELECT id,name,description,expanded,created_at,updated_at FROM groups `+whereClause)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		group := api.Group{}
		err := rows.Scan(&group.ID, &group.Name, &group.Description, &group.Expanded, &group.CreatedAt, &group.UpdatedAt)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func (g GroupsStore) GetByID(ctx context.Context, id string) (*api.Group, error) {
	row := g.conn.QueryRow(ctx, "SELECT id,name,description,expanded,created_at,updated_at FROM groups WHERE id = $1", id)
	group := api.Group{}
	err := row.Scan(&group.ID, &group.Name, &group.Description, &group.Expanded, &group.CreatedAt, &group.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &group, nil
}
