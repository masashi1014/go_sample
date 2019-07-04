package config

import (
	"github.com/BurntSushi/toml"
)

const ConfPath = "./config/env/conf.toml"

type DBConfig struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Db       string `toml:"dbname"`
}

type Config struct {
	DB DBConfig `toml:"database"`
}

func NewConfig() (Config, error) {
	var conf Config

	if _, err := toml.DecodeFile(ConfPath, &conf); err != nil {
		return conf, err
	}

	return conf, nil
}
