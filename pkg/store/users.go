package store

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lbrictson/squint/pkg/api"
)

type UsersStore struct {
	conn *pgxpool.Pool // conn holds the already initialized db connection
}

func (u UsersStore) Create(ctx context.Context, input CreateUserInput) (*api.User, error) {
	_, err := u.conn.Exec(ctx, "INSERT INTO users(id, email, hashed_password, role, api_key) VALUES($1, $2, $3, $4, $5)",
		uuid.New().String(), input.Email, input.HashedPassword, input.Role, uuid.New().String())
	if err != nil {
		return nil, err
	}
	return u.GetByEmail(ctx, input.Email)
}

func (u UsersStore) Delete(ctx context.Context, id string) error {
	_, err := u.conn.Exec(ctx, `DELETE FROM users WHERE id = $1`, id)
	return err
}

func (u UsersStore) Update(ctx context.Context, user api.User) (*api.User, error) {
	_, err := u.conn.Exec(ctx, `UPDATE users SET (email, hashed_password, role, api_key, updated_at) = ($1, $2, $3, $4, $5) WHERE id = $6`,
		user.Email, user.HashedPassword, user.Role, user.APIKey, time.Now(), user.ID)
	if err != nil {
		return nil, err
	}
	return u.GetByID(ctx, user.ID)
}

func (u UsersStore) List(ctx context.Context, options UserListOptions) ([]api.User, error) {
	whereClause := ``
	if options.RoleEqual != nil {
		whereClause = whereClause + `WHERE role = '` + *options.RoleEqual + `'`
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
	users := []api.User{}
	rows, err := u.conn.Query(ctx, `SELECT id,email,hashed_password,role,api_key,created_at,updated_at FROM users `+whereClause)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := api.User{}
		err := rows.Scan(&user.ID, &user.Email, &user.HashedPassword, &user.Role, &user.APIKey, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u UsersStore) GetByID(ctx context.Context, id string) (*api.User, error) {
	row := u.conn.QueryRow(ctx, "select id,email,hashed_password,role,api_key,created_at,updated_at from users WHERE id = $1", id)
	user := api.User{}
	err := row.Scan(&user.ID, &user.Email, &user.HashedPassword, &user.Role, &user.APIKey, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (u UsersStore) GetByEmail(ctx context.Context, email string) (*api.User, error) {
	row := u.conn.QueryRow(ctx, "select id,email,hashed_password,role,api_key,created_at,updated_at from users WHERE email = $1", email)
	user := api.User{}
	err := row.Scan(&user.ID, &user.Email, &user.HashedPassword, &user.Role, &user.APIKey, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
