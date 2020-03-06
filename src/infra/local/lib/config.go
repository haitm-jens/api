package lib

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		root string
	}
)

func NewConfig(root string) *Config {
	return &Config{root: root}
}

func (s *Config) Get(name string) string {
	return viper.GetString(name)
}

func (s *Config) GetStringMapString(name string) map[string]string {
	return viper.GetStringMapString(name)
}

func (s *Config) LoadDir(args ...string) {
	dir := s.root
	if len(args) > 0 {
		dir = dir + "/" + args[0]
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
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
	path := s.root
	if len(args) > 0 {
		path += "/" + args[0]
	}
	load(name, path)
}

func load(name string, dir string) {
	viper.AddConfigPath(dir)
	viper.SetConfigName(name)

	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
