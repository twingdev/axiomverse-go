package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	NatsURL  string `mapstructure:"nats_url"`
	NodeType string `mapstructure:"node_type"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func init() {
	_, err := LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}
}
