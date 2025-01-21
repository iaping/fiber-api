package http

import (
	"fiber-api/app/ctx"
	"fiber-api/app/http/api"
	v1 "fiber-api/app/http/api/v1"

	_ "fiber-api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog/log"
)

var (
	// your api
	apis = []api.IApi{
		v1.Example,
	}
)

type Serve struct {
	ctx  *ctx.Ctx
	s    *fiber.App
	addr string
}

func New(ctx *ctx.Ctx) *Serve {
	cfg := fiber.Config{
		DisableStartupMessage: true,
	}

	return &Serve{
		ctx:  ctx,
		s:    fiber.New(cfg),
		addr: ctx.Cfg.Http.Addr,
	}
}

func (s *Serve) Run() (err error) {
	s.debug()
	s.cors()
	s.router()

	l := log.Info()
	l.Str("Addr", s.addr)
	l.Msg("Server is running")

	return s.s.Listen(s.addr)
}

func (s *Serve) debug() {
	if !s.ctx.Cfg.Debug {
		return
	}

	l := log.Info()
	l.Str("Path", "/_doc/index.html")
	l.Msg("Swagger")

	s.s.Get("/_doc/*", swagger.HandlerDefault)
}

func (s *Serve) cors() {
	cfg := cors.Config{}
	s.s.Use(cors.New(cfg))
}

func (s *Serve) router() {
	r := api.NewRouter(s.s.Group("/v1"), s.ctx)
	r.Inject(apis...)
}
