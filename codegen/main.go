package codegen

// import (
// 	"fmt"
// 	"os"
// 	"codegen/parser"
// 	"github.com/antlr4-go/antlr/v4"
// 	// "github.com/dop251/goja"
// )

// func main() {

// 	fileContent, err := os.ReadFile("./example.prompt") // 文件扩展名你自己定义
// 	if err != nil {
// 		panic(fmt.Errorf("读取 DSL 文件失败: %v", err))
// 	}
// 	input := string(fileContent)
// 	// 创建输入流
// 	inputStream := antlr.NewInputStream(input)
// 	lexer := parser.NewPromptDSLLexer(inputStream)
// 	tokenStream := antlr.NewCommonTokenStream(lexer, 0)

// 	tokenStream.Fill()
// 	tokens := tokenStream.GetAllTokens() // 使用 GetAllTokens() 而不是 GetTokens()
// 	for _, token := range tokens {
// 		if token.GetTokenType() != antlr.TokenEOF {
// 			fmt.Printf("Token: %s, text: %q\n",
// 				lexer.SymbolicNames[token.GetTokenType()],
// 				token.GetText())
// 		}
// 	}

// 	// 解析，调用顶层规则
// 	p := parser.NewPromptDSLParser(tokenStream)
// 	tree := p.PromptFile()
// 	// 打印语法树文本结构
// 	fmt.Println(tree.ToStringTree(nil, p))

// 	// 创建访问者，访问解析树

// 	visitor := parser.NewMyVisitor()
// 	// visitor := &promptdsl.BasePromptDSLVisitor{}
// 	fmt.Println("🌳 start...")

// 	visitor.Visit(tree)
// 	fmt.Println("🌳 ...end")
// 	// 生成最终的 Prompt 字符串
// 	finalPrompt := visitor.GeneratePrompt()

// 	// 输出生成的 Prompt
// 	fmt.Println(finalPrompt)
// 	// fmt.Println(tree.ToStringTree(nil, parser))
// }
