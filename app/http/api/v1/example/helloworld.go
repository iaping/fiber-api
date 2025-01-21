package example

import (
	"fiber-api/app/http/api"
)

// @Tags Example
// @Summary hello world
// @Produce json
// @Router /v1/example/helloworld [get]
// @Success 200
func (i *Example) Helloworld(ctx *api.Ctx) error {
	return ctx.Response("hello world !!!")
}
