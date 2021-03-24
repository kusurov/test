package configs

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Api struct {
		BindAddr   string `yaml:"bind_addr"`
		LogLevel   string `yaml:"log_level"`
		SessionKey string `yaml:"session_key"`
	} `yaml:"api"`
	Database struct {
		Host     string `yaml:"hast"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
	} `yaml:"database"`
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
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
