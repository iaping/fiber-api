package main

import (
	"fiber-api/app"
	cx "fiber-api/app/ctx"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

// @title fiber-api
// @version v1
// @description fiber-api is a basic api framework developed based on fiber
func main() {
	ctx, err := cx.Load("./config.yaml")
	if err != nil {
		log.Panic().Err(err).Msg("Ctx")
		return
	}

	cmd := &cli.App{
		Name: "fiber-api",
		Action: func(*cli.Context) error {
			return app.New(ctx).Run()
		},
	}
	if err = cmd.Run(os.Args); err != nil {
		log.Err(err).Msg("Run")
	}
}
