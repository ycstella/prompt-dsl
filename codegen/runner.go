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
type promptToGenCode struct {
	input             string
	fileName          string
	tree              parser.IPromptFileContext
	promptNode        *PromptNode
	inputStream       *antlr.InputStream
	stream            *antlr.CommonTokenStream
	codeGenUserAndSys *final
	genDir            string
	goModPath         string
}

func NewPromptToGenCode(input, filename string) *promptToGenCode {
	return &promptToGenCode{
		input:    input,
		fileName: filename,
	}
}
func (ptc *promptToGenCode) parsePrompt() error {
	InitLog("llm.log")
	// 1. 解析输入 DSL 文本，生成 Parse Tree
	ptc.inputStream = antlr.NewInputStream(ptc.input)
	lexer := parser.NewPromptDSLLexer(ptc.inputStream)
	ptc.stream = antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPromptDSLParser(ptc.stream)
	ptc.tree = p.PromptFile()
	fmt.Println("🌳 start...")
	fmt.Println(ptc.tree.ToStringTree(nil, p))
	fmt.Println("🌳 ...end")
	return nil
}
func (ptc *promptToGenCode) astToNode() error {
	ctx := ptc.tree.(*parser.PromptFileContext)
	ptc.promptNode = ConvertASTtoPrompt(ctx, ptc.stream, ptc.inputStream)
	// fmt.Printf("%v\n", rootNode)
	fmt.Printf("📦 OutFields: %+v\n", ptc.promptNode.OutFields)
	fmt.Printf("📦 aftercode: %+v\n", ptc.promptNode.AfterCode)
	fmt.Printf("📦 fixcode: %+v\n", ptc.promptNode.FixCode)
	fmt.Printf("📦 BeforeNodes: %+v\n", ptc.promptNode.BeforeCode)
	return nil
}

func (ptc *promptToGenCode) buildPGCxtAndToCode() error {
	str := &PromptGenContext{
		InFields:    ptc.promptNode.InFields,
		OutFields:   ptc.promptNode.OutFields,
		ModuleDefs:  ptc.promptNode.ModuleDefs,
		ModelFields: ptc.promptNode.ModelFields,
	}

	fmt.Println("😅ModelFields:", str.ModelFields)
	var userSysCode *final
	userSysCode, err := ptc.promptNode.Tocode(str)
	ptc.codeGenUserAndSys = userSysCode
	if err != nil {
		return err
	}
	return nil
}

// 生成代码
func (ptc *promptToGenCode) genCode() error {
	CodeBuilder:=NewCodeBuilder(ptc.promptNode, ptc.fileName, ptc.codeGenUserAndSys)
	CodeBuilder.combineSingleCode()
	// code := GeneratepromptCode(ptc.promptNode, "generated", ptc.codeGenUserAndSys, ptc.fileName, ptc.promptNode.Goimport)

	ptc.genDir = "generated_code/" + ptc.fileName
	err := os.MkdirAll(ptc.genDir, os.ModePerm)
	if err != nil {
		fmt.Println("创建目录失败: %v", err)
	}

	outputFile := ptc.genDir + "/main.go"
	err = installGoImports(ptc.promptNode.Goimport, ptc.fileName)
	if err != nil {
		log.Fatalf("安装依赖失败: %v", err)
	}
	err = os.WriteFile(outputFile, []byte(CodeBuilder.b.String()), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "写入文件失败: %v\n", err)
		os.Exit(1)
	}
	return err
}

