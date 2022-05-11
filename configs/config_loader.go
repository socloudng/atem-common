package configs

import (
	"atem/atem-common/utils"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ConfigLoader interface {
	GetConfigType() string
}

func LoadConfigByViper(bindValue ConfigLoader, paths ...string) *viper.Viper {
	v := viper.New()
	v.SetConfigType(bindValue.GetConfigType())
	for _, path := range paths {
		vTemp := viper.New()
		vTemp.SetConfigFile(path)
		vTemp.SetConfigType(bindValue.GetConfigType())
		err := vTemp.ReadInConfig() //读取配置
		if err != nil {
			log.Fatalf("Fatal error config file: %s \n", err.Error())
		}
		mergeChildSettings(vTemp, v) //合并配置
		vTemp.WatchConfig()          //监听配置变化
		vTemp.OnConfigChange(func(e fsnotify.Event) {
			log.Printf("config file changed: %s", e.Name)
			mergeChildSettings(vTemp, v) //合并配置
			unmarshal(v, bindValue)
			//TODO 根据
		})
	}
	unmarshal(v, bindValue)
	return v
}

func SaveConfig(loader ConfigLoader) (err error) {
	viper := viper.New()
	cs := utils.StructToMap(loader)
	for k, v := range cs {
		viper.Set(k, v)
	}
	err = viper.WriteConfig()
	return err
}

func mergeChildSettings(child, parent *viper.Viper) {
	settings := child.AllSettings()
	for k, val := range settings {
		parent.SetDefault(k, val)
	}
}

func unmarshal(v *viper.Viper, bindValue interface{}) {
	if err := v.Unmarshal(bindValue); err != nil {
		log.Println(err)
	}
}
