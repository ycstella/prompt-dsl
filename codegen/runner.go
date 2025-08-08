// promptdsl-core/runner.go
package codegen

import (
	"fmt"
	"log"
	"os"

	// "service"
	"codegen/parser"
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
	// 3. 构造 Eval 上下文
	str := &PromptGenContext{
		InFields:    rootNode.InFields,
		OutFields:   rootNode.OutFields,
		ModuleDefs:  rootNode.ModuleDefs,
		ModelFields: rootNode.ModelFields,
	}

	fmt.Println("😅ModelFields:", str.ModelFields)
	// fmt.Println("fmodeloutput")
	//3.5before
	// 3. 先执行 before 节点，填充 Vars
	// for _, node := range rootNode.BeforeNodes {
	// 	err := node.Tocode(str)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// }

	// 4. 执行 AST，得到 prompt 字符串
	outputParts, err := rootNode.Tocode(str)
	if err != nil {
		return nil, err
	}
	//生成sys+user+after+fix
	// code:=Generateprompthandle(rootNode, getCurrentPackageName())

	code := Generateprompthandle(rootNode, "generated", outputParts, filename, rootNode.Goimport)

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
