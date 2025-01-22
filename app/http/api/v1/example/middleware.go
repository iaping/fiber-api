package example

import (
	"fiber-api/app/http/api"
)

// @Tags Example
// @Summary 中间件
// @Produce json
// @Router /v1/example/middleware [get]
// @Success 200
func (i *Example) Middleware(ctx *api.Ctx) error {
	return ctx.Json("I used a middleware")
}
