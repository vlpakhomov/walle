package app

import (
	"github.com/rs/zerolog"
	"github.com/tebeka/selenium"
	rand2 "math/rand/v2"
	"time"
	"walle/config"
)

type Dep interface {
	Exec() error
}

type app struct {
	cfg       *config.Config
	logger    *zerolog.Logger
	webDriver selenium.WebDriver
	deps      []Dep
}

func New(cfg *config.Config, logger *zerolog.Logger, webDriver selenium.WebDriver, deps []Dep) *app {
	return &app{
		cfg:       cfg,
		logger:    logger,
		webDriver: webDriver,
		deps:      deps,
	}
}

func (a *app) Start() {
	ticker := time.NewTicker(time.Second * 10)

	for range ticker.C {
		if err := a.deps[rand2.IntN(len(a.deps))].Exec(); err != nil {
			a.logger.Err(err).Msg("")
		}

		time.Sleep(time.Hour)
	}
}

func (a *app) Stop() {

}
