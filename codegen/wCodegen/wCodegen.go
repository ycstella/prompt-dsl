package wcodegen

import (
	"log"

	"github.com/antlr4-go/antlr/v4"
	Wparser "github.com/ycstella/prompt-dsl/codegen/Wparser"
)

type WCodeGen struct {
	filename  string
	ctx       *Wparser.WorkflowContext
	astNode   []*TaskNode
	topResult []*TaskNode
	codestr   string
}
func NewWCodeGen (filename string) WCodeGen{
	return WCodeGen{
		filename: filename,
	}
}
// 解析文件内容
func (w *WCodeGen) parseWDSL() (err error) {
	input, err := antlr.NewFileStream(w.filename)
	if err != nil {
		log.Fatal(err)
	}
	lexer := Wparser.NewWorkflowDSLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := Wparser.NewWorkflowDSLParser(stream)
	tree := parser.Workflow()
	ctx, ok := tree.(*Wparser.WorkflowContext)
	if !ok {
		log.Fatal("类型转换失败")
	}
	w.ctx = ctx
	return err
}
func (w *WCodeGen) Execute() (err error) {
	//解析文件内容
	err = w.parseWDSL()
	if err != nil {
		return err
	}
	//构造ast
	ast := NewWAstbuilder(w.ctx)
	err = ast.astbuild()
	if err != nil {
		return err
	}
	w.astNode = ast.TaskList
	//top排序
	w.topResult, err = topoSort(w.astNode)
	//生成最终代码
	gener := NewwcodeGener(w.astNode[0],w.topResult,w.ctx)
	gener.Execute()
	if err != nil {
		return err
	}
	w.codestr=gener.fullCode
	return err
}
