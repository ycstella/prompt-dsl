package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"

	// "net/http"
	"os"
	"promptdslcore" 
	"promptdslcore/config"
)

// var llm *service.LLMClient
func main() {
	config.InitConfig()
	// 打开或创建日志文件
	// logFile, err := os.OpenFile("llm.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("无法打开日志文件: %v", err)
	// }
	// // 设置日志输出到文件
	// log.SetOutput(logFile)
	// // 可选: 添加时间戳、文件名等信息
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	config.InitLogger()
	// log.Println()
	// 指定 prompt 文件夹路径
	promptDir := "../promptdslcore/prompts_copy"
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
		path := "../promptdslcore/prompts_copy/" + filename
		// fileContent, err := os.ReadFile("./promptdsl-core/prompts/SplitSolutionSteps.prompt")
		promptfileContent, err := os.ReadFile(path)
		if err != nil {
			panic(fmt.Errorf("读取 DSL 文件失败: %v", err))
		}
		
		promptfileContentstr := string(promptfileContent)
		nameWithoutExt := strings.TrimSuffix(filename, ".pdsl")
		
		prompt, err := promptdslcore.RunPromptDSL(promptfileContentstr, nameWithoutExt)
		if err != nil {
			log.Fatalf("RunPromptDSL error: %v", err)
		}
		log.Println("生成的 Prompt:\n", prompt)
	}
	// code := promptdslcore.Generatworkflow("generated")
	// outputFile := "../generated_code/workflow.go"
	// err = os.WriteFile(outputFile, []byte(code), 0644)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "写入workflow文件失败: %v\n", err)
	// 	os.Exit(1)
	// }
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
