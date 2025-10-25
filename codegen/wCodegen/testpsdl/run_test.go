package test

import (
	"log"
	"testing"
	"github.com/ycstella/prompt-dsl/codegen/wCodegen"
)

func Test_pipe(t *testing.T) {
	// fmt.Println("11111111111")
	// logdfine()
	//1.加载workflow
	w, err := wcodegen.NewWpipeline("./workflowSmp.txt")
	if err != nil {
		log.Fatal()
	}
	w.Execute()
}
