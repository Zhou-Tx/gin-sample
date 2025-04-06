package config

import "github.com/spf13/viper"

const serverPrefix = "server."

type ServerConfig struct {
	Host string
	Port uint16
}

func (c *ServerConfig) DefaultConfig() {
	viper.SetDefault(serverPrefix+"port", 8080)
}
