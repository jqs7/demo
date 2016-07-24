package app

import (
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var v struct {
	v *viper.Viper
	sync.Once
}

func Config() *viper.Viper {
	v.Do(func() {
		v.v = viper.New()
		v.v.SetConfigName("app")
		v.v.AddConfigPath(".")
		v.v.WatchConfig()
		v.v.ReadInConfig()
		v.v.OnConfigChange(func(fsnotify.Event) {})
	})
	return v.v
}
