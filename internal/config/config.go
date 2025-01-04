package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	BaseURL string
	APIKey  string
}

var AppConfig *Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()
	AppConfig = &Config{
		APIKey:  viper.GetString("ARVANCLOUD_APIKEY"),
		BaseURL: viper.GetString("ARVANCLOUD_BASEURL"),
	}
}
