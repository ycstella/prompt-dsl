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


type PromptGenContext struct {
	Vars        map[string]interface{}
	InFields    []FieldDef
	OutFields   []FieldDef
	ModelFields []FieldDef
	Input       any
	ModuleDefs  map[string][]Node
}

type Node interface {
	Tocode(_ *PromptGenContext) ([]string, error)
	// ConvertToCode() string
}


// func firstOrEmpty(list []string) string {
// 	if len(list) == 0 {
// 		return ""
// 	}
// 	return list[0]
// }

type StringNode struct {
	Val string
}

func (node *StringNode) Tocode(_ *PromptGenContext) ([]string, error) {
	escaped := strings.ReplaceAll(node.Val, `"`, `\"`)
	return []string{fmt.Sprintf(`b.WriteString("%s\n")`, escaped)}, nil
	// return []string{node.Val}, nil
}

type OutputSpecNode struct {
	IsArray bool
	RawTyp  string
}

func (node *OutputSpecNode) Tocode(ctx *PromptGenContext) ([]string, error) {
	// TODO
	var b strings.Builder
	fields := ctx.OutFields
	if ctx.ModelFields != nil {
		fields = ctx.ModelFields
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

func (m *ModuleRefNode) Tocode(ctx *PromptGenContext) ([]string, error) {

	nodes := ctx.ModuleDefs[m.Name]
	if nodes == nil {
		return []string{fmt.Sprintf("[Missing module: %s]", m.Name)}, nil
		// return nil,nil
	}
	var result []string
	for _, node := range nodes {
		str, err := node.Tocode(ctx)
		if err != nil {
			return nil, err
		}
		result = append(result, str...)
	}
	return result, nil
}

// type InputNode struct {
// 	Fields []FieldDef
// }

// func (n *InputNode) Tocode(ctx *PromptGenContext) ([]string, error) {
// 	return nil, nil
// }

// type OutputNode struct {
// 	Fields []FieldDef
// }

// func (n *OutputNode) Tocode(ctx *PromptGenContext) ([]string, error) {
// 	// Tocode逻辑根据需要实现
// 	return nil, nil
// }

type MarkdownNode struct {
	Content string
}

func (m *MarkdownNode) Tocode(ctx *PromptGenContext) ([]string, error) {
	return []string{m.Content}, nil
}

type IfNode struct {
	Condition string
	Then      []Node
	Else      []Node
}


func (node *IfNode) Tocode(ctx *PromptGenContext) ([]string, error) {

	var lines []string

	cond := node.Condition // Expr -> string，比如 input.question != ""

	// 开始 if 结构
	lines = append(lines, fmt.Sprintf("if (%s) {", cond))

	// Then 分支
	for _, n := range node.Then {
		vals, err := n.Tocode(ctx)
		if err != nil {
			return nil, fmt.Errorf("Then branch Tocode failed: %w", err)
		}
		for _, v := range vals {
			lines = append(lines, fmt.Sprintf("    %s", v))
		}
	}

	// else 分支（如果有）
	if len(node.Else) > 0 {
		lines = append(lines, "} else {")
		for _, n := range node.Else {
			vals, err := n.Tocode(ctx)
			if err != nil {
				return nil, fmt.Errorf("Else branch Tocode failed: %w", err)
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

// 实现 Node 接口的 Tocode 方法
func (p *ParamNode) Tocode(ctx *PromptGenContext) ([]string, error) {
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

func (node *SwitchNode) Tocode(ctx *PromptGenContext) ([]string, error) {
	var lines []string

	cond := node.Switch // Expr -> string，比如 input.question != ""

	// 开始 switch 结构
	lines = append(lines, fmt.Sprintf("switch %s {", cond))
	for _, caseitem := range node.Cases {
		lines = append(lines, fmt.Sprintf("case %s:", caseitem.Case))
		for _, n := range caseitem.Body {
			vals, err := n.Tocode(ctx)
			if err != nil {
				return nil, fmt.Errorf("Case branch Tocode failed: %w", err)
			}
			for _, v := range vals {
				lines = append(lines, fmt.Sprintf("    %s", v))
			}
		}
	}
	if node.Default != nil {
		lines = append(lines, "default:")
		for _, n := range node.Default {
			vals, err := n.Tocode(ctx)
			if err != nil {
				return nil, fmt.Errorf("Default branch Tocode failed: %w", err)
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

func (node *ForNode) Tocode(ctx *PromptGenContext) ([]string, error) {
	// fmt.Println("解释执行 ForNode:", node.Key, node.Val, node.Range)
	var lines []string

	switch node.ForType {
	case "traditional":
		lines = append(lines, fmt.Sprintf("for %s; %s; %s {", node.Init, node.Cond, node.Post))

		for _, n := range node.Body {
			vals, err := n.Tocode(ctx)
			if err != nil {
				return nil, fmt.Errorf("Then branch Tocode failed: %w", err)
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
			vals, err := n.Tocode(ctx)
			if err != nil {
				return nil, fmt.Errorf("RangeWithIndex branch Tocode failed: %w", err)
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
			vals, err := n.Tocode(ctx)
			if err != nil {
				return nil, fmt.Errorf("RangeNoIndex branch Tocode failed: %w", err)
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

func (r *PromptNode) Tocode(ctx *PromptGenContext) (*final, error) {

	var user []string
	var sys []string
	for _, node := range r.SysNodes {
		out, err := node.Tocode(ctx)
		if err != nil {
			return nil, err
		}
		sys = append(sys, out...)
	}
	for _, node := range r.UserNodes {
		out, err := node.Tocode(ctx)
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

// type ParsedPrompt struct {
// 	InDef  []FieldDef
// 	OutDef []FieldDef

// 	// Sys []Node
// 	// User []Node
// 	// Or:
// 	SysVars  map[string][]Node
// 	UserVars map[string][]Node

// 	BeforeCode string
// 	FixCode    string
// 	AfterCode  string
// }

// func (parsed *ParsedPrompt) GetSystemPrompt(inp any) (string, error) {
// 	return "", nil
// }

// func (parsed *ParsedPrompt) GetUserPrompt(inp any) (string, error) {
// 	return "", nil
// }

// // Or
// func (parsed *ParsedPrompt) GenSystemPromptFn() (string, error) {
// 	return "", nil
// }

// func (parsed *ParsedPrompt) GenUserPromptFn() (string, error) {
// 	return "", nil
// }
