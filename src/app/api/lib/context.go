package lib

import (
	"pandog/infra/local/lib"
	"sync"
)

var config *lib.Config
var once sync.Once

func NewConfig() *lib.Config {
	once.Do(func() {
		config = lib.NewConfig()
	})
	return config
}
