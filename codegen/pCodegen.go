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

type PCodeGen struct {
	cwd            string
	PdslFile       string
	outDir         string
	nameWithoutExt string
	srcDir string
	workflow_name string
}

func NewPCodeGen(pdsl string) *PCodeGen {
	c, _ := os.Getwd()
	fmt.Println("pdslcode gening")
	return &PCodeGen{
		cwd:      c,
		PdslFile: pdsl,
		workflow_name: filepath.Base(c),   
	}
}
func NewPCodeGentest(pdsl string) *PCodeGen {
	c, _ := os.Getwd()

	fmt.Println("pdslcode gening")
	return &PCodeGen{
		cwd:      c,
		PdslFile: pdsl,
		workflow_name: "testpsdl",   
	}
}
// 验证参数合法性
func (p *PCodeGen) validateArgs() error {
	config.InitConfig(p.cwd)
	config.InitLogger()
	if len(os.Args) < 2 {
		log.Fatalf("请提供 pdsl 文件路径作为参数")
	}
	// 从命令行获取单个 pdsl 文件路径
	if !strings.HasSuffix(p.PdslFile, ".pdsl") {
		log.Fatalf("文件不是 .pdsl 文件: %s", p.PdslFile)
	}
	return nil
}
func (p *PCodeGen) initGenDirs() error {
	genDir := "generated_code"
	err := os.MkdirAll(genDir, os.ModePerm)
	if err != nil {
		log.Fatalf("创建生成目录失败: %v", err)
	}
	filename := filepath.Base(p.PdslFile)
	p.nameWithoutExt = strings.TrimSuffix(filename, ".pdsl")
	p.outDir = filepath.Join(genDir, p.nameWithoutExt)
	//子目录
	err = os.MkdirAll(p.outDir, os.ModePerm)
	if err != nil {
		log.Fatalf("创建子目录失败: %v", err)
	}
	return nil
}
func (p *PCodeGen) winitGenDirs() error {
	genDir := "../generated_code"
	err := os.MkdirAll(genDir, os.ModePerm)
	if err != nil {
		log.Fatalf("创建生成目录失败: %v", err)
	}
	filename := filepath.Base(p.PdslFile)
	p.nameWithoutExt = strings.TrimSuffix(filename, ".pdsl")
	p.outDir = filepath.Join(genDir,p.workflow_name)
	//子目录
	err = os.MkdirAll(p.outDir, os.ModePerm)
	if err != nil {
		log.Fatalf("创建子目录失败: %v", err)
	}
	return nil
}

func (p *PCodeGen) parserAndGen() error {
	// 读取 pdsl 文件
	content, err := os.ReadFile(p.PdslFile)
	if err != nil {
		log.Fatalf("读取 pdsl 文件失败: %v", err)
	}

	// 生成 Prompt
	ptc := NewPromptToGenCode(string(content), p.nameWithoutExt)
	err = ptc.PromptToGenCode()
	if err != nil {
		log.Fatalf("PromptToGenCode error: %v", err)
	}
	return nil
}
func (p *PCodeGen) parserAndGentest() error {
	// 读取 pdsl 文件
	content, err := os.ReadFile(p.PdslFile)
	if err != nil {
		log.Fatalf("读取 pdsl 文件失败: %v", err)
	}
	// 生成 Prompt
	ptc := NewPromptToGenCode(string(content), p.nameWithoutExt)
	err = ptc.PromptToGenCode()
	if err != nil {
		log.Fatalf("PromptToGenCode error: %v", err)
	}
	return nil
}
func (p *PCodeGen) wParserAndGen() error {
	// 读取 pdsl 文件
	content, err := os.ReadFile(p.PdslFile)
	if err != nil {
		log.Fatalf("读取 pdsl 文件失败: %v", err)
	}

	// 生成 Prompt
	ptc := NewPromptToGenCode(string(content), p.nameWithoutExt)
	err = ptc.WPromptToGenCode()
	if err != nil {
		log.Fatalf("PromptToGenCode error: %v", err)
	}
	return nil
}
func (p *PCodeGen) TempparserAndGen() error {
	p.srcDir = filepath.Dir(p.PdslFile)
	content, err := os.ReadFile(p.PdslFile)
	if err != nil {
		log.Fatalf("读取 pdsl 文件失败: %v", err)
	}
	filename := filepath.Base(p.PdslFile)
	p.nameWithoutExt = strings.TrimSuffix(filename, ".pdsl")
	// 生成 Prompt
	ptc := NewPromptToGenCode(string(content), p.nameWithoutExt)
	ptc.GenDirSet(p.srcDir)
	err = ptc.TempGoGen()
	if err != nil {
		log.Fatalf("PromptToGenCode error: %v", err)
	}
	return nil
}
func (p *PCodeGen) copyGoFiles() error {
	// 将 PdslFile 同目录下的 .go 文件和 exe 文件复制到生成的 exe 同目录下
	srcDir := filepath.Dir(p.PdslFile)

	// 先拷贝 .go 文件
	for _, filename := range config.Cfg.Utils {
		srcFile := filepath.Join(srcDir, filename)
		dstFile := filepath.Join(p.outDir, filename)
		data, err := os.ReadFile(srcFile)
		if err != nil {
			log.Fatalf("读取 %s 失败: %v", srcFile, err)
		}
		if err := os.WriteFile(dstFile, data, 0644); err != nil {
			log.Fatalf("写入 %s 失败: %v", dstFile, err)
		}
		log.Printf("已复制 %s 到: %s", srcFile, dstFile)
	}
	return nil
}
// 编译生成 exe
func (p *PCodeGen) buildExe() error {

	exeName := p.nameWithoutExt + ".exe"
	cmd := exec.Command("go", "build", "-o", exeName, ".")
	cmd.Dir = p.outDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(output))
		log.Fatalf("执行 go build 失败: %v", err)
	}
	log.Println("Go 程序编译完成，生成了", exeName)
	return nil
}
func (p *PCodeGen) Pcodegen() error {
	if err := p.validateArgs(); err != nil {
		return err
	}
	if err := p.initGenDirs(); err != nil {
		return err
	}
	if err := p.parserAndGen(); err != nil {
		return err
	}
	if err := p.copyGoFiles(); err != nil {
		return err
	}
	if err := p.buildExe(); err != nil {
		return err
	}
	return nil
}
func (p *PCodeGen) Wpcodegen() error {
	fmt.Println("Wpcodegen gening")
	if err := p.validateArgs(); err != nil {
		fmt.Println("err:",err)
		return err
	}
	if err := p.winitGenDirs(); err != nil {
		fmt.Println("err:",err)
		return err
	}
	if err := p.wParserAndGen(); err != nil {
		fmt.Println("err:",err)
		return err
	}
	if err := p.copyGoFiles(); err != nil {
		fmt.Println("err:",err)
		return err
	}
	fmt.Println("结束")
	// if err := p.buildExe(); err != nil {
	// 	return err
	// }
	return nil
}
func (p *PCodeGen) Wpcodegentest() error {
	fmt.Println("Wpcodegen gening")

	if err := p.winitGenDirs(); err != nil {
		fmt.Println("err:",err)
		return err
	}
	if err := p.wParserAndGen(); err != nil {
		fmt.Println("err:",err)
		return err
	}
	if err := p.copyGoFiles(); err != nil {
		fmt.Println("err:",err)
		return err
	}
	fmt.Println("结束")

	return nil
}