package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Environment               string `mapstructure:"ENVIRONMENT"`
	Debug                     bool   `mapstructure:"DEBUG"`
	DbMasterConnectionString  string `mapstructure:"DB_MASTER_CONNECTION_STRING"`
	DbReplicaConnectionString string `mapstructure:"DB_REPLICA_CONNECTION_STRING"`
	APIAddress                string `mapstructure:"API_ADDRESS"`
	RedisHost                 string `mapstructure:"REDIS_HOST"`
	RedisPort                 int    `mapstructure:"REDIS_PORT"`
	CacheTTL                  int    `mapstructure:"CACHE_TTL"`
}

func NewConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Could not load config with error: %s", err.Error())
		return nil, err
	}

	cfg := Config{}
	err = viper.Unmarshal(&cfg)

	if err != nil {
		log.Printf("Failed to load env variables. %+v\n", err)
		return nil, err
	}
	return &cfg, nil
}
