package config

import (
	"github.com/spf13/viper"
)

func Load(path string, config interface{}) error {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(config)
}
