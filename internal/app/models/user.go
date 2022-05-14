package models

import "time"

type User struct {
	ID         int64
	TgUsername string
	TgChatID   int64
	CreatedAt  time.Time
}
