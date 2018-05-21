package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Database DatabaseConfig
	Daemon   DaemonConfig
	Log      LogConfig
}

type DatabaseConfig struct {
	ConnectionInfo string `toml:"connection_info"`
}

type DaemonConfig struct {
	Address string `toml:"address"`
}

type LogConfig struct {
	Level string `toml:"level"`
}

func Load(filename string) *Config {
	log.Println("Loading config `" + filename + "`")

	var cfg Config
	if _, err := toml.DecodeFile(filename, &cfg); err != nil {
		log.Fatal(err.Error())
	}

	return &cfg
}
