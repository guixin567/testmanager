package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	AppName    string `json:"app_name"`
	Port       string `json:"port"`
	StaticPath string `json:"static_path"`
	Model      string `json:"model"`
}

func InitConfig() *AppConfig {
	file, err := os.Open("./config/config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	appConfig := AppConfig{}
	err = decoder.Decode(&appConfig)
	if err != nil {
		panic(err.Error())
	}
	return &appConfig
}
