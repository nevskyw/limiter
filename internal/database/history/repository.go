package history

import "limiter/internal/database"

// Repository represents history table repo
type Repository interface{}

type repository struct {
	database.Pool
}

// New returns new Repository interface
func New(pool database.Pool) Repository {
	return &repository{pool}
}
