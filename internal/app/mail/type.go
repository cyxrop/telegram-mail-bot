package mail

import (
	"fmt"
	gomail "net/mail"
	"strings"
)

type Type int16

const (
	TypeUnknown Type = iota
	TypeGoogle
	TypeMailRu
	TypeYandex
)

func ParseType(mail string) Type {
	switch true {
	case strings.HasSuffix(mail, "@gmail.com"):
		return TypeGoogle
	case strings.HasSuffix(mail, "@mail.ru"):
		return TypeMailRu
	case strings.HasSuffix(mail, "@yandex.ru"):
		return TypeYandex
	}

	return TypeUnknown
}

func ValidateMail(mail string) error {
	_, err := gomail.ParseAddress(mail)
	if err != nil {
		return fmt.Errorf("invalid mail: %s", mail)
	}

	if ParseType(mail) == TypeUnknown {
		return fmt.Errorf("mail is not supported: %s", mail)
	}

	return nil
}
