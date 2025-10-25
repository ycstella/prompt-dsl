package wcodegen

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ycstella/prompt-dsl/codegen"
)

type Wpipeline struct {
	InputFilePath string
	FileName      string
	Cmd           string
	w             WCodeGen
	PSDLFileList  []string
}

func NewWpipeline(fileName string) (*Wpipeline, error) {
	cmd, err := os.Getwd()
	return &Wpipeline{
		InputFilePath: "",
		FileName:      fileName,
		Cmd:           cmd,
	}, err
}
func (p *Wpipeline) GenW() (err error) {
	w := NewWCodeGen(p.FileName)
	p.w = w
	w.Execute()
	return err
}
func (p *Wpipeline) GenP() (err error) {
	p.PSDLFileList, err = ScanPdslFiles(p.Cmd)
	for _, f := range p.PSDLFileList {
		println("Found functional file:", f)
	}
	for _, v := range p.PSDLFileList {
		pdslPath := filepath.Join(p.Cmd, v+".pdsl")
		fmt.Println("文件：", pdslPath)
		p := codegen.NewPCodeGen(pdslPath)
		p.Wpcodegen()
	}
	return err
}

func ScanPdslFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(strings.ToLower(d.Name()), ".pdsl") {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func (p *Wpipeline) BuildExe() (err error)  {
	exeName := p.FileName + ".exe"
	cmd := exec.Command("go", "build", "-o", exeName, ".")
	cmd.Dir = ""
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(output))
		log.Fatalf("执行 go build 失败: %v", err)
	}
	log.Println("Go 程序编译完成，生成了", exeName)
	return err
}
func (p *Wpipeline) Execute() (err error) {
	//0.解析input
	
	//1.生成workflow.go文件
	if p.GenW() != nil {
		return err
	}
	//2.生成pdsl.go文件
	if p.GenP() != nil {
		return err
	}
	//3.go biuld exe
	if p.BuildExe()  != nil {
		return err
	}
	return err
}
