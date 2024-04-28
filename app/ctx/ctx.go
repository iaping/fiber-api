package ctx

import (
	"fiber-api/app/config"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/extra/bundebug"
)

type Ctx struct {
	Config *config.Config
	Db     *bun.DB
}

func New(cfg *config.Config) *Ctx {
	return &Ctx{
		Config: cfg,
	}
}

func (ctx *Ctx) Init() (err error) {
	if ctx.Db, err = ctx.Config.Db.New(); err != nil {
		return
	}

	ctx.debug()

	return
}

func (ctx *Ctx) debug() {
	if !ctx.Config.Debug {
		return
	}

	opts := []bundebug.Option{
		bundebug.WithVerbose(true),
	}
	ctx.Db.AddQueryHook(bundebug.NewQueryHook(opts...))
}
