package wcodegen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"github.com/antlr4-go/antlr/v4"
	Wparser "github.com/ycstella/prompt-dsl/codegen/Wparser"
)

type Workflow_gen struct {
	Tree Wparser.IWorkflowContext
}

func NewWorkflow_gen() *Workflow_gen {
	return &Workflow_gen{}
}
func (w *Workflow_gen) workflowParser() error {
	b, err := ioutil.ReadFile("workflowSmp.txt")
	if err != nil {
		log.Print(err)
	}
	inputStream := antlr.NewInputStream(string(b))
	lexer := Wparser.NewWorkflowDSLLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := Wparser.NewWorkflowDSLParser(stream)
	w.Tree = p.Workflow()
	log.Println("🌳 start...")
	log.Println(w.Tree.ToStringTree(nil, p))
	log.Println("🌳 ...end")
	return err
}
func main() {
	w := NewWorkflow_gen()
	err := w.workflowParser()
	log.Println(err)
	// // 构建依赖关系
	// in1 := &Node{Name: "in1"}
	// extract := &Node{Name: "ExtractSteps(in1)", Deps: []*Node{in1}}
	// p2 := &Node{Name: "p2(in1)", Deps: []*Node{in1}}
	// p2Qty := &Node{Name: "p2(in1).quantity", Deps: []*Node{p2}}
	// getPath := &Node{Name: "getPath(...)", Deps: []*Node{extract, p2Qty}}

	// // 执行拓扑排序逻辑（递归）
	// getPath.Exec()

}

// ---------------- 数据模型 ----------------
type Workflow struct {
	Inputs  []StructDef
	Outputs []StructDef
}

type StructDef struct {
	Name    string
	Fields  []FieldDef
	IsArray bool
}

type FieldDef struct {
	Name string
	Type string
	Expr string
	Deps []string
}
type ExprDepExtractor struct {
	deps map[string]struct{}
}

func NewExprDepExtractor() *ExprDepExtractor {
	return &ExprDepExtractor{deps: make(map[string]struct{})}
}

// 收集一个依赖项
func (e *ExprDepExtractor) Add(dep string) {
	e.deps[dep] = struct{}{}
}

func (e *ExprDepExtractor) List() []string {
	res := make([]string, 0, len(e.deps))
	for k := range e.deps {
		res = append(res, k)
	}
	return res
}
func (e *ExprDepExtractor) ExtractExpr(exprCtx Wparser.IExprContext) {
	switch {
	case exprCtx.FunctionCall() != nil:
		fc := exprCtx.FunctionCall()
		e.Add(fc.ID().GetText()) // 函数本身也算依赖
		for _, arg := range fc.AllExpr() {
			e.ExtractExpr(arg) // 递归提取参数依赖
		}

	case exprCtx.ConditionalExpr() != nil:
		ce := exprCtx.ConditionalExpr()
		e.ExtractCondition(ce.Condition()) // 条件部分
		e.ExtractExpr(ce.Expr(0))          // if 分支
		e.ExtractExpr(ce.Expr(1))          // else 分支

	case exprCtx.ArithmeticExpr() != nil:
		ae := exprCtx.ArithmeticExpr()

		switch a := ae.(type) {
		case *Wparser.DollarSumExprContext:
			// $sum(...)
			if content := a.ArithExprContent(); content != nil {
				e.ExtractArithExpr(content)
			}

		default:
			// 普通算术表达式
			for _, child := range ae.GetChildren() {
				if c, ok := child.(Wparser.IArithExprContentContext); ok {
					e.ExtractArithExpr(c)
				}
			}
		}

	case exprCtx.Atom() != nil:
		e.ExtractAtom(exprCtx.Atom())

	case exprCtx.INT() != nil:
		// 整数无需依赖
	}
}
func (e *ExprDepExtractor) ExtractCondition(condCtx Wparser.IConditionContext) {
	if condCtx.AllExpr() != nil {
		for _, value := range condCtx.AllExpr() {
			e.ExtractExpr(value)
		}

	} else if len(condCtx.AllCondition()) > 0 {
		for _, c := range condCtx.AllCondition() {
			e.ExtractCondition(c)
		}
	}
}
func (e *ExprDepExtractor) ExtractArithExpr(aeCtx Wparser.IArithExprContentContext) {

}

