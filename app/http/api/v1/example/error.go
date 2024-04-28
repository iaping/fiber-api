package example

import (
	"errors"
	"fiber-api/app/http/api"
)

// @Tags Example
// @Summary 错误（接口不会暴露错误信息，应对一些未知或未处理的错误）
// @Produce json
// @Router /v1/example/error [get]
// @Success 200
func (i *Example) Error(ctx *api.Ctx) (interface{}, error) {
	return nil, errors.New("this is an example of what went wrong")
}
