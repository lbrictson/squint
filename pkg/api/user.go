package api

import "time"

type User struct {
	ID             string
	Email          string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	HashedPassword string
	Role           string
	APIKey         string
}
