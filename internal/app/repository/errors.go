package repository

import (
	"errors"

	"github.com/jackc/pgconn"
)

var (
	ErrNotFound            = errors.New("not found")
	ErrUniqueViolation     = errors.New("uniqueness violation")
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrNotNullViolation    = errors.New("not null violation")
)

func ErrorIs(err error, code string) bool {
	if err == nil {
		return false
	}

	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return false
	}

	return pgErr.Code == code
}
