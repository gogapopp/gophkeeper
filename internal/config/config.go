package config

import (
	"sync"

	"github.com/spf13/viper"
)

const (
	configType     = "yaml"
	configFileName = "config"
	configPath     = "../../config"
)

var (
	errconfig error
	once      sync.Once
	config    *viper.Viper
)

// LoadConfig парсит конфиг
func LoadConfig() (*viper.Viper, error) {
	once.Do(func() {
		v := viper.New()
		v.SetConfigName(configFileName)
		v.SetConfigType(configType)
		v.AddConfigPath(configPath)
		v.BindEnv()
		err := v.ReadInConfig()
		if err != nil {
			errconfig = err
		}
		config = v
	})
	return config, errconfig
}
