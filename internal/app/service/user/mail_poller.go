package service

import (
	"time"

	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/mail"
)

type MailPoller interface {
	Poll(string, string, time.Time, int64) (mail.PollResult, error)
}
