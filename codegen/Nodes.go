package codegen

import (
	"fmt"
	"strings"
	// "codegen/util"
)

type FieldDef struct {
	Name        string
	Type        string
	JsonName    string
	Hint        string
	Annotations []string
	SubFields   []FieldDef
}

type PromptEvalContext struct {
	Vars       map[string]interface{}
	InFields   []FieldDef
	OutFields  []FieldDef
	ModelFields  []FieldDef
	Input      any
	ModuleDefs map[string][]Node
}

type Node interface {
	Eval(_ *PromptEvalContext) ([]string, error)
	// ConvertToCode() string
}

// 引用结点

type ExprOp int

const (
	ExprOp_And ExprOp = iota
	ExprOp_Or
	ExprOp_Not
	ExprOp_Equal
	ExprOp_NotEqual
	ExprOp_GreaterThan
	ExprOp_LessThan
	ExprOp_GreaterThanOrEqual
	ExprOp_LessThanOrEqual
	ExprOp_Add
	ExprOp_Sub
	ExprOp_Mul
	ExprOp_Div
	ExprOp_Mod

	ExprOp_None // for leaf nodes
)

type Expr struct {
	Leaf *string

	Op       ExprOp
	Operant0 []*Expr
	Operant1 []*Expr
}

func (e *Expr) EvalToBool(ctx *PromptEvalContext) (bool, error) {
	// 如果是叶子节点，不能直接判断真假，必须是比较表达式
	if e.Leaf != nil {
		return false, fmt.Errorf("cannot evaluate leaf as boolean directly: %s", *e.Leaf)
	}

	// 支持的布尔比较运算
	switch e.Op {
	case ExprOp_Equal, ExprOp_NotEqual:
		leftExpr := e.Operant0[0]
		rightExpr := e.Operant1[0]

		leftVals, err := leftExpr.Eval(ctx)
		if err != nil {
			return false, err
		}
		rightVals, err := rightExpr.Eval(ctx)
		if err != nil {
			return false, err
		}

		left := firstOrEmpty(leftVals)
		right := firstOrEmpty(rightVals)
		fmt.Println("😶left:", left)
		fmt.Println("😶right:", right)
		// fmt.Println("🔍 Comparing:", left, e.Op, right)

		switch e.Op {
		case ExprOp_Equal:
			return left == right, nil
		case ExprOp_NotEqual:
			return left != right, nil
		}
	}

	return false, fmt.Errorf("unsupported boolean operator: %v", e.Op)
}
func (e *Expr) Eval(ctx *PromptEvalContext) ([]string, error) {
	// 如果是叶子节点

	if e.Leaf != nil {
		leafVal := *e.Leaf
		// 假设叶子是参数路径格式，比如 "input.question" 或 "output.answer"
		parts := strings.SplitN(leafVal, ".", 2)

		if len(parts) == 2 {
			section, key := parts[0], parts[1]
			var val any
			var ok bool

			switch section {
			case "input":
				// 假设 ctx.Input 是 map[string]any
				if m, okCast := ctx.Input.(map[string]any); okCast {
					val, ok = m[key]
				}
			case "outputspec":
				// 你也可以在 ctx 中保存输出数据结构，这里示例用 OutFields 不一样，按实际需求调整
				// 这里仅示例直接返回 key
				val, ok = key, true
			default:
				return nil, fmt.Errorf("unknown section in leaf: %s", section)
			}

			if !ok {
				return []string{""}, nil // 找不到值返回空字符串，或者你也可以返回错误
			}
			return []string{fmt.Sprintf("%v", val)}, nil
		}

		// 如果不是带点的路径，就直接返回叶子内容
		return []string{leafVal}, nil
	}

	// 非叶子节点暂时不支持Eval
	return nil, fmt.Errorf("Eval not supported on non-leaf expressions yet")
}
func firstOrEmpty(list []string) string {
	if len(list) == 0 {
		return ""
	}
	return list[0]
}
func (e *Expr) EvalToInt(_ *PromptEvalContext) (int, error) {
	return 0, nil
}

func (e *Expr) EvalToFloat(_ *PromptEvalContext) (float32, error) {
	return 0.0, nil
}

func (e *Expr) EvalToString(_ *PromptEvalContext) (string, error) {
	return "", nil
}

type StringNode struct {
	Val string
}

func (node *StringNode) Eval(_ *PromptEvalContext) ([]string, error) {
	escaped := strings.ReplaceAll(node.Val, `"`, `\"`)
	return []string{fmt.Sprintf(`b.WriteString("%s\n")`, escaped)}, nil
	// return []string{node.Val}, nil
}

type OutputSpecNode struct {
	IsArray bool
	RawTyp  string
}

