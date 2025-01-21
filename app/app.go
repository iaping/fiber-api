package app

import (
	"fiber-api/app/cron"
	"fiber-api/app/ctx"
	"fiber-api/app/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.DateTime,
	})
}

type App struct {
	ctx   *ctx.Ctx
	serve *http.Serve
	cron  *cron.Cron
}

func New(ctx *ctx.Ctx) *App {
	return &App{
		ctx:   ctx,
		serve: http.New(ctx),
		cron:  cron.New(ctx),
	}
}

func (app *App) Run() (err error) {
	l := log.Info()
	l.Bool("Debug", app.ctx.Cfg.Debug)
	l.Msg("App")

	if err = app.cron.Run(); err != nil {
		return
	}

	return app.serve.Run()
}
