package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(p *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: p,
	}
}
