package repository

import (
	"context"

	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/models"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/repository"
)

func (r MailboxRepository) Update(ctx context.Context, mailbox models.Mailbox) error {
	const query = `
		update mailboxes
		set
			mail = $2,
			password = $3,
			user_id = $4,
			last_message_id = $5,
			polled_at = $6
		where id = $1;
	`

	res, err := r.pool.Exec(ctx, query,
		mailbox.ID,
		mailbox.Mail,
		mailbox.Password,
		mailbox.UserID,
		mailbox.LastMessageID,
		mailbox.PolledAt,
	)
	if res.RowsAffected() == 0 {
		return repository.ErrNotFound
	}

	return err
}
