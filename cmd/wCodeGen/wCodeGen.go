package main

import (
	"fmt"
	"os"

	"github.com/ycstella/prompt-dsl/codegen"
)

func main() {
	//1.加载workflow.yaml
	w := codegen.NewWCodeGen()
	w.Workflowgen()
	fmt.Println(os.Args[1])
}
