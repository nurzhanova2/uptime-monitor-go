package utils

import (
	"os"
	"gopkg.in/yaml.v3"
)

type Config struct {
	IntervalSeconds int `yaml:"interval_seconds"`
	TelegramToken   string `yaml:"telegram_token"`
	TelegramChatID  int64 `yaml:"telegram_chat_id"`
	Targets         []TargetConfig `yaml:"targets"`
}

type TargetConfig struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
