package main

import (
	// "encoding/json"
	"github.com/along416/promptDSL/codegen"
	"fmt"
	"log"
	"os/exec"
	"strings"

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

		promptfileContentstr := string(promptfileContent)
		nameWithoutExt := strings.TrimSuffix(filename, ".pdsl")

		prompt, err := codegen.RunPromptDSL(promptfileContentstr, nameWithoutExt)
		if err != nil {
			log.Fatalf("RunPromptDSL error: %v", err)
		}
		log.Println("生成的 Prompt:\n", prompt)
	}

	exename := os.Args[1]
	dir := "generated_code"
	// go build
	cmd := exec.Command("go", "build", "-o", exename, ".")
	cmd.Dir = dir
	// 获取命令的输出和错误信息
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(cmdOutput)) // 打印 go build 的错误输出
		log.Fatalf("执行 go build 失败: %v\n", err)
	}
	fmt.Println("Go 程序编译完成，生成了 CodeRunner.exe 文件")
}

