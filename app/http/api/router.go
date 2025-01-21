package api

import (
	"fiber-api/app/ctx"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type HandlerFunc func(*Ctx) error

type IRouter interface {
	Use(...interface{}) IRouter
	Group(path string) IRouter

	Get(string, HandlerFunc) IRouter
	Post(string, HandlerFunc) IRouter
	Put(string, HandlerFunc) IRouter
	Delete(string, HandlerFunc) IRouter

	App() *ctx.Ctx
}

type Router struct {
	fiber.Router
	app *ctx.Ctx
}

func NewRouter(router fiber.Router, app *ctx.Ctx) *Router {
	return &Router{
		Router: router,
		app:    app,
	}
}

func (r *Router) Inject(apis ...IApi) {
	for _, i := range apis {
		i.SetRouter(r)
	}
}

func (r *Router) Use(args ...interface{}) IRouter {
	r.Router.Use(args...)
	return r
}

func (r *Router) Group(path string) IRouter {
	return NewRouter(r.Router.Group(path), r.app)
}

func (r *Router) App() *ctx.Ctx {
	return r.app
}

func (r *Router) Get(path string, handler HandlerFunc) IRouter {
	r.Router.Get(path, func(ctx *fiber.Ctx) error {
		return r.handle(ctx, handler)
	})
	return r
}

func (r *Router) Post(path string, handler HandlerFunc) IRouter {
	r.Router.Post(path, func(ctx *fiber.Ctx) error {
		return r.handle(ctx, handler)
	})
	return r
}

func (r *Router) Put(path string, handler HandlerFunc) IRouter {
	r.Router.Put(path, func(ctx *fiber.Ctx) error {
		return r.handle(ctx, handler)
	})
	return r
}

func (r *Router) Delete(path string, handler HandlerFunc) IRouter {
	r.Router.Delete(path, func(ctx *fiber.Ctx) error {
		return r.handle(ctx, handler)
	})
	return r
}

func (r *Router) handle(ctx *fiber.Ctx, handler HandlerFunc) error {
	if err := handler(&Ctx{Fiber: ctx, App: r.app}); err != nil {
		log.Err(err).Str("Path", string(ctx.Request().URI().Path())).Msg("Api")

		resp := r.error(err)
		return ctx.Status(resp.Status).JSON(resp)
	}
	return nil
}

func (r *Router) error(err error) Response {
	var i *Error

	switch e := err.(type) {
	case *Error:
		i = e
	case validator.ValidationErrors:
		i = ErrorParameter
	default:
		i = ErrorServer
	}

	return NewErrorResponse(i)
}
