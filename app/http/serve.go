package http

import (
	"fiber-api/app/ctx"
	"fiber-api/app/http/api"
	v1 "fiber-api/app/http/api/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

type Serve struct {
	ctx *ctx.Ctx
	s   *fiber.App
}

func New(ctx *ctx.Ctx) *Serve {
	cfg := fiber.Config{}

	return &Serve{
		ctx: ctx,
		s:   fiber.New(cfg),
	}
}

func (s *Serve) Run() (err error) {
	s.init()
	return s.s.Listen(s.ctx.Config.Http.Addr)
}

func (s *Serve) init() {
	s.cors()
	s.router()
}

func (s *Serve) cors() {
	cfg := cors.Config{
		AllowOrigins:     s.ctx.Config.Http.Cors.AllowOrigins,
		AllowMethods:     s.ctx.Config.Http.Cors.AllowMethods,
		AllowHeaders:     s.ctx.Config.Http.Cors.AllowHeaders,
		AllowCredentials: s.ctx.Config.Http.Cors.AllowCredentials,
		ExposeHeaders:    s.ctx.Config.Http.Cors.ExposeHeaders,
	}
	s.s.Use(cors.New(cfg))
}

func (s *Serve) router() {
	if s.ctx.Config.Debug {
		s.s.Get("/_doc/*", swagger.HandlerDefault)
	}

	// your api
	r := api.NewRouter(s.s.Group("/v1"), s.ctx)
	r.Inject([]api.IApi{
		v1.Test,
	}...)
}
