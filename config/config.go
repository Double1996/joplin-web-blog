package config

import (
	"github.com/BurntSushi/toml"
	"github.com/double1996/smart-evernote-blog/helper"
	"path"
	"sync"
)

var (
	once sync.Once
	Cfg  config
)

const filePath = "deployments/config.toml"

type config struct {
	Name  string `toml:name`
	Url   string `toml:url` //
	Host  string `toml:host`
	Debug bool
	Log   Log `toml: log`
}

type Log struct {
	Mode   string `toml:mode`
	Dir    string `toml:dir`
	Format string
	Access bool // request website log
}

func initcfg() {
	once.Do(func() {
		path := configPath()
		if _, err := toml.Decode(path, &Cfg); err != nil {
			helper.Panicf("cannot parse config file. %v", err)
		}
	})
}

func configPath() string {
	if configDir == "" {
		configDir = path.Join(Root, filePath)
	}
	return configDir
}
