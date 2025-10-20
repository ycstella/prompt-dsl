package wcodegen

import "testing"

func Test_wcodegen(t *testing.T) {
	logdfine()
	w:=NewWCodeGen("workflowSmp.txt")
	w.Execute()
}
