package config

import (
	"gopkg.in/yaml.v3"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	App struct {
		GRPC struct {
			Port int `yaml:"port" validate:"gt=0"`
		} `yaml:"grpc"`
		HTTP struct {
			Port int `yaml:"port" validate:"gt=0"`
		} `yaml:"http"`
	} `yaml:"app"`
	Imap struct {
		Hosts struct {
			Google string `yaml:"google" validate:"required"`
			MailRu string `yaml:"mail_ru" validate:"required"`
			Yandex string `yaml:"yandex" validate:"required"`
		} `yaml:"hosts"`
	} `yaml:"imap"`
	Telegram struct {
		Timeout int  `yaml:"timeout" validate:"gt=0"`
		Debug   bool `yaml:"debug"`
	} `yaml:"telegram"`
	Notifier struct {
		Every int `yaml:"every" validate:"gt=0"`
	} `yaml:"notifier"`
}

func Parse(bytes []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal(bytes, &cfg); err != nil {
		return nil, err
	}

	if err := validator.New().Struct(cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
