package lib

import (
	_ "pandog/infra/lib"

	"github.com/spf13/viper"
)

type (
	config struct {
	}
)

func NewConfig() *config {
	return &config{}
}

func (*config) Get(name string) string {
	return viper.GetString(name)
}
