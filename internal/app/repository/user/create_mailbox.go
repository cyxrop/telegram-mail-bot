package repository

import (
	"context"

	"github.com/jackc/pgerrcode"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/models"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/repository"
)

//CreateMailbox creates a user's mailbox.
func (r UserRepository) CreateMailbox(ctx context.Context, tu string, mailbox models.Mailbox) (int64, error) {
	const query = `
		insert into mailboxes (
			mail,
			password,
			last_message_id,
			user_id,
			polled_at
		) VALUES (
			$1, $2, $3, (select id from users where tg_username = $4), now()
		) returning id
	`

	var ID int64
	err := r.pool.QueryRow(ctx, query,
		mailbox.Mail,
		mailbox.Password,
		mailbox.LastMessageID,
		tu,
	).Scan(&ID)
	if repository.ErrorIs(err, pgerrcode.UniqueViolation) {
		// If a mailbox with the specified mail already exists.
		return 0, repository.ErrUniqueViolation
	}

	if repository.ErrorIs(err, pgerrcode.ForeignKeyViolation) {
		// If user with specified id does not exist.
		return 0, repository.ErrForeignKeyViolation
	}

	if repository.ErrorIs(err, pgerrcode.NotNullViolation) {
		// If user with specified telegram username does not exist.
		return 0, repository.ErrNotNullViolation
	}

	return ID, err
}
