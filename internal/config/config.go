package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true" `
}

type Config struct {
	Env         string `yaml:"env" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config Path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file doesn't exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Unable to Parse Config File: %s", err.Error())
	}

	return &cfg
}
