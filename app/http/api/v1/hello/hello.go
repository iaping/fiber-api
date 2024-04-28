package hello

import (
	"fiber-api/app/http/api"
)

// @Tags v1
// @Summary Test
// @Produce json
// @Router /v1/hello/world [get]
// @Success 200
func (api *Hello) World(ctx *api.Ctx) (interface{}, error) {
	return "hello world !!!", nil
}
