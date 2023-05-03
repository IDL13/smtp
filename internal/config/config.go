package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	GmailKey string `yaml:"gmail_key"`
}

func GetConf() *Config {
	c := &Config{}
	info, err := os.ReadFile("./conf.yaml")
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(info, c)
	return c
}
