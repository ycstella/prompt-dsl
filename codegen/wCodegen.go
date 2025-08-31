package codegen

import (
	// "encoding/json"

	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ycstella/prompt-dsl/config"
)

type WCodeGen struct {
	cwd      string
	workflow config.Workflow
}

func NewWCodeGen() *WCodeGen {
	c, _ := os.Getwd()
	filename := ""
	config.InitWorkflow("", filename)
	return &WCodeGen{
		cwd:      c,
		workflow: config.WF,
	}
}
func (w *WCodeGen) looptask() error {
	for _, task := range w.workflow.Task {
		pdslPath := filepath.Join(w.cwd, task)
		p := NewPCodeGen(pdslPath)
		p.pcodegen()
	}
	return nil
}

func (w *WCodeGen) wgen() error {
	for _, task := range w.workflow.Task {

	}
	return nil
}
func (w *WCodeGen) buildExe() error {

	exeName := w.nameWithoutExt + ".exe"
	cmd := exec.Command("go", "build", "-o", exeName, ".")
	cmd.Dir = w.outDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(output))
		log.Fatalf("执行 go build 失败: %v", err)
	}
	log.Println("Go 程序编译完成，生成了", exeName)
	return nil
}
func (w *WCodeGen) executeExe() error {
	return nil
}
