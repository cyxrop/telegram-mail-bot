package service

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/mail"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/models"
)

func TestRegister(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		mc := minimock.NewController(t)
		defer mc.Finish()

		ctx := context.Background()
		user := models.User{
			TgUsername: "test",
			TgChatID:   123,
		}

		expectedID := int64(10)
		ur := NewUserRepositoryMock(mc).
			CreateMock.
			Expect(ctx, user).
			Return(expectedID, nil)

		mr := NewMailboxRepositoryMock(mc)
		crypt := NewCryptographerMock(mc)
		sender := NewMessageSenderMock(mc)
		mailPoller := NewMailPollerMock(mc)

		s := NewUserService(ur, mr, crypt, sender, mailPoller)
		ID, err := s.Register(ctx, user.TgUsername, user.TgChatID)
		require.NoError(t, err)
		assert.Equal(t, expectedID, ID)
	})

	t.Run("negative", func(t *testing.T) {
		testCases := map[string]struct {
			username string
			chatID   int64
		}{
			"invalid telegram username": {chatID: 123},
			"invalid telegram chat id":  {username: "valid"},
		}

		for message, data := range testCases {
			t.Run(message, func(t *testing.T) {
				mc := minimock.NewController(t)
				defer mc.Finish()

				ctx := context.Background()
				ur := NewUserRepositoryMock(mc)
				mr := NewMailboxRepositoryMock(mc)
				crypt := NewCryptographerMock(mc)
				sender := NewMessageSenderMock(mc)
				mailPoller := NewMailPollerMock(mc)

				s := NewUserService(ur, mr, crypt, sender, mailPoller)
				_, err := s.Register(ctx, data.username, data.chatID)

				assert.Errorf(t, err, message)
			})
		}
	})
}

