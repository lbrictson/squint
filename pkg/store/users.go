package store

import "database/sql"

type UsersStore struct {
	conn *sql.DB // conn holds the already initialized db connection
}
