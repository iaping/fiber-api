package middleware

import (
	"fiber-api/app/ctx"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Test struct {
	ctx *ctx.Ctx
}

func NewTest(ctx *ctx.Ctx) *Test {
	return &Test{
		ctx: ctx,
	}
}

func (i *Test) Handle(app *fiber.Ctx) error {
	fmt.Println("I'm a middleware, debug:", i.ctx.Cfg.Debug)
	return app.Next()
}
