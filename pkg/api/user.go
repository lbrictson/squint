package api

import "time"

type User struct {
	ID             string
	Username       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	HashedPassword string
	Role           string
	APIKey         string
}
