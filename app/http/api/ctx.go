package api

import (
	"fiber-api/app/ctx"

	"github.com/gofiber/fiber/v2"
)

type Ctx struct {
	App *fiber.Ctx
	Ctx *ctx.Ctx
}

func (ctx *Ctx) Json(data interface{}) error {
	var resp *Response
	switch i := data.(type) {
	case *Response:
		resp = i
	default:
		resp = NewResponse(data)
	}

	return ctx.JsonWithStatus(resp.Status, resp)
}

func (ctx *Ctx) JsonWithStatus(status int, data interface{}, ctype ...string) error {
	return ctx.App.Status(status).JSON(data, ctype...)
}

func (ctx *Ctx) IP() string {
	return ctx.App.IP()
}
