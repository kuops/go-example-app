package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	Port string
	Database string
	Environment string
)

type Config struct {
	Port  string  `yaml:"port"`
	Database string `yaml:"database"`
	Environment string `yaml:"environment"`
}

func InitConfig(cfile string) {
	var config Config
	v := viper.New()
	v.SetConfigFile(cfile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Warnf("Config file changed: %v", e.Name)
		if err := v.Unmarshal(&config); err != nil {
			log.Panicf("Config file format error %v",err)
		}
	})
	if err := v.Unmarshal(&config); err != nil {
		log.Panicf("Config file format error %v",err)
	}
	Port = config.Port
	Database = config.Database
	Environment = config.Environment
}
