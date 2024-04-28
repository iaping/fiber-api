package middleware

import (
	"fiber-api/app/ctx"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Test struct {
	app *ctx.Ctx
}

func NewTest(app *ctx.Ctx) *Test {
	return &Test{
		app: app,
	}
}

func (i *Test) Handle(ctx *fiber.Ctx) error {
	fmt.Println("I'm a middleware, debug:", i.app.Config.Debug)
	return ctx.Next()
}
