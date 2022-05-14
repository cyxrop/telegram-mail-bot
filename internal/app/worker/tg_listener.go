package worker

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/models"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/repository"
	service "gitlab.ozon.dev/cyxrop/homework-2/internal/app/service/user"
	"gitlab.ozon.dev/cyxrop/homework-2/pkg/telegram"
)

const (
	addMailboxCmd    = "/add_mailbox"
	removeMailboxCmd = "/remove_mailbox"
	purgeCmd         = "/purge"
	newMailsCmd      = "/new_mails"
	helpCmd          = "/help"

	helpDescription = `
add_mailbox - Add a mailbox that will be used to read mail (example: /add_mailbox mymail@gmail.com/secretpassword)
remove_mailbox - Remove the specified mailbox credentials (example: /remove_mailbox mymail@gmail.com)
purge - Delete all user data
new_mails - Receive new unread messages from added mailboxes
help - Description of commands
`
)

type TgListener struct {
	bot *telegram.Bot

	us service.UserService
}

func NewTgListener(bot *telegram.Bot, us service.UserService) TgListener {
	return TgListener{
		bot: bot,
		us:  us,
	}
}

func (l TgListener) Run(ctx context.Context) {
	go func() {
		updates := l.bot.Updates()
		for {
			select {
			case <-ctx.Done():
				log.Println("tg listener ctx done")
				return
			case update := <-updates:
				if update.Message == nil || !update.Message.IsCommand() {
					continue
				}

				chatID := update.Message.Chat.ID
				tu := update.Message.From.UserName
				text := update.Message.Text

				var (
					err error
					cmd string
					msg string
				)

				switch true {
				case strings.HasPrefix(text, addMailboxCmd):
					cmd = addMailboxCmd
					err = l.handleAddMailbox(ctx, tu, chatID, strings.TrimLeft(text, addMailboxCmd))

				case strings.HasPrefix(text, removeMailboxCmd):
					cmd = removeMailboxCmd
					err = l.handleRemoveMailbox(ctx, tu, strings.TrimLeft(text, removeMailboxCmd))

				case strings.HasPrefix(text, purgeCmd):
					cmd = purgeCmd
					err = l.handlePurge(ctx, tu)

				case strings.HasPrefix(text, newMailsCmd):
					cmd = newMailsCmd
					err = l.handleNewMails(ctx, tu)

				case strings.HasPrefix(text, helpCmd):
					_ = l.bot.SendMessage(chatID, helpDescription)
					continue

				default:
					err = errors.New("unknown command")
				}

				if err != nil {
					msg = fmt.Sprintf("❗ %s", err)
				} else {
					msg = fmt.Sprintf("✅ %s done.", cmd)
				}

				_ = l.bot.SendMessage(chatID, msg)
			}
		}
	}()
}

func (l TgListener) handleAddMailbox(ctx context.Context, tu string, chatID int64, args string) error {
	_, err := l.us.Get(ctx, tu)
	if err != nil {
		if err != repository.ErrNotFound {
			return fmt.Errorf("%s failed: get user %q: %w", addMailboxCmd, tu, err)
		}

		// Register user when creating first mailbox
		_, err = l.us.Register(ctx, tu, chatID)
		if err != nil {
			return fmt.Errorf("%s failed: create user %q: %w\n", addMailboxCmd, tu, err)
		}
	}

	split := strings.SplitN(args, "/", 2)
	if len(split) != 2 {
		return fmt.Errorf("%s failed: invalid arguments, please see /help\n", addMailboxCmd)
	}

	_, err = l.us.CreateUserMailbox(ctx, tu, models.Mailbox{
		Mail:     strings.TrimSpace(split[0]),
		Password: strings.TrimSpace(split[1]),
	})
	if errors.Is(err, repository.ErrUniqueViolation) {
		return fmt.Errorf("%s failed: mailbox with the specified mail already exists\n", addMailboxCmd)
	}
	if err != nil {
		return fmt.Errorf("%s failed: create mailbox: %w\n", addMailboxCmd, err)
	}

	return nil
}

func (l TgListener) handleRemoveMailbox(ctx context.Context, tu string, mail string) error {
	mail = strings.TrimSpace(mail)
	err := l.us.DeleteUserMailbox(ctx, tu, mail)
	if errors.Is(err, repository.ErrNotFound) {
		return fmt.Errorf("%s failed: mailbox %q not found", removeMailboxCmd, mail)
	}

	if err != nil {
		return fmt.Errorf("%s failed: %w\n", removeMailboxCmd, err)
	}

	return nil
}

func (l TgListener) handlePurge(ctx context.Context, tu string) error {
	err := l.us.DeleteByTelegramUsername(ctx, tu)
	if err != nil {
		return fmt.Errorf("%s failed: %w\n", purgeCmd, err)
	}

	return nil
}

func (l TgListener) handleNewMails(ctx context.Context, tu string) error {
	user, err := l.us.Get(ctx, tu)
	if err != nil {
		return fmt.Errorf("%s failed: get user %q: %w\n", newMailsCmd, tu, err)
	}

	errs := l.us.Notify(ctx, user)
	if len(errs) != 0 {
		return fmt.Errorf("%s failed: notify user %q: %s\n", newMailsCmd, tu, errs)
	}

	return nil
}
