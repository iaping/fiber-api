package example

import (
	"context"
	"fiber-api/app/http/api"
	"time"
)

// @Tags Example
// @Summary redis示例
// @Produce json
// @Router /v1/example/redis [get]
// @Success 200
func (i *Example) Redis(ctx *api.Ctx) error {
	key := "example"

	status := ctx.Ctx.Rds.Set(context.Background(), key, time.Now(), time.Minute)
	if err := status.Err(); err != nil {
		return err
	}

	cmd := ctx.Ctx.Rds.Get(context.Background(), key)
	if err := cmd.Err(); err != nil {
		return err
	}

	val, err := cmd.Result()
	if err != nil {
		return err
	}

	return ctx.Json(map[string]interface{}{
		"key":   key,
		"value": val,
	})
}
