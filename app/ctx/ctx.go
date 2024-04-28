package ctx

import (
	"fiber-api/app/config"

	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/extra/bundebug"
)

type Ctx struct {
	Config *config.Config
	Db     *bun.DB
	Rds    *redis.Client
}

func New(cfg *config.Config) *Ctx {
	return &Ctx{
		Config: cfg,
	}
}

func (ctx *Ctx) Init() (err error) {
	if ctx.Config.Db.Enable {
		if ctx.Db, err = ctx.Config.Db.New(); err != nil {
			return
		}
	}

	if ctx.Config.Redis.Enable {
		if ctx.Rds, err = ctx.Config.Redis.New(); err != nil {
			return
		}
	}

	ctx.debug()

	return
}

func (ctx *Ctx) debug() {
	if !ctx.Config.Debug {
		return
	}

	if ctx.Config.Db.Enable {
		opts := []bundebug.Option{
			bundebug.WithVerbose(true),
		}
		ctx.Db.AddQueryHook(bundebug.NewQueryHook(opts...))
	}
}
