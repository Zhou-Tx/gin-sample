package config

var Config ApplicationConfig

type ApplicationConfig struct {
	Server ServerConfig
	Mysql  MysqlConfig
	Redis  RedisConfig
}

func (c *ApplicationConfig) DefaultConfig() {
	c.Server.DefaultConfig()
	c.Mysql.DefaultConfig()
	c.Redis.DefaultConfig()
}
