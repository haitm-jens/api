package lib

import (
	v "github.com/spf13/viper"
)

type Config struct {
}

func (c *Config) Get(name string) string {
	return v.GetString(name)
}
