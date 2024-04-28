package config

type Http struct {
	Addr string `yaml:"Addr"`
	Cors struct {
		AllowOrigins     string `yaml:"AllowOrigins"`
		AllowMethods     string `yaml:"AllowMethods"`
		AllowHeaders     string `yaml:"AllowHeaders"`
		AllowCredentials bool   `yaml:"AllowCredentials"`
		ExposeHeaders    string `yaml:"ExposeHeaders"`
	} `yaml:"Cors"`
}
