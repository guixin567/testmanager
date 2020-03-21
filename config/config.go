package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	AppName    string   `json:"app_name"`
	Port       string   `json:"port"`
	StaticPath string   `json:"static_path"`
	Model      string   `json:"model"`
	DataBase   DataBase `json:"database"`
	Redis      Redis    `json:"redis"`
}

//mysql数据库信息
type DataBase struct {
	User   string `json:"user"`
	Pwd    string `json:"pwd"`
	Port   string `json:"port"`
	Driver string `json:"driver"`
	DBName string `json:"dbname"`
}

//redis数据库信息
type Redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	NetWork  string `json:"network"`
	Password string `json:"password"`
	Prefix   string `json:"prefix"`
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
