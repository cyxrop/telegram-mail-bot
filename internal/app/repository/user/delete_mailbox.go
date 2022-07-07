package repository

import (
	"context"

	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/repository"
)

//DeleteMailbox delete the user's mailbox.
func (r UserRepository) DeleteMailbox(ctx context.Context, tu, mail string) error {
	const query = `
		delete from mailboxes
		where mail = $1
		  and user_id = (select id from users where tg_username = $2)
	`

	res, err := r.pool.Exec(ctx, query, mail, tu)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return repository.ErrNotFound
	}

	return nil
}
