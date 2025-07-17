package config

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	Config        config
	viperInstance *viper.Viper
	configLock    = new(sync.RWMutex)
)

func Init() {

}

func initConfig() *viper.Vipper {

}
