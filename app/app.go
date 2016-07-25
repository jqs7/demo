package app

import (
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var config struct {
	v *viper.Viper
	sync.Once
}

func Config() *viper.Viper {
	config.Do(func() {
		config.v = viper.New()
		config.v.SetConfigName("app")
		config.v.AddConfigPath(".")
		config.v.WatchConfig()
		config.v.ReadInConfig()
		config.v.OnConfigChange(func(fsnotify.Event) {})
	})
	return config.v
}
