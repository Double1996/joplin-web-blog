package config

import (
	"github.com/BurntSushi/toml"
)

const filePath = "config/config.toml"

type Config struct {
	Name string `toml: `
}

func configPath() string {

}
