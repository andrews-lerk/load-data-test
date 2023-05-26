package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func LoadConfig(path string) *Config {
	path = path + "/" + "app.yaml"
	f, err := os.Open(path)
	if err != nil {
		log.Print("Enable to open config file")
		log.Fatal(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Print("Enable to decode config file")
		log.Fatal(err)
	}
	return &cfg
}
