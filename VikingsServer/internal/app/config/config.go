package config

import (
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"time"
)

type Config struct {
	ServiceHost string
	ServicePort int
	JWT         JWTConfig
}

type JWTConfig struct {
	Token         string
	ExpiresIn     time.Duration
	SigningMethod jwt.SigningMethod
}

func NewConfig(log *logrus.Logger) (*Config, error) {
	var err error

	configName := "config"
	_ = godotenv.Load()
	if os.Getenv("CONFIG_NAME") != "" {
		configName = os.Getenv("CONFIG_NAME")
	}

	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	viper.WatchConfig()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	log.Info("config parsed")
	cfg.JWT.Token = "test"
	cfg.JWT.ExpiresIn = time.Hour
	cfg.JWT.SigningMethod = jwt.SigningMethodHS256
	return cfg, nil
}
