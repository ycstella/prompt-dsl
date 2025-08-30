package main

import (
	// "encoding/json"

	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ycstella/prompt-dsl/codegen"
	"github.com/ycstella/prompt-dsl/config"
)

type WCodeGen struct {
	cwd            string
	pdslFile       string
	outDir         string
	nameWithoutExt string
	srcDir string
}

func NewWCodeGen(pdsl string) *WCodeGen {
	c, _ := os.Getwd()
	return &WCodeGen{
		cwd:      c,
		pdslFile: pdsl,
	}
}

// 验证参数合法性
func (p *WCodeGen) validateArgs() error {
	config.InitConfig(p.cwd)
	config.InitLogger()
	if len(os.Args) < 2 {
		log.Fatalf("请提供 pdsl 文件路径作为参数")
	}
	// 从命令行获取单个 pdsl 文件路径
	if !strings.HasSuffix(p.pdslFile, ".pdsl") {
		log.Fatalf("文件不是 .pdsl 文件: %s", p.pdslFile)
	}
	return nil
}
func (p *WCodeGen) initGenDirs() error {
	genDir := "generated_code"
	err := os.MkdirAll(genDir, os.ModePerm)
	if err != nil {
		log.Fatalf("创建生成目录失败: %v", err)
	}

	filename := filepath.Base(p.pdslFile)
	p.nameWithoutExt = strings.TrimSuffix(filename, ".pdsl")
	p.outDir = filepath.Join(genDir, p.nameWithoutExt)
	//子目录
	err = os.MkdirAll(p.outDir, os.ModePerm)
	if err != nil {
		log.Fatalf("创建子目录失败: %v", err)
	}
	return nil
}
func (p *WCodeGen) parserAndGen() error {
	// 读取 pdsl 文件
	content, err := os.ReadFile(p.pdslFile)
	if err != nil {
		log.Fatalf("读取 pdsl 文件失败: %v", err)
	}

	// 生成 Prompt
	ptc := codegen.NewPromptToGenCode(string(content), p.nameWithoutExt)
	err = ptc.PromptToGenCode()
	if err != nil {
		log.Fatalf("PromptToGenCode error: %v", err)
	}
	return nil
}
func (p *WCodeGen) tempparserAndGen() error {
	p.srcDir = filepath.Dir(p.pdslFile)
	content, err := os.ReadFile(p.pdslFile)
	if err != nil {
		log.Fatalf("读取 pdsl 文件失败: %v", err)
	}
	filename := filepath.Base(p.pdslFile)
	p.nameWithoutExt = strings.TrimSuffix(filename, ".pdsl")
	// 生成 Prompt
	ptc := codegen.NewPromptToGenCode(string(content), p.nameWithoutExt)
	ptc.GenDirSet(p.srcDir)
	err = ptc.TempGoGen()
	if err != nil {
		log.Fatalf("PromptToGenCode error: %v", err)
	}
	return nil
}
func (p *WCodeGen) copyGoFiles() error {
	// 将 pdslFile 同目录下的 .go 文件和 exe 文件复制到生成的 exe 同目录下
	srcDir := filepath.Dir(p.pdslFile)

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

	// // 再拷贝 exe 文件
	// for _, exe := range exeNames {
	// 	srcExe := filepath.Join(srcDir, exe)
	// 	dstExe := filepath.Join(p.outDir, exe)
	// 	data, err := os.ReadFile(srcExe)
	// 	if err != nil {
	// 		log.Fatalf("读取 %s 失败: %v", srcExe, err)
	// 	}
	// 	if err := os.WriteFile(dstExe, data, 0755); err != nil { // exe 权限
	// 		log.Fatalf("写入 %s 失败: %v", dstExe, err)
	// 	}
	// 	log.Printf("已复制 %s 到: %s", srcExe, dstExe)
	// }

	return nil
}
// 编译生成 exe
func (p *WCodeGen) buildExe() error {

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
func (p *WCodeGen) WCodeGen() error {
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

//用workflow1的路径去掉最后生成文件夹路径，在文件夹里查找所有的.pdsl文件去生成所有的.go文件，
//集成之前的copy文件，gomod，goget
//go文件不需要main，package也叫main
//查找wpdsl文件生成对应workflow1.go文件作为main去调用psdl文件生成的gowenjian
func main() {
	p := NewWCodeGen(os.Args[1])
	log.Println("参数数量：", len(os.Args))
	if len(os.Args) < 4 {
		err := p.WCodeGen()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("偷偷编译中,路径：", filepath.Dir(p.pdslFile))
		config.InitConfig(filepath.Dir(p.pdslFile))
		config.InitLogger()
		err := p.tempparserAndGen()
		if err != nil {
			log.Fatal(err)
		}
	}
}
