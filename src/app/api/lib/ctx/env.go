package ctx

import (
	"sync"
)

var env map[string]interface{}

var once_env sync.Once

func NewEnv() map[string]interface{} {
	once_env.Do(func() {
		env = make(map[string]interface{})
	})
	return env
}

func PushEnv(key string, value interface{}) {
	env[key] = value
}

func GetEnv(key string, args ...interface{}) interface{} {
	return env[key]
}
