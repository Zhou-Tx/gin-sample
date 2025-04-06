package config

import "github.com/spf13/viper"

const mysqlPrefix = "mysql."

type MysqlConfig struct {
	Host     string
	Port     uint16
	Username string
	Password string
	Schema   string
}

func (c *MysqlConfig) DefaultConfig() {
	viper.SetDefault(mysqlPrefix+"host", "127.0.0.1")
	viper.SetDefault(mysqlPrefix+"port", 3306)
	viper.SetDefault(mysqlPrefix+"username", "root")
}
