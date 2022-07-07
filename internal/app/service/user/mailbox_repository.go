package service

import (
	"context"

	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/models"
)

type MailboxRepository interface {
	Update(context.Context, models.Mailbox) error
}
