package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/models"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/repository"
)

func (r UserRepository) Get(ctx context.Context, ID int64) (models.User, error) {
	const query = `
		select id,
			tg_username,
			tg_chat_id,
			created_at
		from users
		where id = $1;
	`

	var user models.User
	err := r.pool.QueryRow(ctx, query, ID).Scan(
		&user.ID,
		&user.TgUsername,
		&user.TgChatID,
		&user.CreatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return user, repository.ErrNotFound
	}

	return user, err
}

func (r UserRepository) GetByTgUsername(ctx context.Context, tu string) (models.User, error) {
	const query = `
		select id,
			tg_username,
			tg_chat_id,
			created_at
		from users
		where tg_username = $1;
	`

	var user models.User
	err := r.pool.QueryRow(ctx, query, tu).Scan(
		&user.ID,
		&user.TgUsername,
		&user.TgChatID,
		&user.CreatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return user, repository.ErrNotFound
	}

	return user, err
}

func (r UserRepository) GetList(ctx context.Context) ([]models.User, error) {
	const query = `
		select id,
			tg_username,
			tg_chat_id,
			created_at
		from users;
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		var u models.User
		if err = rows.Scan(
			&u.ID,
			&u.TgUsername,
			&u.TgChatID,
			&u.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}
