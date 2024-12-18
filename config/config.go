package config

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"log"
	"os"
)

type Config struct {
	UserAgent string `validate:"required"`
	Proxy     struct {
		Host string `validate:"required"`
		Port string `validate:"required"`
	} `validate:"required"`
	URLs    []string `validate:"required"`
	Timeout string   `validate:"required"`
	Delay   string   `validate:"required"`
	Deps    struct {
		Habr    Habr    `validate:"required"`
		YouTube YouTube `validate:"required"`
	} `validate:"required"`
}

type Habr struct {
	URLs []string `validate:"required"`
}

type YouTube struct {
	URLs []string `validate:"required"`
}

func Load(path string) Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("open config file fail: %v\n", err)
	}

	cfg := Config{}

	if err = json.NewDecoder(file).Decode(&cfg); err != nil {
		log.Fatalf("decode config fail: %v\n", err)
	}

	if err = validator.New().Struct(&cfg); err != nil {
		log.Fatalf("validate config fail: %v\n", err)
	}

	return cfg
}
