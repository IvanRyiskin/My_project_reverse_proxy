package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

// Config теги yaml указывают на соответствующие поля в yaml конфиге из-за разницы в регистре
type Config struct {
	Listener struct {
		Addr string `yaml:"addr"`
	} `yaml:"listener"`
	Cache struct {
		Size int `yaml:"size"`
	} `yaml:"cache"`
}

func loadConfig(path string) (*Config, error) {
	var config Config

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal разбирает yaml в структуру по указанным тегам
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	err = checkConfig(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// checkConfig проверяет конфиг на корректность
func checkConfig(config *Config) error {
	addr := strings.TrimSpace(config.Listener.Addr)
	if addr == "" {
		return errors.New("передан пустой адрес")
	}

	if config.Cache.Size <= 0 {
		return errors.New("передан некорректный размер кэша")
	}
	return nil
}
