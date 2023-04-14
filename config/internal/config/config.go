package config

import (
	"github.com/spf13/viper"
)

var config *MyConfig

type MyConfig struct {
	Port int    `mapstructure:"port"`
	Key  string `mapstructure:"key"`
}

func Initialize(fileName string) error {
	viper.SetConfigFile(fileName)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.AutomaticEnv()

	c := &MyConfig{}
	if err := viper.Unmarshal(c); err != nil {
		return err
	}
	config = c
	return nil
}

func GetConfig() *MyConfig {
	if config == nil {
		panic("config is not initialized.")
	}
	return config
}
