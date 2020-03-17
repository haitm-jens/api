package lib

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		root string
	}
)

func NewConfig() *Config {
	return &Config{}
}

func (s *Config) Root(root string) *Config {
	s.root = root
	return s
}

func (s *Config) Get(name string) string {
	return viper.GetString(name)
}

func (s *Config) GetInt(name string) int {
	return viper.GetInt(name)
}

func (s *Config) GetStringMapString(name string) map[string]string {
	return viper.GetStringMapString(name)
}

func (s *Config) LoadDir(args ...string) {
	dir := s.root
	if len(args) > 0 {
		dir = path.Join(dir, args[0])
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}

		t := info.Name()
		n := strings.TrimSuffix(t, filepath.Ext(t))
		load(n, filepath.Dir(path))
		return nil
	})

	if err != nil {
		panic(err)
	}
}

func (s *Config) Load(name string, args ...string) {
	dir := s.root
	if len(args) > 0 {
		dir = path.Join(dir, args[0])
	}
	load(name, dir)
}

func load(name string, dir string) {
	viper.AddConfigPath(dir)
	viper.SetConfigName(name)

	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
