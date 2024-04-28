package example

import (
	"context"
	"fiber-api/app/http/api"
	"fiber-api/app/model"

	"github.com/uptrace/bun"
)

// @Tags Example
// @Summary example mysql
// @Produce json
// @Router /v1/example/mysql [get]
// @Success 200
func (api *Example) Mysql(ctx *api.Ctx) (interface{}, error) {
	type timeZoneName struct {
		model.Model
		bun.BaseModel `bun:"table:time_zone_name,alias:tzn"`
		TimeZoneId    int    `json:"time_zone_id"`
		Name          string `json:"name"`
	}

	var data []timeZoneName
	q := ctx.App.Db.NewSelect().Model(&data)
	q.Limit(15)
	q.Order("Time_zone_id ASC")
	err := q.Scan(context.Background())

	return data, err
}
