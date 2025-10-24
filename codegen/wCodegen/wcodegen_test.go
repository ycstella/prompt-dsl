package wcodegen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/ycstella/prompt-dsl/codegen"
	Wparser "github.com/ycstella/prompt-dsl/codegen/Wparser"
	gened "github.com/ycstella/prompt-dsl/codegen/wcodegen/generated"
	"github.com/ycstella/prompt-dsl/config"
)

func Test_wcodegen(t *testing.T) {
	logdfine()
	w := NewWCodeGen("workflowSmp.txt")
	w.Execute()
}
func Test_genPDSL(t *testing.T) {
	logdfine()
	b, err := ioutil.ReadFile("workflowSmp.txt")
	if err != nil {
		log.Print(err)
	}
	g := NewWPDSLGener(string(b), "./generated")
	g.Execute()
}
func Test_outextract(t *testing.T) {
	input := `workflow {
  in1 {
      Id: int
      Question: string
      Solution: string
  }
  out { 
      Id:int                  <-in1.Id
      Question: string        <-in1.Question
      Solution: string        <-in1.Solution
      Process:struct{
          Id:int
          StepsOut:struct{
              Step:  int  
              StepContent: string 
              Path:struct{
                  FromePhenomena:struct{ 
                      Phenomenon: string
                      Projection: string
                      Concept: string
                      IsDerived: string
                  }
                  Expression: string
                  ToPhenomena:string
              }   
          }    <-Getpath(ExtractSteps(in1))
      }     <-ExtractSteps(in1)                
  }     
  models {
    default: 0
    getId: 1              
    ExtractSteps: 2                  
  }
}`

	is := antlr.NewInputStream(input)
	lexer := Wparser.NewWorkflowDSLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := Wparser.NewWorkflowDSLParser(stream)

	tree := p.Workflow()

	listener := NewArrowStructListener("Getpath", input)
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	fmt.Println(listener.result)
}

func Test_pdslgencode(t *testing.T) {
	// fmt.Println("1111111111")
	logdfine()
	w := codegen.NewWCode2()
	// fmt.Println("1111111111")
	filelist := []string{"testpsdl\\ExtractSteps", "testpsdl\\Getpath"}
	w.Looptask(filelist)
	// fmt.Println("1111111111")
}
func Test_runtest(t *testing.T) {
	// 执行依赖计算
	var in1 gened.In1
	b, err := ioutil.ReadFile("generated/test.json")
	if err != nil {
		log.Print(err)
	}
	err = json.Unmarshal(b, &in1)
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("u: %+v\n", in1)
	cd, _ := os.Getwd()
	fmt.Println("当前路径：", cd)

	config.WF.Model = "gemini-2.0-flash"
	config.WF.Config = "./generated"
	gened.Run(in1)
}
