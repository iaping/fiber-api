package job

import (
	"errors"
	"fiber-api/app/ctx"

	"github.com/rs/zerolog/log"
)

type Test struct{}

func (j *Test) Name() string {
	return "test"
}

func (j *Test) Spec() string {
	// return "*/5 * * * * *"
	return "@every 1m"
}

func (j *Test) Run(ctx *ctx.Ctx) error {
	log.Info().Msgf("I'm cron job, debug: %t", ctx.Cfg.Debug)
	return errors.New("test error")
}
