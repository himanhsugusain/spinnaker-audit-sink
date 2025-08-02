// Package config
package config

import (
	"os"
)

type Config struct {
	UserName string
	PassWord string
	Port     string
}

func GetConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	userName := os.Getenv("USERNAME")
	if userName == "" {
		userName = "foo"
	}
	passWord := os.Getenv("PASSWORD")
	if passWord == "" {
		passWord = "bar"
	}
	return Config{
		UserName: userName,
		PassWord: passWord,
		Port:     port,
	}
}
