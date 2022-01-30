package api

import (
	"time"
)

type Group struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Expanded    bool
}

type Service struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      string
	StatusSince time.Time
	Pages       []string
	Group       string
}
