package service

type MessageSender interface {
	SendMessage(int64, string) error
}
