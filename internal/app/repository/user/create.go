package repository

import (
	"context"

	"github.com/jackc/pgerrcode"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/models"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/repository"
)

func (r UserRepository) Create(ctx context.Context, user models.User) (int64, error) {
	const query = `
		insert into users (
			tg_username,
			tg_chat_id,
			created_at
		) VALUES (
			$1, $2, now()
		) returning id
	`

	var ID int64
	err := r.pool.QueryRow(ctx, query, user.TgUsername, user.TgChatID).Scan(&ID)
	if repository.ErrorIs(err, pgerrcode.UniqueViolation) {
		// If a user with the specified telegram username already exists.
		return 0, repository.ErrUniqueViolation
	}

	return ID, err
}
