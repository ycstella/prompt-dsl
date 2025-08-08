package main

import (
	// "encoding/json"
	"codegen"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"

	// "strings"

	// "net/http"
	// "codegen"
	"codegen/config"
	"os"
)

// var llm *service.LLMClient
func main() {
	config.InitConfig()
	config.InitLogger()
	// 指定 prompt 文件夹路径

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
	cmd := exec.Command("go", "build", "-o",exename, dir)
	cmd.Dir = dir
	// 获取命令的输出和错误信息
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("执行 go build 失败: %v\n", err)
	}
	// 打印命令执行的输出
	fmt.Println(string(cmdOutput))
	fmt.Println("Go 程序编译完成，生成了 output.exe 文件")

}

// 假设 result 是 string，包含 JSON 数组（即模型返回结果）
func CallUserPostProcessor[T any](jsonStr string) (T, error) {
	var zeroValue T // 用于返回零值

	fmt.Println("CallUserPostProcessor进入前")

	// 执行外部命令：运行 Go 程序
	cmd := exec.Command("go", "run", "code_gen/code_gen.go")
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return zeroValue, fmt.Errorf("无法获取 StdinPipe: %v", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return zeroValue, fmt.Errorf("无法获取 StdoutPipe: %v", err)
	}

	// 启动外部命令
	if err := cmd.Start(); err != nil {
		return zeroValue, fmt.Errorf("cmd.Start 错误: %v", err)
	}

	// 传入 JSON 输入
	_, err = stdin.Write([]byte(jsonStr))
	if err != nil {
		return zeroValue, fmt.Errorf("写入 stdin 错误: %v", err)
	}
	stdin.Close()

	// 读取命令输出
	outputBytes, err := io.ReadAll(stdout)
	if err != nil {
		return zeroValue, fmt.Errorf("读取 Stdout 错误: %v", err)
	}
	fmt.Println("outputBytes:", string(outputBytes))

	// 等待命令完成
	if err := cmd.Wait(); err != nil {
		return zeroValue, fmt.Errorf("等待命令完成时发生错误: %v", err)
	}

	// 将输出的 JSON 解析为目标类型
	var result T
	err = json.Unmarshal(outputBytes, &result)
	if err != nil {
		return zeroValue, fmt.Errorf("无法解析 AfterProcess 输出: %v", err)
	}

	return result, nil
}
