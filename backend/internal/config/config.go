package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	PostgresUser string `mapstructure:"POSTGRES_USER"`
	PostgresPass string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB   string `mapstructure:"POSTGRES_DB"`
	PostgresHost string `mapstructure:"POSTGRES_HOST"`
	PostgresPort string `mapstructure:"POSTGRES_PORT"`
	RedisHost    string `mapstructure:"REDIS_HOST"`
	RedisPort    string `mapstructure:"REDIS_PORT"`
	RedisPass    string `mapstructure:"REDIS_PASSWORD"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	viper.SetDefault("PORT", "8080")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	t := reflect.TypeOf(cfg)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("mapstructure")
		if tag != "" {
			viper.BindEnv(tag)
		}
	}

	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if cfg.PostgresPass == "" || cfg.PostgresUser == "" || cfg.PostgresDB == "" || cfg.PostgresHost == "" || cfg.PostgresPort == "" {
		return nil, fmt.Errorf("variáveis obrigatórias do banco de dados não configuradas")
	}
	if cfg.RedisHost == "" || cfg.RedisPort == "" || cfg.RedisPass == "" {
		return nil, fmt.Errorf("variáveis obrigatórias do Redis não configuradas")
	}

	return &cfg, nil
}
