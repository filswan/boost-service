package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Configuration struct {
	Port  int    `toml:"port"`
	Boost *Boost `toml:"boost"`
	Log   *Log   `toml:"log"`
}

type Boost struct {
	FullAPI string `toml:"full_api"`
	Repo    string `toml:"repo"`
}

type Log struct {
	Env   string `toml:"env"`
	Level int    `toml:"level"`
}

var conf Configuration

var (
	confPath = "config/config.toml"
)

func Init(configPath ...string) error {
	if len(configPath) > 0 && configPath[0] != "" {
		confPath = configPath[0]
	}
	_, err := toml.DecodeFile(confPath, &conf)
	if err != nil {
		log.Fatal("Error:", err)
	}
	return nil
}

func Conf() *Configuration {
	return &conf
}
