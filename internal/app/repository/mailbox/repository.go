package repository

import "github.com/jackc/pgx/v4/pgxpool"

type MailboxRepository struct {
	pool *pgxpool.Pool
}

func NewMailboxRepository(p *pgxpool.Pool) *MailboxRepository {
	return &MailboxRepository{
		pool: p,
	}
}
