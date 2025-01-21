package main

import (
	"fiber-api/app"
	cx "fiber-api/app/ctx"

	"github.com/rs/zerolog/log"
)

// @title fiber-api
// @version v1
// @description fiber-api is a basic api framework developed based on fiber
func main() {
	ctx, err := cx.Load("./config.yaml")
	if err != nil {
		log.Panic().Err(err).Msg("Ctx")
	}

	if err = app.New(ctx).Run(); err != nil {
		log.Panic().Err(err).Msg("App")
	}
}
