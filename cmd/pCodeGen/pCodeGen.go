package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/along416/promptDSL/codegen"

	// "strings"

	// "net/http"
	// "codegen"

	"os"
)

// var llm *service.LLMClient
func main() {
	if len(os.Args) < 2 {
		log.Fatalf("请提供 pdsl 文件路径作为参数")
	}

	// 从命令行获取单个 pdsl 文件路径
	pdslFile := os.Args[1]

	if !strings.HasSuffix(pdslFile, ".pdsl") {
		log.Fatalf("文件不是 .pdsl 文件: %s", pdslFile)
	}

	// 确保生成代码目录存在
	genDir := "generated_code"
	err := os.MkdirAll(genDir, os.ModePerm)
	if err != nil {
		log.Fatalf("创建生成目录失败: %v", err)
	}

	filename := filepath.Base(pdslFile)
	nameWithoutExt := strings.TrimSuffix(filename, ".pdsl")
	outDir := filepath.Join(genDir, nameWithoutExt)

	// 为该文件创建单独子目录
	err = os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		log.Fatalf("创建子目录失败: %v", err)
	}

	// 读取 pdsl 文件
	content, err := os.ReadFile(pdslFile)
	if err != nil {
		log.Fatalf("读取 pdsl 文件失败: %v", err)
	}

	// 生成 Prompt
	prompt, err := codegen.RunPromptDSL(string(content), nameWithoutExt)
	if err != nil {
		log.Fatalf("RunPromptDSL error: %v", err)
	}

	log.Println("生成的 Prompt:\n", prompt)

	// 编译生成 exe
	exeName := nameWithoutExt + ".exe"
	cmd := exec.Command("go", "build", "-o", exeName, ".")
	cmd.Dir = outDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		log.Fatalf("执行 go build 失败: %v", err)
	}

	fmt.Println("Go 程序编译完成，生成了", exeName)

	exePath := filepath.Join(outDir, exeName)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cmdRun := exec.Command(exePath, os.Args[2], cwd,os.Args[3])

	cmdRun.Dir = outDir
	cmdRun.Stdout = os.Stdout
	cmdRun.Stderr = os.Stderr

	err = cmdRun.Run()
	if err != nil {
		log.Fatalf("执行生成的 exe 失败: %v", err)
	}

	fmt.Println("生成的 exe 执行完成")

}
