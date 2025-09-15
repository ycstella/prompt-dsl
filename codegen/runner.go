// promptdsl-core/runner.go
package codegen

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	// "service"
	"github.com/ycstella/prompt-dsl/codegen/parser"
	"github.com/ycstella/prompt-dsl/config"

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
	workflowName string
}

func NewPromptToGenCode(input, filename string) *promptToGenCode {
	cwd, _:= os.Getwd()
	return &promptToGenCode{
		input:    input,
		fileName: filename,
		workflowName:filepath.Base(cwd),
	}
}
func (ptc *promptToGenCode) parsePrompt() error {
	config.InitLogger()
	// 1. 解析输入 DSL 文本，生成 Parse Tree
	ptc.inputStream = antlr.NewInputStream(ptc.input)
	lexer := parser.NewPromptDSLLexer(ptc.inputStream)
	ptc.stream = antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPromptDSLParser(ptc.stream)
	ptc.tree = p.PromptFile()
	log.Println("🌳 start...")
	log.Println(ptc.tree.ToStringTree(nil, p))
	log.Println("🌳 ...end")
	return nil
}
func (ptc *promptToGenCode) astToNode() error {
	ctx := ptc.tree.(*parser.PromptFileContext)
	ptc.promptNode = ConvertASTtoPrompt(ctx, ptc.stream, ptc.inputStream)
	// fmt.Printf("%v\n", rootNode)
	log.Printf("📦 OutFields: %+v\n", ptc.promptNode.OutFields)
	log.Printf("📦 aftercode: %+v\n", ptc.promptNode.AfterCode)
	log.Printf("📦 fixcode: %+v\n", ptc.promptNode.FixCode)
	log.Printf("📦 BeforeNodes: %+v\n", ptc.promptNode.BeforeCode)
	return nil
}

// 准备use&sys的构造代码
func (ptc *promptToGenCode) buildPGCxtAndToCode() error {
	str := &PromptGenContext{
		InFields:    ptc.promptNode.InFields,
		OutFields:   ptc.promptNode.OutFields,
		ModuleDefs:  ptc.promptNode.ModuleDefs,
		ModelFields: ptc.promptNode.ModelFields,
	}

	fmt.Println("😅ModelFields:", str.ModelFields)
	// var userSysCode *final
	// log.Println("tocode处理问题")
	userSysCode, err := ptc.promptNode.Tocode(str)
	if err!= nil {
		log.Println("节点解析失败：",err)
		return err
	}
	ptc.codeGenUserAndSys = userSysCode
	return err
}

// 生成代码
func (ptc *promptToGenCode) genCode() error {
	CodeBuilder := NewCodeBuilder(ptc.promptNode, ptc.fileName, ptc.codeGenUserAndSys)
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
func (ptc *promptToGenCode) WgenCode() error {
	CodeBuilder := NewCodeBuilder(ptc.promptNode, ptc.fileName, ptc.codeGenUserAndSys)
	CodeBuilder.WcombineSingleCode()
	// code := GeneratepromptCode(ptc.promptNode, "generated", ptc.codeGenUserAndSys, ptc.fileName, ptc.promptNode.Goimport)

	ptc.genDir = filepath.Join("../generated_code",ptc.workflowName)
	err := os.MkdirAll(ptc.genDir, os.ModePerm)
	if err != nil {
		fmt.Println("创建目录失败: %v", err)
	}
	outputFile := filepath.Join(ptc.genDir,ptc.fileName+".go")
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

	os.Unsetenv("HTTP_PROXY")
	os.Unsetenv("HTTPS_PROXY")
	os.Setenv("GOPROXY", "https://goproxy.cn,direct")

	ptc.goModPath = filepath.Join(ptc.genDir, "go.mod")
	if _, err := os.Stat(ptc.goModPath); os.IsNotExist(err) {
		// 使用合法 module 名
		modInitCmd := exec.Command("go", "mod", "init", ptc.fileName)
		modInitCmd.Dir = ptc.genDir
		if output, err := modInitCmd.CombinedOutput(); err != nil {
			log.Printf("go mod init 输出: %s\n", output)
			log.Fatalf("go mod init 失败: %v", err)
		}
		log.Println("成功初始化 Go module")
	} else if err != nil {
		log.Fatalf("检查 go.mod 文件失败: %v", err)
	} else {
		log.Println("go.mod 已存在，跳过初始化")
	}
	return nil
}

// 在 go.mod 所在目录执行 go get
func (ptc *promptToGenCode) runGoGet() error {
    // 设置私有仓库，避免 go.sum 校验
    os.Setenv("GOPRIVATE", "github.com/ycstella/*")
    os.Setenv("GONOSUMDB", "github.com/ycstella/*")

    // 使用 token 拉取私有模块
    getCmd := exec.Command("go", "get", "github.com/ycstella/prompt-dsl@v0.1.37")
    getCmd.Dir = filepath.Dir(ptc.goModPath)
    getCmd.Env = append(os.Environ(),
        "GOPRIVATE=github.com/ycstella/*",
        "GONOSUMDB=github.com/ycstella/*",
    )

    log.Println("执行目录：", getCmd.Dir)

    output, err := getCmd.CombinedOutput()
    log.Printf("go get 输出:\n%s", string(output))
    if err != nil {
        log.Printf("go get 失败: %v", err)
        return err
    }

    log.Println("go get 成功 ✅")
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
//只需要生成代码
func (ptc *promptToGenCode) WPromptToGenCode() error {
	if err := ptc.parsePrompt(); err != nil {
		return err
	}
	if err := ptc.astToNode(); err != nil {
		return err
	}
	if err := ptc.buildPGCxtAndToCode(); err != nil {
		return err
	}
	if err := ptc.WgenCode(); err != nil {
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
func (ptc *promptToGenCode) tempGenCode() error {
	var temp *final
	CodeBuilder := NewCodeBuilder(ptc.promptNode, ptc.fileName, temp)
	CodeBuilder.TcombineSingleCode()
	// code := GeneratepromptCode(ptc.promptNode, "generated", ptc.codeGenUserAndSys, ptc.fileName, ptc.promptNode.Goimport)
	err := os.WriteFile(filepath.Join(ptc.genDir, ptc.fileName+".go"), []byte(CodeBuilder.b.String()), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "偷偷编译失败: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "偷偷编译成功: \n")
	return err
}
func (ptc *promptToGenCode) GenDirSet(g string) error {
	ptc.genDir = g
	return nil
}
func (ptc *promptToGenCode) TempGoGen() error {
	if err := ptc.parsePrompt(); err != nil {
		return err
	}
	if err := ptc.astToNode(); err != nil {
		return err
	}
	if err := ptc.buildPGCxtAndToCode(); err != nil {
		return err
	}
	if err := ptc.tempGenCode(); err != nil {
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
