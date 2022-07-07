package service

import (
	"context"

	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/models"
)

type UserRepository interface {
	Create(context.Context, models.User) (int64, error)
	GetByTgUsername(context.Context, string) (models.User, error)
	GetList(context.Context) ([]models.User, error)
	DeleteByTgUsername(context.Context, string) error

	CreateMailbox(context.Context, string, models.Mailbox) (int64, error)
	DeleteMailbox(context.Context, string, string) error
	GetMailboxes(context.Context, string) ([]models.Mailbox, error)
}