func TestGet(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	ctx := context.Background()
	expected := models.User{TgUsername: "username"}

	ur := NewUserRepositoryMock(mc).
		GetByTgUsernameMock.
		Expect(ctx, expected.TgUsername).
		Return(expected, nil)

	mr := NewMailboxRepositoryMock(mc)
	crypt := NewCryptographerMock(mc)
	sender := NewMessageSenderMock(mc)
	mailPoller := NewMailPollerMock(mc)

	s := NewUserService(ur, mr, crypt, sender, mailPoller)
	actual, err := s.Get(ctx, expected.TgUsername)

	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestDeleteByTelegramUsername(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	ctx := context.Background()
	expected := "username"

	ur := NewUserRepositoryMock(mc).
		DeleteByTgUsernameMock.
		Expect(ctx, expected).
		Return(nil)

	mr := NewMailboxRepositoryMock(mc)
	crypt := NewCryptographerMock(mc)
	sender := NewMessageSenderMock(mc)
	mailPoller := NewMailPollerMock(mc)

	s := NewUserService(ur, mr, crypt, sender, mailPoller)
	err := s.DeleteByTelegramUsername(ctx, expected)

	require.NoError(t, err)
}

func TestCreateUserMailbox(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	ctx := context.Background()
	expectedID := int64(55)
	username := "username"
	mb := models.Mailbox{
		Mail:     "test@gmail.com",
		Password: "secret",
	}

	crypt := NewCryptographerMock(mc).
		EncryptMock.
		Expect(mb.Password).
		Return("encrypted", nil)

	ur := NewUserRepositoryMock(mc).
		CreateMailboxMock.
		Expect(ctx, username, models.Mailbox{
			Mail:     mb.Mail,
			Password: "encrypted",
		}).
		Return(expectedID, nil)

	mr := NewMailboxRepositoryMock(mc)
	sender := NewMessageSenderMock(mc)
	mailPoller := NewMailPollerMock(mc)

	s := NewUserService(ur, mr, crypt, sender, mailPoller)
	ID, err := s.CreateUserMailbox(ctx, username, mb)

	require.NoError(t, err)
	assert.Equal(t, expectedID, ID)
}

func TestDeleteUserMailbox(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	ctx := context.Background()
	username, m := "username", "test@test.com"

	ur := NewUserRepositoryMock(mc).
		DeleteMailboxMock.
		Expect(ctx, username, m).
		Return(nil)

	mr := NewMailboxRepositoryMock(mc)
	crypt := NewCryptographerMock(mc)
	sender := NewMessageSenderMock(mc)
	mailPoller := NewMailPollerMock(mc)

	s := NewUserService(ur, mr, crypt, sender, mailPoller)
	err := s.DeleteUserMailbox(ctx, username, m)
	require.NoError(t, err)
}

func TestNotify(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	ctx := context.Background()
	user := models.User{
		TgUsername: "username",
		TgChatID:   95,
	}

	mbs := []models.Mailbox{
		{
			Mail:          "mail1@test.com",
			Password:      "secret1",
			PolledAt:      time.Now().Add(time.Hour * -1),
			LastMessageID: 10,
		},
		{
			Mail:          "mail2@test.com",
			Password:      "secret2",
			PolledAt:      time.Now().Add(time.Hour * -5),
			LastMessageID: 15,
		},
		{
			Mail:          "mail3@test.com",
			Password:      "secret3",
			PolledAt:      time.Now().Add(time.Hour * -8),
			LastMessageID: 20,
		},
	}

	mb1Messages := []mail.Message{
		{
			From:    "from1@gmail.com",
			Date:    time.Now().Add(time.Hour * -3),
			Subject: "subject1",
		},
		{
			From:    "from2@gmail.com",
			Date:    time.Now().Add(time.Hour * -2),
			Subject: "subject2",
		},
	}
	mb2Messages := []mail.Message{
		{
			From:    "from3@gmail.com",
			Date:    time.Now().Add(time.Hour * -1),
			Subject: "subject3",
		},
	}

	ur := NewUserRepositoryMock(mc).
		GetMailboxesMock.
		Expect(ctx, user.TgUsername).
		Return(mbs, nil)

	crypt := NewCryptographerMock(mc).
		DecryptMock.When("secret1").Then("decrypted1", nil).
		DecryptMock.When("secret2").Then("decrypted2", nil).
		DecryptMock.When("secret3").Then("", errors.New("fake error"))

	mailPoller := NewMailPollerMock(mc).
		PollMock.
		When(mbs[0].Mail, "decrypted1", mbs[0].PolledAt, mbs[0].LastMessageID).
		Then(mail.PollResult{
			Messages:      mb1Messages,
			Failed:        2,
			LastMessageID: 51,
		}, nil).
		PollMock.
		When(mbs[1].Mail, "decrypted2", mbs[1].PolledAt, mbs[1].LastMessageID).
		Then(mail.PollResult{
			Messages:      mb2Messages,
			LastMessageID: 52,
		}, nil)

	sender := NewMessageSenderMock(mc).
		SendMessageMock.
		When(user.TgChatID, fmt.Sprintf(
			`📬 Mailbox: %s

📩 %s
From: %s Date: %s

📩 %s
From: %s Date: %s

❗ Number of unread emails that could not be processed: 2`,
			mbs[0].Mail,
			mb1Messages[0].Subject,
			mb1Messages[0].From,
			mb1Messages[0].Date.Format(time.RFC822),
			mb1Messages[1].Subject,
			mb1Messages[1].From,
			mb1Messages[1].Date.Format(time.RFC822),
		)).
		Then(nil).
		SendMessageMock.
		When(user.TgChatID, fmt.Sprintf(`📬 Mailbox: %s

📩 %s
From: %s Date: %s
`, mbs[1].Mail, mb2Messages[0].Subject, mb2Messages[0].From, mb2Messages[0].Date.Format(time.RFC822))).
		Then(nil)

	mr := NewMailboxRepositoryMock(mc).
		UpdateMock.
		Inspect(func(ctx context.Context, m1 models.Mailbox) {
			switch m1.Mail {
			case mbs[0].Mail:
				assert.Equal(t, mbs[0].Password, m1.Password)
				assert.Equal(t, int64(51), m1.LastMessageID)
			case mbs[1].Mail:
				assert.Equal(t, mbs[1].Password, m1.Password)
				assert.Equal(t, int64(52), m1.LastMessageID)
			default:
				assert.Fail(t, fmt.Sprintf("unexpected mail: %s", m1.Mail))
			}
		}).
		Return(nil)

	s := NewUserService(ur, mr, crypt, sender, mailPoller)
	errs := s.Notify(ctx, user)

	assert.Len(t, errs, 1)
	assert.Errorf(t, errs[0], fmt.Sprintf("password decrypt of %q: fake error", mbs[2].Mail))
}
