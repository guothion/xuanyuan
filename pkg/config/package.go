package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var (
	Config        config
	viperInstance *viper.Viper
	configLock    = new(sync.RWMutex)
)

func Init() {
	viperInstance = initConfig()
	// 成功读取了 config
	loadConfig()
}

func initConfig() *viper.Viper {
	viperInstance = viper.New()

	// set config name
	viperInstance.SetConfigName("local.application") // name of config file

	viperInstance.SetConfigType("yaml")
	viperInstance.AddConfigPath("../cfg/") // path to look for the config file in
	viperInstance.AddConfigPath("./cfg/")  // call multiple times to add many search paths

	err := viperInstance.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	return viperInstance

}

func loadConfig() {
	var tempConfig config
	if err := viperInstance.Unmarshal(&tempConfig); err != nil {
		panic(fmt.Errorf("unable to decode config: %s \n", err))
	}

	// we use lock to read the config,because we use goroutine to open multi-thread
	// for prevent data race we locked it.
	configLock.Lock()
	defer configLock.Unlock()
	Config = tempConfig
}
