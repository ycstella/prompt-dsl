package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/along416/promptDSL/codegen"

	// "strings"

	// "net/http"
	// "codegen"

	"os"
)

// var llm *service.LLMClient
func main() {

	promptDir := os.Args[2]
	entries, err := os.ReadDir(promptDir)
	if err != nil {
		log.Println("读取目录失败：", err)
	}
	// 确保生成代码目录存在
	
	genDir := "generated_code/generated"
	err = os.MkdirAll(genDir, os.ModePerm)
	if err != nil {
		fmt.Println("创建目录失败: %v", err)
	}
	fmt.Println("创建目录: %v", err)
	for _, entry := range entries {
		filename := entry.Name()
		log.Println("开始处理文件：", filename)
		// 只处理 .prompt 文件
		fmt.Println("entry.Name():", filename)
		if !strings.HasSuffix(filename, ".pdsl") {
			log.Println("不是 .pdsl 文件")
			continue
		}
		// 拼接完整路径
		path := promptDir + filename
		// fileContent, err := os.ReadFile("./promptdsl-core/prompts/SplitSolutionSteps.prompt")
		promptfileContent, err := os.ReadFile(path)
		if err != nil {
			panic(fmt.Errorf("读取 DSL 文件失败: %v", err))
		}

		promptFileContentstr := string(promptfileContent)
		nameWithoutExt := strings.TrimSuffix(filename, ".pdsl")

		prompt, err := codegen.RunPromptDSL(promptFileContentstr, nameWithoutExt)
		if err != nil {
			log.Fatalf("RunPromptDSL error: %v", err)
		}
		log.Println("生成的 Prompt:\n", prompt)
	}

	exeName := os.Args[1]
	
	// go build
	cmd := exec.Command("go", "build", "-o", exeName, ".")
	cmd.Dir = "generated_code"
	// 获取命令的输出和错误信息
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(cmdOutput)) // 打印 go build 的错误输出
		log.Fatalf("执行 go build 失败: %v\n", err)
	}
	fmt.Println("Go 程序编译完成，生成了 CodeRunner.exe 文件")
}