func (e *ExprDepExtractor) ExtractAtom(atomCtx Wparser.IAtomContext) {
	if atomCtx.FunctionCall() != nil {
		for _, value := range atomCtx.FunctionCall().AllExpr() {
			e.ExtractExpr(value)
		}
	} else {
		e.Add(atomCtx.GetText()) // 原子标识符直接作为依赖
	}
}

// parseOutField 解析字段
func parseField(of Wparser.IOutFieldContext, parentName string, structs *[]StructDef) FieldDef {
	fd := FieldDef{Name: of.ID().GetText()}

	if of.Expr() != nil {
		fd.Expr = of.Expr().GetText()
		extractor := NewExprDepExtractor()
		extractor.ExtractExpr(of.Expr())
		fd.Deps = extractor.List()
	}

	// 处理类型
	fd.Type = parsetype(of.Outtype(), fd.Name, parentName, structs)
	return fd
}

// parseOuttype 解析类型，如果是 struct 则生成新的 StructDef
func parsetype(ot Wparser.IOuttypeContext, fieldName, parentName string, structs *[]StructDef) string {
	if ot.STRING_TYPE() != nil {
		return "string"
	}
	if ot.INT_TYPE() != nil {
		return "int"
	}
	if ot.FLOAT_TYPE() != nil {
		return "float64"
	}
	if ot.ID() != nil {
		return ot.ID().GetText()
	}
	if ot.FLOAT_TYPE() != nil {
		return "float64"
	}

	if ot.STRUCT() != nil {
		// struct 名称规则: 字段名（首字母大写）
		structName := strings.Title(fieldName)

		var children []FieldDef
		for _, inner := range ot.AllOutField() {
			child := parseField(inner, structName, structs)
			children = append(children, child)
		}
		*structs = append(*structs, StructDef{
			Name:   structName,
			Fields: children,
		})
		return structName
	}
	if ot.LBRACK() != nil {
		elemType := parsetype(ot.Outtype(), fieldName, parentName, structs)
		return "[]" + elemType
	}
	return "interface{}"
	//如果内部有定义则取内部定义（声明内部定义结构体，改写外部结构体）
}

// ---------------- AST → 数据模型 ----------------
func BuildModel(ctx *Wparser.WorkflowContext) *Workflow {
	model := &Workflow{}

	// 输入
	if inCtx := ctx.WorkflowBody().Inputs(); inCtx != nil {
		for _, input := range inCtx.AllInputBlock() {
			fmt.Println(input)
			sd := StructDef{}
			if inId := input.IN_ID(); inId != nil {
				sd.Name = inId.GetText()
			} else if inIdArr := input.IN_ID_ARRAY(); inIdArr != nil {
				sd.Name = strings.Trim(inIdArr.GetText(), "[]")
				sd.IsArray = true // 可以加一个标记表示这是数组
			} else {
				log.Println("⚠️ unexpected inputBlock without IN_ID or IN_ID_ARRAY")
			}

			for _, field := range input.AllFieldDef() {
				sd.Fields = append(sd.Fields, FieldDef{
					Name: field.ID().GetText(),
					Type: field.Type_().GetText(),
				})
			}
			model.Inputs = append(model.Inputs, sd)
		}
	}

	// 输出
	var structs []StructDef
	if outCtx := ctx.WorkflowBody().Out(); outCtx != nil {
		out := StructDef{Name: "Out"}
		for _, of := range outCtx.AllOutField() {
			field := parseField(of, "Out", &structs)
			out.Fields = append(out.Fields, field)
		}
		structs = append(structs, out)
	}
	model.Outputs = structs
	return model
}

// ---------------- 模板 ----------------
const tmpl = `

{{range .Inputs}}
type {{.Name}} struct {
    {{- range .Fields}}
    {{.Name}} {{.Type}}
    {{- end}}
}
{{end}}

{{range .Outputs}}
type {{.Name}} struct {
    {{- range .Fields}}
    {{.Name}} {{.Type}}
    {{- end}}
}
{{end}}
`

func GenerateCode(model *Workflow) string {
	t := template.Must(template.New("code").Parse(tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, model); err != nil {
		panic(err)
	}
	return buf.String()
}