func (node *OutputSpecNode) Eval(ctx *PromptEvalContext) ([]string, error) {
	// TODO
	var b strings.Builder
	fields:=ctx.OutFields
	if ctx.ModelFields!=nil{
		fields=ctx.ModelFields
	}
	if node.IsArray {
		for _, line := range BuildModelOutputSpecLines(fields, true) {
			b.WriteString(fmt.Sprintf("    b.WriteString(\"%s\\n\")\n", strings.ReplaceAll(line, "\"", "\\\"")))
		}
		return []string{b.String()}, nil
	}
	for _, line := range BuildModelOutputSpecLines(fields, false) {
		b.WriteString(fmt.Sprintf("    b.WriteString(\"%s\\n\")\n", strings.ReplaceAll(line, "\"", "\\\"")))
	}
	return []string{b.String()}, nil

	// return []string{}, nil
}

type ModuleRefNode struct {
	Name string
}

func (m *ModuleRefNode) Eval(ctx *PromptEvalContext) ([]string, error) {

	nodes := ctx.ModuleDefs[m.Name]
	if nodes == nil {
		return []string{fmt.Sprintf("[Missing module: %s]", m.Name)}, nil
		// return nil,nil
	}
	var result []string
	for _, node := range nodes {
		str, err := node.Eval(ctx)
		if err != nil {
			return nil, err
		}
		result = append(result, str...)
	}
	return result, nil
}

type InputNode struct {
	Fields []FieldDef
}

func (n *InputNode) Eval(ctx *PromptEvalContext) ([]string, error) {
	// Eval逻辑根据需要实现
	return nil, nil
}

type OutputNode struct {
	Fields []FieldDef
}

func (n *OutputNode) Eval(ctx *PromptEvalContext) ([]string, error) {
	// Eval逻辑根据需要实现
	return nil, nil
}

type MarkdownNode struct {
	Content string
}

func (m *MarkdownNode) Eval(ctx *PromptEvalContext) ([]string, error) {
	return []string{m.Content}, nil
}

type IfNode struct {
	Condition string
	Then      []Node
	Else      []Node
}

func (e *Expr) String() string {
	if e.Leaf != nil {
		return *e.Leaf
	}
	if e.Op != ExprOp_None && len(e.Operant0) > 0 && len(e.Operant1) > 0 {
		return fmt.Sprintf("%s %s %s", e.Operant0[0].String(), e.Op.String(), e.Operant1[0].String())
	}
	return ""
}
func (op ExprOp) String() string {
	switch op {
	case ExprOp_Equal:
		return "=="
	case ExprOp_NotEqual:
		return "!="
	default:
		return ""
	}
}
func (node *IfNode) Eval(ctx *PromptEvalContext) ([]string, error) {

	var lines []string

	cond := node.Condition // Expr -> string，比如 input.question != ""

	// 开始 if 结构
	lines = append(lines, fmt.Sprintf("if (%s) {", cond))

	// Then 分支
	for _, n := range node.Then {
		vals, err := n.Eval(ctx)
		if err != nil {
			return nil, fmt.Errorf("Then branch Eval failed: %w", err)
		}
		for _, v := range vals {
			lines = append(lines, fmt.Sprintf("    %s", v))
		}
	}

	// else 分支（如果有）
	if len(node.Else) > 0 {
		lines = append(lines, "} else {")
		for _, n := range node.Else {
			vals, err := n.Eval(ctx)
			if err != nil {
				return nil, fmt.Errorf("Else branch Eval failed: %w", err)
			}
			for _, v := range vals {
				lines = append(lines, fmt.Sprintf("    %s", v))
			}
		}
	}

	// 结束
	lines = append(lines, "}")

	return lines, nil
}

type ParamNode struct {
	Path string
}

// 实现 Node 接口的 Eval 方法
func (p *ParamNode) Eval(ctx *PromptEvalContext) ([]string, error) {

	// var section, key string
	// parts := strings.SplitN(p.Path, ".", 2)
	// if len(parts) == 2 {
	// 	section, key = parts[0], parts[1]
	// } else if len(parts) == 1 {
	// 	section = parts[0]
	// 	key = ""
	// } else {
	// 	return nil, fmt.Errorf("invalid param path: %s", p.Path)
	// }
	// // section, key := parts[0], parts[1]

	// switch section {
	// case "input":
	// 	// var input string
	// 	// 从 ctx.Input 中获取真实值
	// 	// fmt.Println("解释执行 ParamNode:", ctx.Input)
	// 	inputMap, ok := ctx.Input.(map[string]any)
	// 	if !ok {
	// 		return nil, fmt.Errorf("input is not a valid map[string]any")
	// 	}

	// 	if val, ok := inputMap[key]; ok {
	// 		return []string{fmt.Sprintf("%v", val)}, nil
	// 	}
	// case "output":
	// default:
	// 	return nil, fmt.Errorf("unknown param section: %s", section)
	// }

	return []string{fmt.Sprintf(`b.WriteString(%s)`, p.Path)}, nil
}

type CasePair struct {
	Case string
	Body []Node
}

type SwitchNode struct {
	Switch  string
	Cases   []CasePair
	Default []Node
}

