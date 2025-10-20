package wcodegen

import (
	"log"
	"strings"

	Wparser "github.com/ycstella/prompt-dsl/codegen/Wparser"
)

// TaskNode 类型
type NodeType int

const (
	Direct NodeType = iota
	Func
	Arithmetic
	Conditional
	Struct
	Atom
	Field
	FAtom
	SubField
)

type StructMode int

const (
	StructNone StructMode = iota
	StructFieldBased
	StructWholeBased
)

type TaskNode struct {
	Id    string
	Type  NodeType
	Value any
	Expr  string

	Func         func() any
	Dependencies []*TaskNode
	Children     []*TaskNode

	VarName    string
	CallName   string
	ReturnType string

	StructMode StructMode
}
type TaskNodeJSON struct {
	Id           string          `json:"id"`
	Type         NodeType        `json:"type"`
	Value        any             `json:"value"`
	Expr         string          `json:"expr"`
	Dependencies []*TaskNodeJSON `json:"dependencies"`
	Children     []*TaskNodeJSON `json:"children"`
	VarName      string
	CallName     string
}

type WAstbuilder struct {
	TaskList []*TaskNode
	ctx      *Wparser.WorkflowContext
}

func NewWAstbuilder(ctx *Wparser.WorkflowContext) WAstbuilder {
	return WAstbuilder{
		ctx: ctx,
	}
}

func (w *WAstbuilder) ParseWorkflowBodyToTaskNode() (err error) {
	var nodes []*TaskNode

	// 解析 out 部分
	outCtx := w.ctx.WorkflowBody().Out()
	if outCtx != nil {
		outNode := &TaskNode{
			Id:       "out",
			Type:     Struct,
			Children: []*TaskNode{},
		}

		for _, field := range outCtx.AllOutField() {
			child := w.parseOutField(field)
			outNode.Children = append(outNode.Children, child)
		}
		outNode.Dependencies = outNode.Children
		nodes = append(nodes, outNode)
		w.TaskList = nodes
	}
	return err
}

// 递归解析 outField
func (w *WAstbuilder) parseOutField(ctx Wparser.IOutFieldContext) *TaskNode {
	node := &TaskNode{
		Id:         ctx.ID().GetText(),
		Type:       Field,
		ReturnType: ctx.Outtype().GetText(),
	}

	rt := strings.TrimSpace(node.ReturnType)
	if strings.HasPrefix(rt, "struct{") && strings.HasSuffix(rt, "}") {
		// 如果是结构体定义，则返回类型设为字段 Id 首字母大写
		node.ReturnType = strings.Title(node.Id)
	}

	exprCtx := ctx.Expr()
	if exprCtx != nil {
		exprNode := w.parseExpr(exprCtx, node.ReturnType)
		node.Dependencies = append(node.Dependencies, exprNode)
		node.Expr = exprCtx.GetText()
	}

	//如果是 struct 类型
	structCtx := ctx.Outtype().STRUCT()
	if structCtx != nil {
		if exprCtx != nil {
			node.StructMode = StructWholeBased
		} else {
			node.StructMode = StructFieldBased
		}
		node.Children = []*TaskNode{}
		for _, subField := range ctx.Outtype().AllOutField() {
			child := w.parseOutField(subField)
			node.Children = append(node.Children, child)
		}
	}

	return node
}

// 解析任意 expr → TaskNode
func (w *WAstbuilder) parseExpr(ctx Wparser.IExprContext, returnType string) *TaskNode {
	switch {
	case ctx.FunctionCall() != nil:
		fnNode := &TaskNode{
			Type:       Func,
			Id:         ctx.GetText(),
			Func:       func() any { return nil },
			Expr:       ctx.GetText(),
			ReturnType: returnType,
		}
		// 递归解析参数
		for _, arg := range ctx.FunctionCall().AllExpr() {
			argNode := w.parseExpr(arg, returnType)

			fnNode.Dependencies = append(fnNode.Dependencies, argNode)
		}
		return fnNode

	case ctx.ArithmeticExpr() != nil:
		arithNode := &TaskNode{
			Type:       Arithmetic,
			Id:         ctx.GetText(),
			Expr:       ctx.GetText(),
			ReturnType: returnType,
		}
		// log.Println("ArithmeticExpr生效：", arithNode.Id)
		// 判断是否是 $sum([...])
		if ds, ok := ctx.ArithmeticExpr().(*Wparser.DollarSumExprContext); ok {
			top := ds.ArithExprContent()
			w.collectArithDependencies(top, &arithNode.Dependencies, returnType)
		}
		return arithNode

	case ctx.ConditionalExpr() != nil:
		condNode := &TaskNode{
			Type: Conditional,
			Expr: ctx.ConditionalExpr().Condition().GetText(),
		}
		condNode.Children = []*TaskNode{
			w.parseExpr(ctx.ConditionalExpr().Expr(0), returnType),
			w.parseExpr(ctx.ConditionalExpr().Expr(1), returnType),
		}
		return condNode

	case ctx.Atom() != nil:
		return w.parseAtom(ctx.Atom(), returnType)
	default:
		return &TaskNode{
			Type: Direct,
			Id:   ctx.GetText(),
		}
	}
}

func (w *WAstbuilder) parseAtom(atomCtx Wparser.IAtomContext, returnType string) *TaskNode {
	node := &TaskNode{
		Id:   atomCtx.GetText(),
		Expr: atomCtx.GetText(),
		Type: Atom,
	}
	if atomCtx.FunctionCall() != nil {
		funcText := atomCtx.FunctionCall().GetText()
		fnNode := &TaskNode{
			Id:   funcText,
			Expr: funcText,
			Type: Func,
		}
		node.Type = FAtom
		for _, arg := range atomCtx.FunctionCall().AllExpr() {
			argNode := w.parseExpr(arg, returnType)
			fnNode.Dependencies = append(fnNode.Dependencies, argNode)
		}
		node.Dependencies = append(node.Dependencies, fnNode)
	}
	return node
}

func (w *WAstbuilder) collectArithDependencies(ctx Wparser.IArithExprContentContext, deps *[]*TaskNode, returnType string) {
	parenList := ctx.AllParenExpr()
	for _, p := range parenList {
		if expr := p.Expr(); expr != nil {
			node := w.parseExpr(expr, returnType)
			*deps = append(*deps, node)
			if expr.ArithmeticExpr() != nil {
				if ds, ok := expr.ArithmeticExpr().(*Wparser.DollarSumExprContext); ok {
					w.collectArithDependencies(ds.ArithExprContent(), deps, returnType)
				}
			}
		}
	}
}
func (w *WAstbuilder)PreprocessDeps()(err error) {
	root:=w.TaskList[0]
	var dfs func(node, parent *TaskNode)
	dfs = func(node, parent *TaskNode) {
		if node.Expr == "" && parent != nil && len(parent.Dependencies) > 0 {
			node.Dependencies = append(node.Dependencies, parent.Dependencies...)
			node.Expr = parent.Expr
			node.Type = SubField
			log.Printf("依赖继承: %s <- %s\n", node.Id, parent.Id)
		}
		for _, child := range node.Children {
			dfs(child, node)
		}
	}
	dfs(root, nil)
	return err
}

func (w *WAstbuilder) astbuild() (err error) {
	err = w.ParseWorkflowBodyToTaskNode()
	if err != nil {
		return err
	}
	err = w.PreprocessDeps()
	if err != nil {
		return err
	}
	return nil
}
