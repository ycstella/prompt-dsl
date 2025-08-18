// promptdsl-core/runner.go
package codegen

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	// "service"
	"github.com/along416/promptDSL/codegen/parser"
	// "strings"

	"github.com/antlr4-go/antlr/v4"
)

// 生成单prompt代码，返回user，sys
func RunPromptDSL(input string, filename string) (*final, error) {
	InitLog("llm.log")
	// 1. 解析输入 DSL 文本，生成 Parse Tree
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewPromptDSLLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPromptDSLParser(stream)
	tree := p.PromptFile()
	fmt.Println("🌳 start...")
	fmt.Println(tree.ToStringTree(nil, p))
	fmt.Println("🌳 ...end")
	// 2. 构建 AST Node

	ctx := tree.(*parser.PromptFileContext)

	rootNode := ConvertASTtoPrompt(ctx, stream, inputStream)
	// fmt.Printf("%v\n", rootNode)
	fmt.Printf("📦 OutFields: %+v\n", rootNode.OutFields)
	fmt.Printf("📦 aftercode: %+v\n", rootNode.AfterCode)
	fmt.Printf("📦 fixcode: %+v\n", rootNode.FixCode)
	fmt.Printf("📦 BeforeNodes: %+v\n", rootNode.BeforeCode)
	// 3. 构造 Eval 上下文
	str := &PromptGenContext{
		InFields:    rootNode.InFields,
		OutFields:   rootNode.OutFields,
		ModuleDefs:  rootNode.ModuleDefs,
		ModelFields: rootNode.ModelFields,
	}

	fmt.Println("😅ModelFields:", str.ModelFields)

	// 4. 执行 AST，得到 prompt 字符串
	outputParts, err := rootNode.Tocode(str)
	if err != nil {
		return nil, err
	}
	//生成sys+user+after+fix
	// code:=Generateprompthandle(rootNode, getCurrentPackageName())

	code := GeneratepromptCode(rootNode, "generated", outputParts, filename, rootNode.Goimport)

	genDir := "generated_code/" + filename
	err = os.MkdirAll(genDir, os.ModePerm)
	if err != nil {
		fmt.Println("创建目录失败: %v", err)
	}

	outputFile := genDir + "/main.go"
	err = installGoImports(rootNode.Goimport, "generated_code")
	if err != nil {
		log.Fatalf("安装依赖失败: %v", err)
	}
	err = os.WriteFile(outputFile, []byte(code), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "写入文件失败: %v\n", err)
		os.Exit(1)
	}

	os.Setenv("GOPROXY", "direct")
	os.Setenv("GOSUMDB", "off")

	// 检查go.mod文件是否存在，不存在才初始化
	// 假设 genDir 是你创建 go.mod 的目录
	goModPath := filepath.Join(genDir, "go.mod")
	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		// 使用合法 module 名
		modInitCmd := exec.Command("go", "mod", "init", "generated_code")
		modInitCmd.Dir = genDir
		if output, err := modInitCmd.CombinedOutput(); err != nil {
			fmt.Printf("go mod init 输出: %s\n", output)
			log.Fatalf("go mod init 失败: %v", err)
		}
		fmt.Println("成功初始化 Go module")
	} else if err != nil {
		log.Fatalf("检查 go.mod 文件失败: %v", err)
	} else {
		fmt.Println("go.mod 已存在，跳过初始化")
	}

	// 在 go.mod 所在目录执行 go get
	getCmd := exec.Command("go", "get", "github.com/along416/promptDSL@v0.1.13")
	getCmd.Dir = filepath.Dir(goModPath) // 指向 go.mod 所在目录
	log.Println("目录：", getCmd.Dir)
	if output, err := getCmd.CombinedOutput(); err != nil {
		fmt.Printf("go get 输出: %s\n", output)
		log.Fatalf("go get 失败: %v", err)
	}

	// 在 go.mod 所在目录执行 go mod tidy
	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Dir = filepath.Dir(goModPath) // 指向 go.mod 所在目录
	if output, err := tidyCmd.CombinedOutput(); err != nil {
		fmt.Printf("go mod tidy 输出: %s\n", output)
		log.Fatalf("go mod tidy 失败: %v", err)
	}

	fmt.Println("成功初始化 Go module 并安装依赖")

	return outputParts, nil
}
func WRunPromptDSL(input string, filename string, workflow string) (*final, error) {
	InitLog("llm.log")
	// 1. 解析输入 DSL 文本，生成 Parse Tree
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewPromptDSLLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPromptDSLParser(stream)
	tree := p.PromptFile()
	fmt.Println("🌳 start...")
	fmt.Println(tree.ToStringTree(nil, p))
	fmt.Println("🌳 ...end")
	// 2. 构建 AST Node

	ctx := tree.(*parser.PromptFileContext)

	rootNode := ConvertASTtoPrompt(ctx, stream, inputStream)
	// fmt.Printf("%v\n", rootNode)
	fmt.Printf("📦 OutFields: %+v\n", rootNode.OutFields)
	fmt.Printf("📦 aftercode: %+v\n", rootNode.AfterCode)
	fmt.Printf("📦 fixcode: %+v\n", rootNode.FixCode)
	fmt.Printf("📦 BeforeNodes: %+v\n", rootNode.BeforeCode)
	// 3. 构造 Eval 上下文
	str := &PromptGenContext{
		InFields:    rootNode.InFields,
		OutFields:   rootNode.OutFields,
		ModuleDefs:  rootNode.ModuleDefs,
		ModelFields: rootNode.ModelFields,
	}

	fmt.Println("😅ModelFields:", str.ModelFields)

	// 4. 执行 AST，得到 prompt 字符串
	outputParts, err := rootNode.Tocode(str)
	if err != nil {
		return nil, err
	}
	//生成sys+user+after+fix
	// code:=Generateprompthandle(rootNode, getCurrentPackageName())

	code := GeneratepromptCode(rootNode, "generated", outputParts, filename, rootNode.Goimport)

	outputFile := "generated_code/generated/" + filename + ".go"
	err = installGoImports(rootNode.Goimport, "generated_code")
	if err != nil {
		log.Fatalf("安装依赖失败: %v", err)
	}
	err = os.WriteFile(outputFile, []byte(code), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "写入文件失败: %v\n", err)
		os.Exit(1)
	}

	return outputParts, nil
}
