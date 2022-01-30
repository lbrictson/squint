package conf

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DBUser     string `default:"postgres"`
	DBPassword string `default:"postgres"`
	DBName     string `default:"squint"`
	DBHost     string `default:"localhost"`
	DBPort     int    `default:"5432"`
	LogLevel   string `default:"info"`
	LogJSON    bool   `default:"true"`
	LogFile    string `default:""`
}

func New() (Config, error) {
	c := Config{}
	err := envconfig.Process("SQUINT", &c)
	return c, err
}
