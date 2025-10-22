package wcodegen

import (
	"github.com/antlr4-go/antlr/v4"
	Wparser "github.com/ycstella/prompt-dsl/codegen/Wparser"
)

type ArrowStructListener struct {
	*Wparser.BaseWorkflowDSLParserListener
	funcName string
	content  string
	result   string
	Matches  [][]string
}

func NewArrowStructListener(funcName, content string) *ArrowStructListener {
	return &ArrowStructListener{
		funcName: funcName,
		content:  content,
		Matches:  make([][]string, 0),
	}
}
func (l *ArrowStructListener) EnterFunctionCall(ctx *Wparser.FunctionCallContext) {
	id := ctx.ID().GetText()
	if id != l.funcName {
		return
	}

	// 找到 funcName 调用
	start := ctx.GetStart().GetStart()
	stop := ctx.GetStop().GetStop()

	// 沿父节点向上找到最外层 struct 或 outField
	parent := ctx.GetParent()
	for parent != nil {
		switch n := parent.(type) {
		case *Wparser.OutFieldContext:
			start = n.GetStart().GetStart()
			stop = n.GetStop().GetStop()
			goto FOUND
		case *Wparser.TypeContext:
			start = n.GetStart().GetStart()
			stop = n.GetStop().GetStop()
			goto FOUND
		}
		parent = parent.(antlr.ParserRuleContext).GetParent()
	}
FOUND:
	content := l.content

	// ✅ 向前扩展 start，使其包含行首缩进
	for start > 0 {
		ch := content[start-1]
		if ch == ' ' || ch == '\t' {
			start--
			continue
		}
		if ch == '\n' || ch == '\r' {
			break
		}
		// 如果前面是非空白字符（比如上一个字段），就停下
		break
	}
	// 截取源码原始字符串
	structStr := l.content[start : stop+1]
	//处理内部非本次任务的内容
	l.result = replaceArrowsInStruct(structStr, l.funcName)
}

// 监听 outField 节点的进入事件
func (l *ArrowStructListener) EnterOutField(ctx *Wparser.OutFieldContext) {
	if ctx.ARROW() == nil {
		return // 没有箭头，跳过
	}

	exprCtx := ctx.Expr() // 拿到右侧表达式
	if exprCtx == nil {
		return
	}

	fc := exprCtx.FunctionCall()
	if fc == nil {
		return // 右侧不是函数调用
	}

	funcName := fc.ID().GetText()
	argText := ""
	if fc.GetChildCount() >= 3 {
		fullText := fc.GetText() // Getpath(ExtractSteps(in1))
		if len(fullText) > len(funcName)+2 {
			argText = fullText[len(funcName)+1 : len(fullText)-1] // 括号里面
		}
	}

	allin := fc.GetText() // 假设箭头已经在 outFieldContext 里

	l.Matches = append(l.Matches, []string{allin, funcName, argText})
}
