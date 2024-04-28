package example

import (
	"fiber-api/app/http/api"
)

// @Tags Example
// @Summary Example
// @Produce json
// @Router /v1/example/helloworld [get]
// @Success 200
func (api *Example) Helloworld(ctx *api.Ctx) (interface{}, error) {
	return "hello world !!!", nil
}
