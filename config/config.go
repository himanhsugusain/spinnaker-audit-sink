// Package config
package config

import (
	"github.com/spf13/viper"
)

type Filter struct {
	DetailsType []string
}
type Config struct {
	UserName string
	PassWord string
	Port     string
	Filter   Filter
	Sinks    []string
}

func GetConfig() (Config, error) {
	cfg := Config{
		UserName: "foo",
		PassWord: "bar",
		Port:     "8080",
	}
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
