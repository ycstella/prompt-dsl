package wcodegen

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
	Wparser "github.com/ycstella/prompt-dsl/codegen/Wparser"
)

type Range struct {
	functionName string
	start        int
	stop         int
	code         string
}
type ArrowStructListener struct {
	*Wparser.BaseWorkflowDSLParserListener
	funcName     string
	content      string
	result       string
	Matches      [][]string
	ignoreRanges []Range
}

func NewArrowStructListener(funcName, content string) *ArrowStructListener {
	return &ArrowStructListener{
		funcName: funcName,
		content:  content,
		Matches:  make([][]string, 0),
	}
}

func replaceArrowsInStruct(structStr, funcName string) string {
	var sb strings.Builder
	i := 0
	n := len(structStr)

	for i < n {
		// 找到下一个 "<-"
		idx := strings.Index(structStr[i:], "<-")
		if idx == -1 {
			sb.WriteString(structStr[i:])
			break
		}
		idx += i

		sb.WriteString(structStr[i:idx])
		//跳箭头+空白
		j := idx + 2
		for j < n && unicode.IsSpace(rune(structStr[j])) {
			j++
		}

		// 读取函数名（ID）
		startName := j
		for j < n {
			r := structStr[j]
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_' || (r >= '0' && r <= '9' || r == '$' || r == '.' ||
				r == '[' || r == ']' && j > startName) {
				j++
				continue
			}

			//下划线和.[]不能一起用
			// 遇到空格但后面可能是 '(' -> 允许跳过一次
			//if
			if unicode.IsSpace(rune(r)) {
				j++
				continue
			}
			break
		}
		funcIdent := strings.TrimSpace(structStr[startName:j])
		// 跳过空白
		for j < n && unicode.IsSpace(rune(structStr[j])) {
			j++
		}
		//匹配完整调用
		open := 0
		k := j
		for k < n {
			c := structStr[k]
			if c == '(' {
				open++
			} else if c == ')' {
				open--
				if open == 0 {
					k++
					break
				}
			}
			k++
		}

		if open != 0 {
			// 写回 "<-" 并推进
			sb.WriteString("<-")
			i = idx + 2
			continue
		}

		fullCall := structStr[idx:k]
		fmt.Println("found arrow call:", fullCall)

		if funcIdent == funcName {

		} else {

			sb.WriteString("@ignore")
		}

		// 继续扫描，从 k 开始
		i = k
	}

	return sb.String()
}
func (l *ArrowStructListener) EnterFunctionCall(ctx *Wparser.FunctionCallContext) {
	id := ctx.ID().GetText()

	if l.funcName == "" {
		return
	}
	// fmt.Println("id:", id)
	// fmt.Println("funcName:", l.funcName)
	ignoreRanges := Range{
		functionName: id,
	}

	// 找到 funcName 调用
	start := ctx.GetStart().GetStart()
	stop := ctx.GetStop().GetStop()

	// 沿父节点向上找到最外层 struct 或 outField
	parent := ctx.GetParent()

	for parent != nil {
		fmt.Printf("parent dynamic type: %s\n", reflect.TypeOf(parent))
		for i := 0; i < parent.GetChildCount(); i++ {
			child := parent.GetChild(i)
			fmt.Printf("child[%d] dynamic type: %s\n", i, reflect.TypeOf(child))
		}

		switch n := parent.(type) {
		case *Wparser.OutFieldContext:
			fmt.Println("parent:", n.GetText())

			start = n.GetStart().GetStart()
			stop = n.GetStop().GetStop()

			goto FOUND
		case *Wparser.TypeContext:
			start = n.GetStart().GetStart()
			stop = n.GetStop().GetStop()
			goto FOUND
		
		case *Wparser.FunctionCallContext:
			return
		}
		parent = parent.(antlr.ParserRuleContext).GetParent()

	}

FOUND:
	// fmt.Println("strat:",start)
	ignoreRanges.stop = stop

	fmt.Println("找到了")
	content := l.content
	// 向前扩展 start，使其包含行首缩进
	for start > 0 {
		ch := content[start-1]
		if ch == ' ' || ch == '\t' {
			start--
			continue
		}
		if ch == '\n' || ch == '\r' {
			break
		}
		// 如果前面是非空白字符（比如上一个字段）
		break
	}
	ignoreRanges.start = start
	// 截取源码原始字符串
	structStr := l.content[start : stop+1]
	ignoreRanges.code = structStr
	l.ignoreRanges = append(l.ignoreRanges, ignoreRanges)
	for _, value := range l.ignoreRanges {
		fmt.Println("ignoreRanges:", value)
	}

}

func extractTargetContent(content string, funcName string, ignoreRanges []Range) string {
	var targetStart, targetStop int
	found := false

	// 找到目标函数的范围
	for _, v := range ignoreRanges {
		if v.functionName == funcName {
			targetStart = v.start
			targetStop = v.stop
			found = true
			break
		}
	}

	if !found {
		fmt.Println("⚠️ 未找到目标函数范围")
		return ""
	}

	fmt.Println("目标函数范围Start:", targetStart, "stop:", targetStop)
	if targetStart < 0 {
		targetStart = 0
	}
	if targetStop >= len(content) {
		targetStop = len(content) - 1
	}

	targetBlock := []byte(content[targetStart : targetStop+1])
	log.Println("name:", funcName)
	log.Println("targetBlock:", string(targetBlock))

	// 收集所有非目标函数的区间（在目标范围内的）
	var toRemove []Range
	log.Println("targetrange:", targetStart, ",", targetStop)
	for _, v := range ignoreRanges {
		if v.functionName != funcName &&
			v.start >= targetStart && v.stop <= targetStop {
			toRemove = append(toRemove, v)
			// log
			log.Println("toRemove:", v.start, ",", v.stop)
		}
	}
	// 从后往前删除，防止索引错位
	sort.Slice(toRemove, func(i, j int) bool {
		return toRemove[i].start > toRemove[j].start
	})

	for _, r := range toRemove {
		innerStart := r.start - targetStart
		innerStop := r.stop - targetStart
		if innerStart < 0 || innerStop >= len(targetBlock) {
			continue
		}
		// 删除对应子区间
		targetBlock = append(targetBlock[:innerStart], targetBlock[innerStop+1:]...)
	}

	return string(targetBlock)
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
		fullText := fc.GetText()
		if len(fullText) > len(funcName)+2 {
			argText = fullText[len(funcName)+1 : len(fullText)-1] // 括号里面
		}
	}

	allin := fc.GetText()

	l.Matches = append(l.Matches, []string{allin, funcName, argText})
}
