package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Listen  Listen  `yaml:"listen"`
	Storage Storage `yaml:"storage"`
}

type Listen struct {
	BindIp string `yaml:"bind_ip"`
	Port   string `yaml:"port"`
}

type Storage struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"database"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

var cfg *Config

func MustConfig() *Config {
	cfg = &Config{}

	if err := cleanenv.ReadConfig("config.yaml", cfg); err != nil {
		panic("cannot read config " + err.Error())
	}

	return cfg
}
