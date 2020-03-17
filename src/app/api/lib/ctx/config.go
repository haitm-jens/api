package ctx

import (
	"sync"

	"pandog/infra/local/lib"
)

var config *lib.Config

var once_config sync.Once

func NewConfig() *lib.Config {
	once_config.Do(func() {
		config = lib.NewConfig()
	})
	return config
}
