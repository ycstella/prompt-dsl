package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string
	Port int
}

type LogConfig struct {
	Level string
	File  string
}

type ModelConfig struct {
	Provider    string
	ApiKey      string `mapstructure:"api_key"`
	BaseURL     string `mapstructure:"base_url"`
	Model       string
	Temperature float64
	MaxTokens   int `mapstructure:"max_tokens"`
}

type AppConfig struct {
	Server ServerConfig
	Log    LogConfig
	Model  ModelConfig
}

var Cfg AppConfig

func InitConfig() {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config") // 相对路径，可按需调整

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}

	fmt.Println("✅ 配置加载成功")
}
func InitLogger() {
	logPath := Cfg.Log.File
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件: %v", err)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if Cfg.Log.Level == "debug" {
		log.Println("日志等级为 DEBUG")
	}
}
