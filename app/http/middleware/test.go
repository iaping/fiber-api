package middleware

import (
	"fiber-api/app/ctx"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Test struct{}

func (i *Test) Handle(app *ctx.Ctx) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Println("I'm a middleware, debug:", app.Config.Debug)
		return ctx.Next()
	}
}
