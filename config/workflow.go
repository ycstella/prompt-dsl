package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Workflow struct {
	Model  string
	Config string
	Input  string
	Task   []string
}

var WF Workflow

func InitWorkflow(configpath, filename string) {
	path:=filepath.Join(configpath, filename+".wyaml")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("读取文件失败: %v", err)
	}

	if err := yaml.Unmarshal(data, &WF); err != nil {
		log.Fatalf("解析 wyaml 文件失败: %v", err)
	}

	fmt.Println("✅ workflow加载成功:",)
}
