package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)
type Workflow struct {
	Model  string
	Config string
	Input  string
	Task   []string
}

var WF Workflow
func InitWorkflow(configpath, filename string) {
	viper.SetConfigName(filename)  //名
	viper.SetConfigType("yaml")     //类型
	viper.AddConfigPath(configpath) //
	log.Println("workflow文件: %v",filename)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取workflow文件失败: %v", err)
	}
	err = viper.Unmarshal(&WF)
	if err != nil {
		log.Fatalf("解析workflow文件失败: %v", err)
	}
	fmt.Println("✅ workflow加载成功")
}
