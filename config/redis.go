package config

import "github.com/spf13/viper"

const redisPrefix = "redis."

type RedisConfig struct {
	Host     string
	Port     uint16
	Username string
	Password string
	Database uint16
}

func (c *RedisConfig) DefaultConfig() {
	viper.SetDefault(redisPrefix+"host", "127.0.0.1")
	viper.SetDefault(redisPrefix+"port", 6379)
}
