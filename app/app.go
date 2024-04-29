package app

import (
	"fiber-api/app/config"
	"fiber-api/app/cron"
	"fiber-api/app/ctx"
	"fiber-api/app/http"
)

type App struct {
	ctx   *ctx.Ctx
	serve *http.Serve
	cron  *cron.Cron
}

func New(cfg *config.Config) *App {
	return &App{
		ctx: ctx.New(cfg),
	}
}

func (app *App) Run() (err error) {
	if err = app.ctx.Init(); err != nil {
		return
	}

	app.init()

	if err = app.cron.Run(); err != nil {
		return
	}

	return app.serve.Run()
}

func (app *App) init() {
	app.serve = http.New(app.ctx)
	app.cron = cron.New(app.ctx)
}
