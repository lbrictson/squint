package store

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lbrictson/squint/pkg/api"
)

// Stores hold various store objects
type Stores struct {
	Users    UsersStorage
	Groups   GroupsStorage
	Services ServicesStorage
}

func New(dbConn *pgxpool.Pool) (*Stores, error) {
	return &Stores{
		Users:    UsersStore{conn: dbConn},
		Groups:   GroupsStore{conn: dbConn},
		Services: ServicesStore{conn: dbConn},
	}, nil
}

type UserListOptions struct {
	Limit     *int
	OffSet    *int
	OrderBy   *string
	RoleEqual *string
}

type CreateUserInput struct {
	Email          string
	HashedPassword string
	Role           string
}

type UsersStorage interface {
	Create(ctx context.Context, input CreateUserInput) (*api.User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, user api.User) (*api.User, error)
	List(ctx context.Context, options UserListOptions) ([]api.User, error)
	GetByID(ctx context.Context, id string) (*api.User, error)
	GetByEmail(ctx context.Context, email string) (*api.User, error)
}

type CreateGroupInput struct {
	Name        string
	Description string
	Expanded    bool
}

type GroupListOptions struct {
	Limit   *int
	OffSet  *int
	OrderBy *string
}

type GroupsStorage interface {
	Create(ctx context.Context, input CreateGroupInput) (*api.Group, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, group api.Group) (*api.Group, error)
	List(ctx context.Context, options GroupListOptions) ([]api.Group, error)
	GetByID(ctx context.Context, id string) (*api.Group, error)
}

type CreateServiceInput struct {
	Name        string
	Description string
	Group       string
}

type ServiceListOptions struct {
	Limit       *int
	OffSet      *int
	OrderBy     *string
	StatusEqual *string
	GroupEqual  *string
	PageEqual   *string
}

type ServicesStorage interface {
	Create(ctx context.Context, input CreateServiceInput) (*api.Service, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, service api.Service) (*api.Service, error)
	List(ctx context.Context, options ServiceListOptions) ([]api.Service, error)
	GetByID(ctx context.Context, id string) (*api.Service, error)
}
