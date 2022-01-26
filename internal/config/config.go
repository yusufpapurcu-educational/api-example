package config

import "github.com/BurntSushi/toml"

type Config struct {
	Server   ServerConfig   `toml:"server"`
	Database DatabaseConfig `toml:"database"`
}

func ParseConfig(file_name string) Config {
	var res Config
	a, err := toml.DecodeFile(file_name, &res)
	_, _ = a, err
	return res
}

type DatabaseConfig struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
	TimeZone string `toml:"timezone"`
	Port     int    `toml:"port"`
}

type ServerConfig struct {
	Port    int `toml:"port"`
	Timeout int `toml:"timeout"`
}
