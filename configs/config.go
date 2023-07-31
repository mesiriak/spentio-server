package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppAddress string `mapstructure:"API_ADDRESS" default:":8080"`
	JwtSecret  string `mapstructure:"JWT_SECRET"`

	RedisDsn    string `mapstructure:"BLACKLIST_URL"`
	PostgresDsn string `mapstructure:"POSTGRES_URL"`
}

func CreateConfig() Config {
	var config Config

	viper.SetConfigFile("configs/.env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Reading config failed, %s", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unmarshaling config failed, %s", err)
	}

	return config
}
