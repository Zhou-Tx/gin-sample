package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type IApplicationConfig interface {
	DefaultConfig()
}

func LoadConfig() ApplicationConfig {
	fmt.Println("Loading config...")
	Config.DefaultConfig()
	// 设置配置文件路径和类型
	viper.AddConfigPath("config")        // 配置文件路径
	viper.SetConfigName("configuration") // 文件名（不带扩展名）
	viper.SetConfigType("yaml")          // 文件类型

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// 将配置映射到结构体中
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
	return Config
}
