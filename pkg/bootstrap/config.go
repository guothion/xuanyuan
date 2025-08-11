package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/guothion/xuanyuan/pkg/global"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() *viper.Viper {
	viperInstance := viper.New()

	// set config name
	viperInstance.SetConfigName("local.application") // name of config file

	viperInstance.SetConfigType("yaml")
	viperInstance.AddConfigPath("../cfg/") // path to look for the config file in
	viperInstance.AddConfigPath("./cfg/")  // call multiple times to add many search paths

	err := viperInstance.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监听配置文件
	viperInstance.WatchConfig()
	viperInstance.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		// 重新载入配置
		if err := viperInstance.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})

	if err := viperInstance.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}

	logrus.Printf("%#v", global.App.Config.Log)

	return viperInstance
}
