package codegen

import (
	// "encoding/json"

	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ycstella/prompt-dsl/config"
)

type WCodeGen struct {
	cwd      string
	workflow config.Workflow
	fileName string
	genDir   string
	b        strings.Builder
}

func NewWCodeGen() *WCodeGen {
	c, _ := os.Getwd()
	filename := os.Args[1]
	config.InitWorkflow(c, filename)
	return &WCodeGen{
		cwd:      c,
		fileName: filename,
		workflow: config.WF,
	}
}
func NewWCode2( ) *WCodeGen {
	c, _ := os.Getwd()
	return &WCodeGen{
		cwd:      c,
	}
}
func (w *WCodeGen) looptask() error {
	for _, task := range w.workflow.Task {
		pdslPath := filepath.Join(w.cwd, task+".pdsl")
		p := NewPCodeGen(pdslPath)
		p.Wpcodegen()
	}
	return nil
}
func (w *WCodeGen) Looptask(list []string ) error {
	for _, v := range list {
		// fmt.Println("进入:",v)
		pdslPath := filepath.Join(w.cwd, v +".pdsl")
		fmt.Println("文件：",pdslPath)
		p := NewPCodeGen(pdslPath)
		p.Wpcodegen()
	}
	return nil
}
func (w *WCodeGen) wgen() error {
	w.b.WriteString("package main\n")
	w.b.WriteString("import (\n")
	w.b.WriteString("\t\"github.com/ycstella/prompt-dsl/codegen\"\n")
	w.b.WriteString("\t\"github.com/ycstella/prompt-dsl/config\"\n")
	w.b.WriteString("\t\"os\"\n")
	w.b.WriteString("\t\"log\"\n")
	w.b.WriteString(")\n")
	w.b.WriteString("func main() {\n")
	w.b.WriteString("\tconfig.InitWorkflow(os.Args[1],os.Args[2])\n")
	for i, task := range w.workflow.Task {
		px := fmt.Sprintf("p%d", i+1)
		w.b.WriteString(fmt.Sprintf("    %s := New%s()\n", px, task))
		w.b.WriteString(fmt.Sprintf("    %s.init()\n", px))
	}
	firstPx := "p1"
	w.b.WriteString(fmt.Sprintf("    %s.loadInput()\n", firstPx))
	w.b.WriteString(fmt.Sprintf("    %s.parsedata()\n", firstPx))
	w.b.WriteString("    var err error\n")
	w.b.WriteString(fmt.Sprintf("    for _, item := range %s.Input {\n", firstPx))
	w.b.WriteString(fmt.Sprintf("        %s.currentData = item\n", firstPx))

	for i := 0; i < len(w.workflow.Task); i++ {
		curr := fmt.Sprintf("p%d", i+1)
		w.b.WriteString(fmt.Sprintf("    \t%s.singleExecute()\n", curr))
		// 如果不是最后一个，拷贝结果给下一个
		if i < len(w.workflow.Task)-1 {
			next := fmt.Sprintf("p%d", i+2)
			w.b.WriteString(fmt.Sprintf("    \terr=codegen.CopyStructRecursive(&%s.afterRet, &%s.currentData)\n", curr, next))
			w.b.WriteString("    \tif err != nil {\n")
			w.b.WriteString("    \t\tlog.Fatal(err)\n")
			w.b.WriteString("    \t}\n")
		}
	}
	lastPx := fmt.Sprintf("p%d", len(w.workflow.Task))
	w.b.WriteString(fmt.Sprintf("        %s.writeOut(%s.afterRet)\n", lastPx, lastPx))
	w.b.WriteString("    }\n")
	w.b.WriteString("}\n")
	w.writewcode()

	return nil
}
func (w *WCodeGen) writewcode() error {

	w.genDir = filepath.Join("../generated_code", w.fileName)
	err := os.MkdirAll(w.genDir, os.ModePerm)
	if err != nil {
		fmt.Println("创建目录失败: ", err)
	}
	outputFile := filepath.Join(w.genDir, w.fileName+".go")
	err = os.WriteFile(outputFile, []byte(w.b.String()), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "写入文件失败: \n", err)
		os.Exit(1)
	}
	return nil
}
func (w *WCodeGen) buildExe() error {
	exeName := w.fileName + ".exe"
	cmd := exec.Command("go", "build", "-o", exeName, ".")
	cmd.Dir = w.genDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(output))
		log.Fatalf("执行 go build 失败: %v", err)
	}
	log.Println("Go 程序编译完成，生成了", exeName)
	return nil
}

func (w *WCodeGen) Workflowgen() error {
	//2.循环task中所有pdsl文件去生成针对workflow的go文件（没有main）
	w.looptask()
	//3.生成workflow.go文件
	w.wgen()
	// 4.buildexe，
	w.buildExe()
	// 5.执行exe
	// w.runExe()
	//循环调用每一个task对应的exe,但是task是input的内部循环，所以task内部的调用要改
	// for _,t:=range w.workflow.Task{
	// }
	return nil
}
