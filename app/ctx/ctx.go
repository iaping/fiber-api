package ctx

import (
	"fiber-api/app/config"

	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/extra/bundebug"
)

type Ctx struct {
	Cfg *config.Config
	Db  *bun.DB
	Rds *redis.Client
}

func New(cfg *config.Config) *Ctx {
	return &Ctx{
		Cfg: cfg,
	}
}

func Load(path string) (*Ctx, error) {
	cfg, err := config.Load(path)
	if err != nil {
		return nil, err
	}

	ctx := New(cfg)
	return ctx, ctx.Init()
}

func (ctx *Ctx) Init() (err error) {
	if err = ctx.db(); err != nil {
		return
	}

	if err = ctx.rds(); err != nil {
		return
	}

	return
}

// init db
func (ctx *Ctx) db() (err error) {
	if !ctx.Cfg.Db.Enable {
		return
	}

	if ctx.Db, err = ctx.Cfg.Db.New(); err != nil {
		return
	}

	// debug
	if ctx.Cfg.Debug {
		opts := []bundebug.Option{
			bundebug.WithVerbose(true),
		}
		ctx.Db.AddQueryHook(bundebug.NewQueryHook(opts...))
	}

	return
}

// init redis
func (ctx *Ctx) rds() (err error) {
	if !ctx.Cfg.Redis.Enable {
		return
	}

	ctx.Rds, err = ctx.Cfg.Redis.New()
	return
}
