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
	if err = ctx.initDb(); err != nil {
		return
	}

	if err = ctx.initRds(); err != nil {
		return
	}

	return
}

func (ctx *Ctx) initDb() (err error) {
	if !ctx.Config.Db.Enable {
		return
	}

	if ctx.Db, err = ctx.Config.Db.New(); err != nil {
		return
	}

	// debug
	if ctx.Config.Debug {
		opts := []bundebug.Option{
			bundebug.WithVerbose(true),
		}
		ctx.Db.AddQueryHook(bundebug.NewQueryHook(opts...))
	}

	return
}

func (ctx *Ctx) initRds() (err error) {
	if !ctx.Config.Redis.Enable {
		return
	}

	if ctx.Rds, err = ctx.Config.Redis.New(); err != nil {
		return
	}

	return
}
