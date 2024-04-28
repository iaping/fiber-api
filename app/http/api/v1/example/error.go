package example

import (
	"errors"
	"fiber-api/app/http/api"
)

// @Tags Example
// @Summary Example
// @Produce json
// @Router /v1/example/error [get]
// @Success 200
func (api *Example) Error(ctx *api.Ctx) (interface{}, error) {
	return nil, errors.New("this is an example of what went wrong")
}