// 检查go.mod文件是否存在，不存在则初始化
func (ptc *promptToGenCode) ensureGoModule() error {

	os.Setenv("GOPROXY", "https://goproxy.cn,direct")

	ptc.goModPath = filepath.Join(ptc.genDir, "go.mod")
	if _, err := os.Stat(ptc.goModPath); os.IsNotExist(err) {
		// 使用合法 module 名
		modInitCmd := exec.Command("go", "mod", "init", ptc.fileName)
		modInitCmd.Dir = ptc.genDir
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
	return nil
}

// 在 go.mod 所在目录执行 go get
func (ptc *promptToGenCode) runGoGet() error {
	// 在 go.mod 所在目录执行 go get
	getCmd := exec.Command("go", "get", "github.com/along416/promptDSL@v0.1.13")
	getCmd.Dir = filepath.Dir(ptc.goModPath)
	log.Println("目录：", getCmd.Dir)
	if output, err := getCmd.CombinedOutput(); err != nil {
		fmt.Printf("go get 输出: %s\n", output)
		log.Fatalf("go get 失败: %v", err)
	}

	return nil
}

// 在 go.mod 所在目录执行 go mod tidy
func (ptc *promptToGenCode) runGoModTidy() error {

	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Dir = filepath.Dir(ptc.goModPath)
	if output, err := tidyCmd.CombinedOutput(); err != nil {
		fmt.Printf("go mod tidy 输出: %s\n", output)
		log.Fatalf("go mod tidy 失败: %v", err)
	}

	fmt.Println("成功初始化 Go module 并安装依赖")
	return nil
}
func (ptc *promptToGenCode) PromptToGenCode() error {
	if err := ptc.parsePrompt(); err != nil {
		return err
	}
	if err := ptc.astToNode(); err != nil {
		return err
	}
	if err := ptc.buildPGCxtAndToCode(); err != nil {
		return err
	}
	if err := ptc.genCode(); err != nil {
		return err
	}
	if err := ptc.ensureGoModule(); err != nil {
		return err
	}
	if err := ptc.runGoGet(); err != nil {
		return err
	}
	if err := ptc.runGoModTidy(); err != nil {
		return err
	}
	return nil
}

// func WRunPromptDSL(input string, filename string, workflow string) (*final, error) {
// 	InitLog("llm.log")
// 	// 1. 解析输入 DSL 文本，生成 Parse Tree
// 	inputStream := antlr.NewInputStream(input)
// 	lexer := parser.NewPromptDSLLexer(inputStream)
// 	stream := antlr.NewCommonTokenStream(lexer, 0)
// 	p := parser.NewPromptDSLParser(stream)
// 	tree := p.PromptFile()
// 	fmt.Println("🌳 start...")
// 	fmt.Println(tree.ToStringTree(nil, p))
// 	fmt.Println("🌳 ...end")
// 	// 2. 构建 AST Node

// 	ctx := tree.(*parser.PromptFileContext)

// 	rootNode := ConvertASTtoPrompt(ctx, stream, inputStream)
// 	// fmt.Printf("%v\n", rootNode)
// 	fmt.Printf("📦 OutFields: %+v\n", rootNode.OutFields)
// 	fmt.Printf("📦 aftercode: %+v\n", rootNode.AfterCode)
// 	fmt.Printf("📦 fixcode: %+v\n", rootNode.FixCode)
// 	fmt.Printf("📦 BeforeNodes: %+v\n", rootNode.BeforeCode)
// 	// 3. 构造 Eval 上下文
// 	str := &PromptGenContext{
// 		InFields:    rootNode.InFields,
// 		OutFields:   rootNode.OutFields,
// 		ModuleDefs:  rootNode.ModuleDefs,
// 		ModelFields: rootNode.ModelFields,
// 	}

// 	fmt.Println("😅ModelFields:", str.ModelFields)

// 	// 4. 执行 AST，得到 prompt 字符串
// 	outputParts, err := rootNode.Tocode(str)
// 	if err != nil {
// 		return nil, err
// 	}
// 	//生成sys+user+after+fix
// 	// code:=Generateprompthandle(rootNode, getCurrentPackageName())
	
// 	code := GeneratepromptCode(rootNode, "generated", outputParts, filename, rootNode.Goimport)

// 	outputFile := "generated_code/generated/" + filename + ".go"
// 	err = installGoImports(rootNode.Goimport, "generated_code")
// 	if err != nil {
// 		log.Fatalf("安装依赖失败: %v", err)
// 	}
// 	err = os.WriteFile(outputFile, []byte(code), 0644)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "写入文件失败: %v\n", err)
// 		os.Exit(1)
// 	}

// 	return outputParts, nil
// }
