package wcodegen

import (
	"fmt"
	"log"
	"strings"

	Wparser "github.com/ycstella/prompt-dsl/codegen/Wparser"
)

type Generater struct {
	astNode   *TaskNode
	topResult []*TaskNode
	structs   string
	funcs     string
	exec      string
	fullCode  string
	importstr string
	util      string
	main      string
	ctx       *Wparser.WorkflowContext
}

func NewwcodeGener(astNode *TaskNode, topResult []*TaskNode, ctx *Wparser.WorkflowContext) Generater {
	return Generater{
		astNode:   astNode,
		topResult: topResult,
		ctx:       ctx,
	}
}

// 变量管理
func (g *Generater) assignNames() (err error) {
	counter := 1
	for _, n := range g.topResult {
		// 给每个节点一个唯一变量名
		n.VarName = fmt.Sprintf("v%d", counter)
		counter++

		// 根据类型决定函数名
		switch n.Type {
		case Func:
			n.CallName = extractFuncName(n.Expr)
		case Arithmetic:
			n.CallName = fmt.Sprintf("Arith_%d", counter)
		case Conditional:
			n.CallName = fmt.Sprintf("Conditional_%d", counter)
		case Atom:
			n.CallName = n.Expr
		case FAtom:
			filed := getLastPart(n.Expr)
			if len(n.Dependencies) > 0 {
				depId := n.Dependencies[0].Id
				depVar := findVarNameById(depId, g.topResult)
				n.CallName = depVar + "." + filed
			} else {
				n.CallName = filed
			}

		}

	}
	return err
}

func (g *Generater) genImport() string {
	return `
package main
import (
	"log"
	"io/ioutil"
	// "github.com/ycstella/prompt-dsl/codegen"
	// "github.com/spf13/cast"
	"github.com/ycstella/prompt-dsl/config"
	"os"
	"encoding/json"
	"fmt"
)`
}
func (g *Generater) genStructs() string {
	w := BuildModel(g.ctx)
	structcode := GenerateCode(w)
	return structcode
}

// 直连单个task的调用
func (g *Generater) genFuncs() string {

	const tmplStr = `
{{- range .}}
	{{- if eq .Type 2 }}  
func W{{ .CallName }}(
	{{- range $j, $d := .Dependencies }}
		{{- if $j}}, {{ end }}{{ findVarNameById $d.Id $ }} any
	{{- end }}) float64 {
	return {{ range $j, $d := .Dependencies }}
	{{- if $j}} + {{ end }}cast.ToFloat64({{ findVarNameById $d.Id $ }})
	{{- end }}
}
	{{- else if eq .Type 1 }}  
func W{{ .CallName }}(
	{{- range $j, $d := .Dependencies }}
		{{- if $j}}, {{ end }}{{ findVarNameById $d.Id $ }}
	{{- end }} any) (result {{.ReturnType}}) { 
	task:= New{{ .CallName }}()
	task.init()
	err := ConvertStruct({{- range $j, $d := .Dependencies }}
		{{- if $j}}, {{ end }}{{ findVarNameById $d.Id $ }}
	{{- end }},&task.currentData{{- if (index .Dependencies 0).ReturnType }}.{{ findReNameById (index .Dependencies 0).Id $ }}{{ end }})
	if err != nil { 
		log.Fatal(err)
	}
	task.singleExecute()
	err = ConvertStruct(task.afterRet.{{.ReturnType}}, &result)
	if err != nil {
		log.Println("结果转换失败：", err)
		return {{.ReturnType}}{}
	}
	return result 	//返回值精确一点就ok(参数值也一样)
}
	{{- else if eq .Type 3 }}  
func {{ .CallName }}(
	{{- range $j, $d := .Dependencies }}
		{{- if $j}}, {{ end }}{{ findVarNameById $d.Id $ }}
	{{- end }} any) any { return nil }
	{{- end}}

{{- end }}
`

	funcStr := renderTemplateWithLookup(g.topResult, tmplStr)
	return funcStr
}
func (g *Generater) genFieldAssignments(node *TaskNode, parentPath string, parentVar string) string {
	var sb strings.Builder

	fieldPath := "out"
	if parentPath != "" {
		fieldPath = parentPath + "." + node.Id
	} else if node.Id != "" {
		fieldPath = node.Id
	}

	rightVar := parentVar

	// 优先匹配当前节点自己的依赖变量
	for _, value := range g.topResult {
		if value.Expr != "" && node.Id == value.Id {
			// 判断是否有子字段（在赋值前加注释）
			if len(node.Children) > 0 {
				sb.WriteString(fmt.Sprintf("\t//"))
			}
			// 如果是顶级字段
			if value.Type == Field {
				rightVar = value.Dependencies[0].VarName
				sb.WriteString(fmt.Sprintf("\t%s = %s\n", fieldPath, rightVar))
			} else if value.Type == SubField {
				// 子字段右边来自父变量
				if parentVar != "" {
					rightVar = fmt.Sprintf("%s.%s", parentVar, node.Id) // 父变量 + 当前字段
				} else {
					rightVar = fmt.Sprintf("%s.%s", value.Dependencies[0].VarName, node.Id) // 没父变量，用依赖起点
				}
				sb.WriteString(fmt.Sprintf("\t%s = %s\n", fieldPath, rightVar))
			}
			break
		}
	}
	// 递归子节点（传递当前 rightVar）
	for _, child := range node.Children {
		sb.WriteString(g.genFieldAssignments(child, fieldPath, rightVar))
	}

	return sb.String()
}

