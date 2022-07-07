package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/mail"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/models"
)

type UserService struct {
	ur         UserRepository
	mr         MailboxRepository
	crypt      Cryptographer
	sender     MessageSender
	mailPoller MailPoller
}

func NewUserService(
	ur UserRepository,
	mr MailboxRepository,
	crypt Cryptographer,
	sender MessageSender,
	poller MailPoller,
) UserService {
	return UserService{
		ur:         ur,
		mr:         mr,
		crypt:      crypt,
		sender:     sender,
		mailPoller: poller,
	}
}

func (s UserService) Register(ctx context.Context, tgUsername string, tgChatID int64) (int64, error) {
	if tgUsername == "" {
		return 0, errors.New("invalid telegram username")
	}

	if tgChatID == 0 {
		return 0, errors.New("invalid telegram chat id")
	}

	return s.ur.Create(ctx, models.User{
		TgUsername: tgUsername,
		TgChatID:   tgChatID,
	})
}

func (s UserService) Get(ctx context.Context, tu string) (models.User, error) {
	return s.ur.GetByTgUsername(ctx, tu)
}

func (s UserService) DeleteByTelegramUsername(ctx context.Context, tu string) error {
	return s.ur.DeleteByTgUsername(ctx, tu)
}

func (s UserService) CreateUserMailbox(ctx context.Context, tu string, mb models.Mailbox) (int64, error) {
	if err := mail.ValidateMail(mb.Mail); err != nil {
		return 0, err
	}

	encrypted, err := s.crypt.Encrypt(mb.Password)
	if err != nil {
		return 0, err
	}

	mb.Password = encrypted
	return s.ur.CreateMailbox(ctx, tu, mb)
}

func (s UserService) DeleteUserMailbox(ctx context.Context, tu, mail string) error {
	return s.ur.DeleteMailbox(ctx, tu, mail)
}

func (s UserService) Notify(ctx context.Context, user models.User) []error {
	mailboxes, err := s.ur.GetMailboxes(ctx, user.TgUsername)
	if err != nil {
		return []error{fmt.Errorf("get user mailboxes: %w", err)}
	}

	if len(mailboxes) == 0 {
		return nil
	}

	errs := make([]error, 0)
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(mailboxes))

	for _, mb := range mailboxes {
		go func(tu string, mb models.Mailbox) {
			defer wg.Done()

			err = s.notify(ctx, user, mb)
			if err == nil {
				return
			}

			mx.Lock()
			errs = append(errs, err)
			mx.Unlock()
		}(user.TgUsername, mb)
	}

	wg.Wait()

	log.Println("poll errors:", errs)
	return errs
}

func (s UserService) notify(ctx context.Context, user models.User, mb models.Mailbox) error {
	decrypted, err := s.crypt.Decrypt(mb.Password)
	if err != nil {
		return fmt.Errorf("password decrypt of %q: %w", mb.Mail, err)
	}

	polledAt := time.Now()
	res, err := s.mailPoller.Poll(mb.Mail, decrypted, mb.PolledAt, mb.LastMessageID)
	if err != nil {
		_ = s.sender.SendMessage(
			user.TgChatID,
			fmt.Sprintf("❗ Unable to read %q messages, please check your credentials", mb.Mail),
		)
		return fmt.Errorf("poll %q mailbox: %w", mb.Mail, err)
	}

	// No new messages
	if len(res.Messages) == 0 && res.Failed == 0 {
		return nil
	}

	if err = s.sender.SendMessage(user.TgChatID, s.buildMessage(mb.Mail, res.Messages, res.Failed)); err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	mb.LastMessageID = res.LastMessageID
	mb.PolledAt = polledAt

	if err = s.mr.Update(ctx, mb); err != nil {
		return fmt.Errorf("update %q mailbox after poll: %w", mb.Mail, err)
	}

	return nil
}

func (s UserService) buildMessage(mail string, messages []mail.Message, failed int) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("📬 Mailbox: %s\n\n", mail))

	if len(messages) != 0 {
		formatted := make([]string, len(messages))
		for i, m := range messages {
			formatted[i] = fmt.Sprintf(
				"📩 %s\nFrom: %s Date: %s\n",
				m.Subject,
				m.From,
				m.Date.Format(time.RFC822),
			)
		}

		sb.WriteString(strings.Join(formatted, "\n"))
	}

	if failed > 0 {
		sb.WriteString(fmt.Sprintf("\n❗ Number of unread emails that could not be processed: %d", failed))
	}

	return sb.String()
}

func (s UserService) NotifyAll(ctx context.Context) []error {
	users, err := s.ur.GetList(ctx)
	if err != nil {
		return []error{fmt.Errorf("get users: %w", err)}
	}

	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(users))
	errs := make([]error, 0)

	for _, user := range users {
		go func(user models.User) {
			defer wg.Done()

			e := s.Notify(ctx, user)

			mx.Lock()
			errs = append(errs, e...)
			mx.Unlock()
		}(user)
	}

	wg.Wait()

	return errs
}
