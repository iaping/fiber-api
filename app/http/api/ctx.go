package api

import (
	"fiber-api/app/ctx"

	"github.com/gofiber/fiber/v2"
)

type Ctx struct {
	Fiber *fiber.Ctx
	App   *ctx.Ctx
}
