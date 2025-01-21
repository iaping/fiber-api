package example

import (
	"fiber-api/app/http/api"
)

// @Tags Example
// @Summary 自定义错误（接口会暴露错误信息，需要定义好错误的相关信息）
// @Produce json
// @Router /v1/example/error/custom [get]
// @Success 200
func (i *Example) ErrorCustom(ctx *api.Ctx) error {
	return api.NewError(10000, "this is a custom error")
}
