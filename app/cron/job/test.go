package job

import (
	"fiber-api/app/ctx"

	"github.com/rs/zerolog/log"
)

type Test struct{}

func (j *Test) Spec() string {
	// return "*/5 * * * * *"
	return "@every 5s"
}

func (j *Test) Run(ctx *ctx.Ctx) {
	log.Info().Msgf("I'm cron job, debug: %t", ctx.Config.Debug)
}
