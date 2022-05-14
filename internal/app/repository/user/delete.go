package repository

import (
	"context"

	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/repository"
)

func (r UserRepository) Delete(ctx context.Context, ID int64) error {
	const query = `
		delete from users
		where id = $1;
	`

	res, err := r.pool.Exec(ctx, query, ID)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return repository.ErrNotFound
	}

	return nil
}

func (r UserRepository) DeleteByTgUsername(ctx context.Context, tu string) error {
	const query = `
		delete from users
		where tg_username = $1;
	`

	res, err := r.pool.Exec(ctx, query, tu)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return repository.ErrNotFound
	}

	return nil
}
