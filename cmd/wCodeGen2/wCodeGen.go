package main

import (
	"log"
	"os"

	wcodegen "github.com/ycstella/prompt-dsl/codegen/wCodegen"
)

func main() {
	wcodegen.Logdfine()
	//1.加载workflow.yaml
	w,err := wcodegen.NewWpipeline(os.Args[1])
	if err != nil {
		log.Fatal()
	}
	w.Execute()
}
