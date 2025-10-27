package wcodegen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	Wparser "github.com/ycstella/prompt-dsl/codegen/Wparser"
)

type WCodeGen struct {
	filename  string
	Dir       string
	ctx       *Wparser.WorkflowContext
	astNode   []*TaskNode
	topResult []*TaskNode
	codestr   string
}

func NewWCodeGen(filename string) WCodeGen {
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

func (w *WCodeGen) writewcode() error {
	w.filename = strings.TrimSuffix(filepath.Base(w.filename), filepath.Ext(w.filename))
	w.Dir,_=os.Getwd()
	w.Dir = strings.TrimSuffix(filepath.Base(w.Dir), filepath.Ext(w.Dir))
	genDir := filepath.Join("../generated_code",w.Dir)
	err := os.MkdirAll(genDir, os.ModePerm)
	if err != nil {
		fmt.Println("创建目录失败:v", err)
	}
	outputFile := filepath.Join(genDir, w.filename+".go")
	err = os.WriteFile(outputFile, []byte(w.codestr), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "写入文件失败: %v\n", err)
		os.Exit(1)
	}
	return nil
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
	gener := NewwcodeGener(w.astNode[0], w.topResult, w.ctx)
	gener.Execute()
	if err != nil {
		return err
	}
	w.codestr = gener.fullCode

	//写入
	w.writewcode()
	return err
}
