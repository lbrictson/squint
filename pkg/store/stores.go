package store

import (
	"context"
	"database/sql"
	"github.com/lbrictson/squint/pkg/api"
)

// Stores hold various store objects
type Stores struct {
}

func New(dbConn *sql.DB) (*Stores, error) {
	return &Stores{}, nil
}

type UserListOptions struct {
	Limit         *int
	OffSet        *int
	OrderBy       *string
	RoleEqual     *string
	UsernameEqual *string
}

type CreateUserInput struct {
	Username       string
	HashedPassword string
	Role           string
}

type UsersStorage interface {
	Create(ctx context.Context, input CreateUserInput) (*api.User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, user api.User) (*api.User, error)
	List(ctx context.Context, options UserListOptions) ([]api.User, error)
}
