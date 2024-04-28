package main

import (
	"fiber-api/app"
	"fiber-api/app/config"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out: os.Stderr,
	})
}

// @title api
// @version v1
// @description fiber-api is a basic api framework developed based on fiber
func main() {
	cfg, err := config.Load("./config.yaml")
	if err != nil {
		log.Panic().Err(err).Msg("Config")
	}

	if err = app.New(cfg).Run(); err != nil {
		log.Panic().Err(err).Msg("App")
	}
}
