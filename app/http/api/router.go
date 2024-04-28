package api

import (
	"fiber-api/app/ctx"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type HandlerFunc func(*Ctx) (interface{}, error)

type IRouter interface {
	Use(...interface{}) IRouter
	Group(path string) IRouter

	Get(string, HandlerFunc) IRouter
	Post(string, HandlerFunc) IRouter
	Put(string, HandlerFunc) IRouter
	Delete(string, HandlerFunc) IRouter
}

type Router struct {
	fiber.Router
	ctx *ctx.Ctx
}

func NewRouter(router fiber.Router, ctx *ctx.Ctx) *Router {
	return &Router{
		Router: router,
		ctx:    ctx,
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
	return NewRouter(r.Router.Group(path), r.ctx)
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
	var resp Response

	if data, err := handler(&Ctx{Fiber: ctx, App: r.ctx}); err != nil {
		log.Err(err).Msgf("path: %s", ctx.Request().URI().Path())
		resp = r.error(err)
	} else {
		resp = NewResponse(data)
	}

	return ctx.Status(resp.Status).JSON(resp)
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
