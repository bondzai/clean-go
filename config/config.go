package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort       string
	RedisURL      string
	RedisUser     string
	RedisPassword string
	RedisDatabase int
}

var AppConfig Config

func LoadConfig() error {
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	AppConfig = Config{
		AppPort:       viper.GetString("app.port"),
		RedisURL:      viper.GetString("redis.url"),
		RedisUser:     viper.GetString("redis.user"),
		RedisPassword: viper.GetString("redis.password"),
		RedisDatabase: viper.GetInt("redis.database"),
	}

	return nil
}
