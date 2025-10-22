package wcodegen

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	Wparser "github.com/ycstella/prompt-dsl/codegen/Wparser"
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
