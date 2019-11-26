package config

import (
	"github.com/double1996/joplin-web-blog/pkg/logger"
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

type GithubConfiguration struct {
	GithubClientId     string `yaml:"github_clientid"`
	GithubClientSecret string `yaml:"github_clientsecret"`
	GithubAuthUrl      string `yaml:"github_authurl"`
	GithubRedirectURL  string `yaml:"github_redirecturl"`
	GithubTokenUrl     string `yaml:"github_tokenurl"`
	GithubScope        string `yaml:"github_scope"`
}

type LogConfiguration struct {
	Dir    string
	Access bool // request website log
}

type JoplinConfig struct {
	Token string
	Host  string
	Port  string
	Cron  string
}

type Configuration struct {
	Server   ServerConfiguration
	DataBase DataBaseConfiguration
	Log      LogConfiguration
	Joplin   JoplinConfig
	Github   GithubConfiguration
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
