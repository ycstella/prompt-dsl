package main

import (
	// "encoding/json"

	"log"
	"os"
	"path/filepath"

	"github.com/ycstella/prompt-dsl/codegen"
	"github.com/ycstella/prompt-dsl/config"
)

func main() {
	p := codegen.NewPCodeGen(os.Args[1])
	log.Println("参数数量：", len(os.Args))
	if len(os.Args) < 4 {
		err := p.Pcodegen()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("偷偷编译中,路径：", filepath.Dir(p.PdslFile))
		config.InitConfig(filepath.Dir(p.PdslFile))
		config.InitLogger()
		err := p.TempparserAndGen()
		if err != nil {
			log.Fatal(err)
		}
	}
}
