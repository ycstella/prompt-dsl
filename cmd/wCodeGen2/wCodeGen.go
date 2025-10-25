package main

import (
	"log"

	wcodegen "github.com/ycstella/prompt-dsl/codegen/wCodegen"
)

func main() {
	wcodegen.Logdfine()
	//1.加载workflow.yaml
	w,err := wcodegen.NewWpipeline()
	if err != nil {
		log.Fatal()
	}
	w.Execute()
}
