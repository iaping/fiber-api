package example

import (
	"context"
	"fiber-api/app/http/api"
	"time"
)

// @Tags Example
// @Summary Example
// @Produce json
// @Router /v1/example/redis [get]
// @Success 200
func (api *Example) Redis(ctx *api.Ctx) (interface{}, error) {
	key := "example"

	status := ctx.App.Rds.Set(context.Background(), key, time.Now(), time.Minute)
	if err := status.Err(); err != nil {
		return nil, err
	}

	cmd := ctx.App.Rds.Get(context.Background(), key)
	if err := cmd.Err(); err != nil {
		return nil, err
	}

	val, _ := cmd.Result()
	return map[string]interface{}{
		"key":   key,
		"value": val,
	}, cmd.Err()
}