func (g *Generater) genExec() string {
	var b strings.Builder
	const topStr = `
	// 拓扑排序执行
	{{- $length := len . -}}
	{{- range $i, $n := . }}
		{{- if or (eq $n.Type 2) (eq $n.Type 3) (eq $n.Type 1) }}
		{{ $n.VarName }} := W{{ $n.CallName }}(
			{{- range $j, $d := $n.Dependencies }}
				{{- if $j}}, {{ end }}{{ findVarNameById $d.Id $ }}
			{{- end }})
		{{- else if eq $n.Type 5 }}
		{{ $n.VarName }} := {{ $n.CallName }}
		 {{- else if eq $n.Type 7 }}
		{{ $n.VarName }} := {{ $n.CallName }}
		{{- end }}
	{{- end }}
`

	//暴力遍历model，赋值out的每一个字段（有expr的通过id匹配map，取其varname）,递归遍历其子字段，子字段的赋值在其父字段之后

	feilds := g.genFieldAssignments(g.astNode, "", "")
	log.Println("进入字段赋值")
	const fieldtpStr2 = `
	//字段赋值
	var out Out
`
	fieldCode2 := renderTemplateWithLookup(g.topResult, fieldtpStr2)
	topCode := renderTemplateWithLookup(g.topResult, topStr)

	b.WriteString("func Run(In1 In1) Out {\n")
	b.WriteString("    // 执行依赖计算\n")

	b.WriteString(topCode)
	b.WriteString(fieldCode2 + feilds)

	b.WriteString("\n")
	b.WriteString(`
    return out
}
`)

	return b.String()
}

// 单部分获取
func (g *Generater) generateCodeSingle() (err error) {
	g.assignNames()
	g.importstr = g.genImport()
	g.structs = g.genStructs()
	g.funcs = g.genFuncs()
	g.exec = g.genExec()
	g.util = g.utilgen()
	g.main = g.maingen()
	return err
}
func (g *Generater) maingen() string {
	var b strings.Builder

	// 模板字符串，支持动态输入结构体名和运行函数
	const mainTpl = `
func main() {
	fmt.Println("workflow执行ing")
	inputFile := os.Args[1]
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("读取文件失败: %v", err)
	}
	config.WF.Model = "gemini-2.0-flash"
	config.WF.Config = "."
	var in In1
	
	json.Unmarshal(b,&in)
	fmt.Println("收到输入: ", in)
	log.Printf("收到输入: %+v\n", in)
	out := Run(in)
	fmt.Println("workflow执行44444444")
	// 输出 JSON 结果
	result, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		log.Fatalf("结果序列化失败: %v", err)
	}
	fmt.Println(string(result))
}
`

	code:= renderTemplateWithLookup(g.topResult, mainTpl)
	b.WriteString(code)
	return b.String()
}
func (g *Generater) utilgen() string {
	return `func ConvertStruct[A any, B any](src A, dst *B) error {
	data, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}
	if err := json.Unmarshal(data, dst); err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}
	return nil
}`
}

// 组装所有内容
func (g *Generater) combine() (err error) {
	g.fullCode = fmt.Sprintf(`// Code generated by WorkflowDSL Generator. DO NOT EDIT.
%s

%s

%s

%s

%s

%s
`, g.importstr, g.structs, g.funcs, g.exec, g.main, g.util)

	// 确保目录存在
	// os.MkdirAll("../generated_code", 0755)
	// path:="generated_code"
	// err = ioutil.WriteFile("generated_code/workflowSmp.go", []byte(g.fullCode), 0666)
	// if err != nil {
	// 	log.Print(err)
	// }
	// fmt.Println("==== STRUCTS ====\n", g.structs)
	// fmt.Println("==== FUNCS ====\n", g.funcs)
	// fmt.Println("==== EXEC ====\n", g.exec)

	return err
}

func (g *Generater) Execute() (err error) {
	err = g.generateCodeSingle()
	if err != nil {
		return err
	}
	err = g.combine()
	if err != nil {
		return err
	}

	return err
}
