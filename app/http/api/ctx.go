package api

import (
	"fiber-api/app/ctx"

	"github.com/gofiber/fiber/v2"
)

type Ctx struct {
	Fiber *fiber.Ctx
	App   *ctx.Ctx
}

func (ctx *Ctx) Response(data interface{}) error {
	resp := NewResponse(data)
	return ctx.Json(resp)
}

func (ctx *Ctx) Json(data interface{}, ctype ...string) error {
	return ctx.JsonWithStatus(fiber.StatusOK, data, ctype...)
}

func (ctx *Ctx) JsonWithStatus(status int, data interface{}, ctype ...string) error {
	return ctx.Fiber.Status(status).JSON(data, ctype...)
}