func (node *SwitchNode) Eval(ctx *PromptEvalContext) ([]string, error) {
	var lines []string

	cond := node.Switch // Expr -> string，比如 input.question != ""

	// 开始 switch 结构
	lines = append(lines, fmt.Sprintf("switch %s {", cond))
	for _, caseitem := range node.Cases {
		lines = append(lines, fmt.Sprintf("case %s:", caseitem.Case))
		for _, n := range caseitem.Body {
			vals, err := n.Eval(ctx)
			if err != nil {
				return nil, fmt.Errorf("Case branch Eval failed: %w", err)
			}
			for _, v := range vals {
				lines = append(lines, fmt.Sprintf("    %s", v))
			}
		}
	}
	if node.Default != nil {
		lines = append(lines, "default:")
		for _, n := range node.Default {
			vals, err := n.Eval(ctx)
			if err != nil {
				return nil, fmt.Errorf("Default branch Eval failed: %w", err)
			}
			for _, v := range vals {
				lines = append(lines, fmt.Sprintf("    %s", v))
			}
		}
	}
	// 结束
	lines = append(lines, "}")
	return lines, nil
}

type ForNode struct {
	ForType string // "traditional" / "rangeWithIndex" / "rangeNoIndex"
	Init    string // traditional
	Cond    string // traditional
	Post    string // traditional
	Key     string // rangeWithIndex
	Val     string // rangeNoIndex or rangeWithIndex
	Range   string // iterable expr
	Body    []Node // 通用内容
}

func (node *ForNode) Eval(ctx *PromptEvalContext) ([]string, error) {
	// fmt.Println("解释执行 ForNode:", node.Key, node.Val, node.Range)
	var lines []string

	switch node.ForType {
	case "traditional":
		lines = append(lines, fmt.Sprintf("for %s; %s; %s {", node.Init, node.Cond, node.Post))

		for _, n := range node.Body {
			vals, err := n.Eval(ctx)
			if err != nil {
				return nil, fmt.Errorf("Then branch Eval failed: %w", err)
			}
			for _, v := range vals {
				lines = append(lines, fmt.Sprintf("    %s", v))
			}
		}

		lines = append(lines, "}")

	case "rangeWithIndex":
		// for key, val := range iterable
		fmt.Println("解释执行 ForNode:", node.Key, node.Val, node.Range)
		lines = append(lines, fmt.Sprintf("for %s, %s := range %s {", node.Key, node.Val, node.Range))

		for _, n := range node.Body {
			vals, err := n.Eval(ctx)
			if err != nil {
				return nil, fmt.Errorf("RangeWithIndex branch Eval failed: %w", err)
			}
			for _, v := range vals {
				lines = append(lines, fmt.Sprintf("    %s", v))
			}
		}

		lines = append(lines, "}")

	case "rangeNoIndex":
		fmt.Println("解释执行 ForNode:", node.Key, node.Val, node.Range)
		// for val := range iterable
		lines = append(lines, fmt.Sprintf("for %s := range %s {", node.Val, node.Range))

		for _, n := range node.Body {
			vals, err := n.Eval(ctx)
			if err != nil {
				return nil, fmt.Errorf("RangeNoIndex branch Eval failed: %w", err)
			}
			for _, v := range vals {
				lines = append(lines, fmt.Sprintf("    %s", v))
			}
		}

		lines = append(lines, "}")

	default:
		return nil, fmt.Errorf("unknown for loop type: %s", node.ForType)
	}

	return lines, nil
}

type PromptNode struct {
	Vars             map[string]interface{}
	SysNodes         []Node
	UserNodes        []Node
	BeforeNodes      []Node
	ModuleDefs       map[string][]Node // 初始化 map
	InFields         []FieldDef
	OutFields        []FieldDef
	ModelFields      []FieldDef
	BeforeCode       string
	FixCode          []string
	AfterCode        []string
	Goimport         []goimport
	outputspectNodes OutputSpecNode
	// IsArray     bool
	// 其它部分
}
type goimport struct {
	Alias string
	Path  string
}
type final struct {
	User []string
	Sys  []string
}

func (r *PromptNode) Eval(ctx *PromptEvalContext) (*final, error) {

	var user []string
	var sys []string
	for _, node := range r.SysNodes {
		out, err := node.Eval(ctx)
		if err != nil {
			return nil, err
		}
		sys = append(sys, out...)
	}
	for _, node := range r.UserNodes {
		out, err := node.Eval(ctx)
		if err != nil {
			return nil, err
		}
		user = append(user, out...)
	}

	return &final{
		User: user,
		Sys:  sys,
	}, nil
}

type ParsedPrompt struct {
	InDef  []FieldDef
	OutDef []FieldDef

	// Sys []Node
	// User []Node
	// Or:
	SysVars  map[string][]Node
	UserVars map[string][]Node

	BeforeCode string
	FixCode    string
	AfterCode  string
}

func (parsed *ParsedPrompt) GetSystemPrompt(inp any) (string, error) {
	return "", nil
}

func (parsed *ParsedPrompt) GetUserPrompt(inp any) (string, error) {
	return "", nil
}

// Or
func (parsed *ParsedPrompt) GenSystemPromptFn() (string, error) {
	return "", nil
}

func (parsed *ParsedPrompt) GenUserPromptFn() (string, error) {
	return "", nil
}
