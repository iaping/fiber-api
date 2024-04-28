package config

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Enable   bool   `yaml:"Enable"`
	Host     string `yaml:"Host"`
	Port     uint16 `yaml:"Port"`
	Password string `yaml:"Password"`
	DB       int    `yaml:"DB"`
}

func (i Redis) New() (*redis.Client, error) {
	rds := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", i.Host, i.Port),
		Password: i.Password,
		DB:       i.DB,
	})

	cmd := rds.Ping(context.Background())
	return rds, errors.Wrap(cmd.Err(), "redis")
}
