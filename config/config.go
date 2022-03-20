package config

import (
	"github.com/allentom/harukap/config"
)

var DefaultConfigProvider *config.Provider

func InitConfigProvider() error {
	var err error
	DefaultConfigProvider, err = config.NewProvider(func(provider *config.Provider) {
		ReadConfig(provider)
	})
	return err
}

var Instance Config

type Config struct {
	UseReg string
}

func ReadConfig(provider *config.Provider) {
	configer := provider.Manager
	configer.SetDefault("addr", ":8000")
	configer.SetDefault("application", "Thumbnail Service")
	configer.SetDefault("instance", "main")

	Instance = Config{
		UseReg: configer.GetString("use_reg"),
	}
}
