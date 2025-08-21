package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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
	Provider       string
	ApiKey         string `mapstructure:"api_key"`
	BaseURL        string `mapstructure:"base_url"`
	Model          string
	Temperature    float64
	MaxTokens      int `mapstructure:"max_tokens"`
	MaxConcurrency int
	Proxy          string
	Stream         bool
}

type Config struct {
	Server ServerConfig
	Log    LogConfig
	Models []ModelConfig
}

var Cfg Config

func InitConfig(configpath string) {
	viper.SetConfigName("default")  //名
	viper.SetConfigType("yaml")     //类型
	viper.AddConfigPath(configpath) //

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

	// 获取目录部分
	logDir := filepath.Dir(logPath)

	// 确保日志目录存在
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("无法创建日志目录: %v", err)
	}
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("无法打开日志文件: %v", err)
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if Cfg.Log.Level == "debug" {
		log.Println("日志等级为 DEBUG")
	}
}
func GetModelConfig(name string) *ModelConfig {
	for _, m := range Cfg.Models {
		if m.Model == name {
			return &m
		}
	}
	return nil
}
