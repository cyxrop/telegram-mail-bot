package models

import "time"

type Mailbox struct {
	ID            int64
	Mail          string
	Password      string
	UserID        int64
	LastMessageID int64
	PolledAt      time.Time
}
