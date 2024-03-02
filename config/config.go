package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DB DBConfig `yaml:"db"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func NewConfig() (*Config, error) {
	config := &Config{}
	file, err := os.Open("config/server.yaml")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
