package config

import (
	"github.com/double1996/smart-evernote-blog/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Conf Configuration

type ServerConfiguration struct {
	Port string
	ICP  string
}

type DataBaseConfiguration struct {
	Driver   string
	DbName   string
	Username string
	Password string
	Host     string
	Port     string
}

type LogConfiguration struct {
	Dir    string
	Access bool // request website log
}

type EverNoteConfiguration struct {
	EverNoteKey         string
	EverNoteSecret      string
	EverNoteAuthorToken string
}

type Configuration struct {
	Server   ServerConfiguration
	DataBase DataBaseConfiguration
	Log      LogConfiguration
}

func InitConfig() {

	viper.SetConfigFile("./deployments/config.yml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("Error reading config file", zap.String("config", err.Error()))
	}
	err := viper.Unmarshal(&Conf)
	if err != nil {
		logger.Fatal("Unable to decode into struct", zap.String("config", err.Error()))
	}
}
