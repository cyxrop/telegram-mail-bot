package repository

import (
	"context"

	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/models"
)

func (r UserRepository) GetMailboxes(ctx context.Context, tu string) ([]models.Mailbox, error) {
	const query = `
		select id,
			mail,
			password,
			user_id,
			last_message_id,
			polled_at
		from mailboxes
		where user_id = (select id from users where tg_username = $1);
	`

	rows, err := r.pool.Query(ctx, query, tu)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mailboxes := make([]models.Mailbox, 0)
	for rows.Next() {
		var mb models.Mailbox
		if err = rows.Scan(
			&mb.ID,
			&mb.Mail,
			&mb.Password,
			&mb.UserID,
			&mb.LastMessageID,
			&mb.PolledAt,
		); err != nil {
			return nil, err
		}

		mailboxes = append(mailboxes, mb)
	}

	return mailboxes, nil
}
