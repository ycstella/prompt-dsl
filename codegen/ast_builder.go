// promptdsl-core/astbuilder.go
package codegen

import (
	"fmt"
	"log"

	"strconv"
	"strings"

	"github.com/ycstella/prompt-dsl/codegen/parser"
	"github.com/ycstella/prompt-dsl/config"

	"github.com/antlr4-go/antlr/v4"
)

func ConvertASTtoPrompt(parseTree *parser.PromptFileContext, stream *antlr.CommonTokenStream, inputStream *antlr.InputStream) *PromptNode {
	config.InitLogger()
	result := NewPromptNode()
	fmt.Println("Building AST...")
	// 先处理 runexeDef
	runexeDefs := parseTree.AllRunexeDef()
	if len(runexeDefs) > 0 {
		result.RunExe = make([]string, len(runexeDefs))
		for i, runexeDef := range runexeDefs {
			txt := runexeDef.STRING().GetText()
			result.RunExe[i] = strings.Trim(txt, "\"")
		}
		log.Println("🚀 RunExe:", result.RunExe)
	}
	def := parseTree.PromptDef(0)
	for _, block := range def.AllPromptBlock() {
		fmt.Println("  Processing block:")
		child := block.GetChild(0)
		switch b := child.(type) {
		case *parser.SystemSectionContext:
			if len(b.AllID()) > 0 {
				// ID 模式，生成 ModuleRefNode
				for _, id := range b.AllID() {
					modName := id.GetText()
					result.SysNodes = append(result.SysNodes, &ModuleRefNode{Name: modName})
				}
				fmt.Println("result.SysNodes:", result.SysNodes)
			} else {
				for _, content := range b.AllSysContent() {
					result.SysNodes = append(result.SysNodes, BuildSysNodesC(content))
				}
			}
		case *parser.ModuleDefContext:
			// 收集模块内容
			modName := b.ID().GetText()
			var ModuleNode []Node
			for _, uc := range b.AllModuleContent() {
				node := BuildModuleNode(uc)
				ModuleNode = append(ModuleNode, node)
				if outSpec, ok := node.(*OutputSpecNode); ok {
					result.outputspectNodes = *outSpec
				}
			}
			result.ModuleDefs[modName] = ModuleNode
		case *parser.UserSectionContext:
			if len(b.AllID()) > 0 {
				// ID 模式，生成 ModuleRefNode
				for _, id := range b.AllID() {
					modName := id.GetText()
					result.UserNodes = append(result.UserNodes, &ModuleRefNode{Name: modName})
				}
				fmt.Println("result.UserNodes:", result.UserNodes)
			} else {
				for _, uc := range b.AllUserContent() {
					node := BuildUserNode(uc)
					result.UserNodes = append(result.UserNodes, node)
					if outSpec, ok := node.(*OutputSpecNode); ok {
						result.outputspectNodes = *outSpec
					}
				}
			}

		case *parser.InputSectionContext:
			// 解析输入字段，放到 result.InDef
			for _, field := range b.AllFieldDef() {
				field := processField(field, nil, result)
				result.InFields = append(result.InFields, field)
			}
			fmt.Println("😅inNode:", result.InFields)

		case *parser.OutputSectionContext:

			// 检查是哪种 output 类型

			if structCtx := b.OutputStruct(); structCtx != nil {
				// 先构建 defaultAnnotation map，方便查找
				defaultAnnoMap := buildDefaultAnnotationMap(b.AllDefaultAnnotation())
				// var fields []FieldDef
				for _, field := range structCtx.AllFieldDef() {
					field := processField(field, defaultAnnoMap, result)
					result.OutFields = append(result.OutFields, field)
				}
			} else if mdCtx := b.OutputMarkdown(); mdCtx != nil {
				// fmt.Println("md👀")
				// text := mdCtx.MARKDOWN().GetText()
				// mdNode := &MarkdownNode{Content: cleanQuotes(text)}
				// result.SysNodes = append(result.SysNodes, mdNode)
			}

		case *parser.BeforeSectionContext:
			result.BeforeCode = extractRawText(b, stream)
		case *parser.AfterSectionContext:
			fmt.Println("😊AfterSection", extractRawText(b, stream))

			result.AfterCode = extractRawText(b, stream)
		case *parser.FixSectionContext:
			result.FixCode = extractRawText(b, stream)
		case *parser.IteraterContext:
			if b.STRING() != nil {
				iteraterPath := b.STRING().GetText()
				result.iteraterPath, _ = strconv.Unquote(iteraterPath)
			}
		case *parser.LoopRangeSectionContext:
			if len(b.AllNUMBER()) > 0 {
				nums := b.AllNUMBER()
				var loopRange []int
				for _, n := range nums {
					text := n.GetText()
					val, err := strconv.Atoi(text)
					if err != nil {
						log.Printf("解析数字失败: %s, err: %v", text, err)
						continue
					}
					loopRange = append(loopRange, val)
				}
				result.loopRange = loopRange
			}
		case *parser.GoimportSectionContext:
			var imports []goimport
			for _, entry := range b.AllGoimportEntry() {
				var imp goimport
				if entry.ID() != nil {
					imp.Alias = entry.ID().GetText()
				} else {
					imp.Alias = "" // 没有别名
				}
				imp.Path = strings.Trim(entry.STRING().GetText(), `"`)
				imports = append(imports, imp)
			}
			result.Goimport = imports
		}
	}
	return result
}

func BuildSysNodesC(ctx parser.ISysContentContext) Node {
	nodeCtx := ctx.(*parser.SysContentContext)
	// 优先判断 ARRAY_OUTPUTSPEC（形如 []outputspec）
	if nodeCtx.ARRAY_OUTPUTSPEC() != nil {
		text := nodeCtx.ARRAY_OUTPUTSPEC().GetText()
		fmt.Println("😊ARRAY_OUTPUTSPEC:", text)

		// 直接去掉前缀 []，拿到实际类型名
		rawType := strings.TrimPrefix(text, "[]")
		return &OutputSpecNode{
			IsArray: true,
			RawTyp:  rawType,
		}
	}

	// 普通 OUTPUTSPEC（非数组形式）
	if nodeCtx.OUTPUTSPEC() != nil {
		text := nodeCtx.OUTPUTSPEC().GetText()
		fmt.Println("😊OUTPUTSPEC:", text)
		return &OutputSpecNode{
			IsArray: false,
			RawTyp:  text,
		}
	}

	for i := 0; i < nodeCtx.GetChildCount(); i++ {
		child := nodeCtx.GetChild(i)
		switch sub := child.(type) {
		case *parser.ParamPathContext:
			// fmt.Println("😊param path:", sub.GetText())
			return &ParamNode{Path: cleanQuotes(sub.GetText())}
		case *parser.TextLineContext:
			fmt.Println("😊stringtext:", sub.GetText())
			return &StringNode{Val: cleanQuotes(sub.GetText())}
		case *parser.IfStatementContext:
			fmt.Println("😊IfStatementContext:", sub.GetText())
			return buildIfNode(sub)
		case *parser.ExprContext:
			return &StringNode{Val: cleanQuotes(sub.GetText())} // 临时
		default:
			continue
		}
	}
	return nil
}

// 构建一个 userContent Node
func BuildUserNode(ctx parser.IUserContentContext) Node {
	// fmt.Println("😊buildUserNode:")
	nodeCtx := ctx.(*parser.UserContentContext)
	// 优先判断 ARRAY_OUTPUTSPEC（形如 []outputspec）
	if nodeCtx.ARRAY_OUTPUTSPEC() != nil {
		text := nodeCtx.ARRAY_OUTPUTSPEC().GetText()
		fmt.Println("😊ARRAY_OUTPUTSPEC:", text)

		// 直接去掉前缀 []，拿到实际类型名
		rawType := strings.TrimPrefix(text, "[]")
		return &OutputSpecNode{
			IsArray: true,
			RawTyp:  rawType,
		}
	}

	// 普通 OUTPUTSPEC（非数组形式）
	if nodeCtx.OUTPUTSPEC() != nil {
		text := nodeCtx.OUTPUTSPEC().GetText()
		fmt.Println("😊OUTPUTSPEC:", text)
		return &OutputSpecNode{
			IsArray: false,
			RawTyp:  text,
		}
	}

	for i := 0; i < nodeCtx.GetChildCount(); i++ {
		child := nodeCtx.GetChild(i)
		switch sub := child.(type) {
		case *parser.ParamPathContext:
			// fmt.Println("😊param path:", sub.GetText())
			return &ParamNode{Path: cleanQuotes(sub.GetText())}
		case *parser.TextLineContext:
			fmt.Println("😊stringtext:", sub.GetText())
			return &StringNode{Val: cleanQuotes(sub.GetText())}
		case *parser.IfStatementContext:
			fmt.Println("😊IfStatementContext:", sub.GetText())
			return buildIfNode(sub)
		case *parser.ExprContext:
			return &StringNode{Val: cleanQuotes(sub.GetText())} // 临时
		default:
			continue
		}
	}
	return nil
}

func BuildModuleNode(ctx parser.IModuleContentContext) Node {
	nodeCtx := ctx.(*parser.ModuleContentContext)
	// 优先判断 ARRAY_OUTPUTSPEC（形如 []outputspec）
	if nodeCtx.ARRAY_OUTPUTSPEC() != nil {
		text := nodeCtx.ARRAY_OUTPUTSPEC().GetText()
		fmt.Println("😊ARRAY_OUTPUTSPEC:", text)

		// 直接去掉前缀 []，拿到实际类型名
		rawType := strings.TrimPrefix(text, "[]")
		return &OutputSpecNode{
			IsArray: true,
			RawTyp:  rawType,
		}
	}

	// 普通 OUTPUTSPEC（非数组形式）
	if nodeCtx.OUTPUTSPEC() != nil {
		text := nodeCtx.OUTPUTSPEC().GetText()
		fmt.Println("😊OUTPUTSPEC:", text)
		return &OutputSpecNode{
			IsArray: false,
			RawTyp:  text,
		}
	}

	for i := 0; i < nodeCtx.GetChildCount(); i++ {
		child := nodeCtx.GetChild(i)
		switch sub := child.(type) {
		case *parser.ParamPathContext:
			fmt.Println("😊param path:", sub.GetText())
			return &ParamNode{Path: cleanQuotes(sub.GetText())}
		case *parser.TextLineContext:
			fmt.Println("😊stringtext:", sub.GetText())
			return &StringNode{Val: cleanQuotes(sub.GetText())}
		case *parser.IfStatementContext:
			fmt.Println("😊IfStatementContext:", sub.GetText())
			return buildIfNode(sub)
		case *parser.ForRangeWithIndexContext:
			fmt.Println("😊ForRangeWithIndexContext:", sub.GetText())
			return buildForNode(sub)
		case *parser.ForTraditionalContext:
			fmt.Println("😊ForTraditionalContext:", sub.GetText())
			return buildForNode(sub)
		case *parser.ForRangeNoIndexContext:
			fmt.Println("😊ForRangeNoIndexContext:", sub.GetText())
			return buildForNode(sub)
		case *parser.SwitchStatementContext:
			fmt.Println("😊SwitchStatementContext:", sub.GetText())
			return buildSwitchNode(sub)
		case *parser.ExprContext:
			return &StringNode{Val: cleanQuotes(sub.GetText())}
		default:
			continue
		}
	}
	return nil
}

func buildIfNode(ctx *parser.IfStatementContext) *IfNode {
	condList := ctx.Condition()
	// // 断言接口为具体类型
	cctx := condList.(*parser.ConditionContext)
	condition := cctx.GetText()
	// fmt.Println("😊Condition:", condition)
	// condExpr := buildExprFromCondition(cctx)

	var thenNodes []Node
	Thencontent := ctx.AllThencontent()
	for _, uc := range Thencontent {
		thenNodes = append(thenNodes, BuildUserNode(uc.UserContent()))
	}
	var elseNodes []Node
	// 遍历 else 分支
	elsecontent := ctx.AllElsecontent()
	for _, uc := range elsecontent {
		elseNodes = append(elseNodes, BuildUserNode(uc.UserContent()))
	}

	return &IfNode{
		Condition: condition,
		Then:      thenNodes,
		Else:      elseNodes,
	}
}
func buildForNode(ctx parser.IForStatementContext) *ForNode {
	var initStr, condStr, updateStr string
	var keyStr, valStr, rangeExprStr string
	var forType string
	var body []Node

	switch fc := ctx.(type) {

	case *parser.ForTraditionalContext:
		// C-style: for i := 0; i < 10; i++
		init := fc.GetInit().(parser.IAssignExprContext)
		update := fc.GetUpdate().(parser.IUpdateExprContext)
		condition := fc.Condition().(*parser.ConditionContext)

		initStr = init.GetText()
		condStr = condition.GetText()
		updateStr = update.GetText()
		forType = "traditional"

		for _, uc := range fc.AllForcontent() {
			body = append(body, BuildUserNode(uc.UserContent()))
		}

	case *parser.ForRangeWithIndexContext:
		keyStr = fc.GetKey().GetText()
		valStr = fc.GetVal().GetText()
		rangeExprStr = fc.GetIterable().GetText()
		forType = "rangeWithIndex"

		for _, uc := range fc.AllForcontent() {
			body = append(body, BuildUserNode(uc.UserContent()))
		}

	case *parser.ForRangeNoIndexContext:
		valStr = fc.GetVal().GetText()
		rangeExprStr = fc.GetIterable().GetText()
		forType = "rangeNoIndex"

		for _, uc := range fc.AllForcontent() {
			body = append(body, BuildUserNode(uc.UserContent()))
		}

	default:
		panic("unknown for statement type")
	}

	return &ForNode{
		ForType: forType, // 新增字段，标明类型
		Init:    initStr,
		Cond:    condStr,
		Post:    updateStr,
		Key:     keyStr,
		Val:     valStr,
		Range:   rangeExprStr,
		Body:    body,
	}
}
func buildSwitchNode(ctx parser.ISwitchStatementContext) *SwitchNode {
	condList := ctx.Condition()
	cctx := condList.(*parser.ConditionContext)
	condition := cctx.GetText()

	caseList := ctx.AllSwitchCase()
	var cases []CasePair
	for _, caseitem := range caseList {
		var casest CasePair
		casest.Case = caseitem.Condition().GetText()
		fmt.Println("😊casestr:", casest.Case)
		for _, uc := range caseitem.AllUserContent() {
			casest.Body = append(casest.Body, BuildUserNode(uc))
		}
		cases = append(cases, casest)
	}
	defult := ctx.SwitchDefault()
	var defultNode []Node
	for _, uc := range defult.AllUserContent() {
		defultNode = append(defultNode, BuildUserNode(uc))
	}
	return &SwitchNode{
		Switch:  condition,
		Cases:   cases,
		Default: defultNode,
	}
}

// 去掉开头和结尾的双引号（配对的）
func cleanQuotes(s string) string {
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		s = strings.TrimPrefix(s, "\"")
		s = strings.TrimSuffix(s, "\"")
	}
	s = strings.TrimPrefix(s, "-")
	return s
}

func buildDefaultAnnotationMap(defaultAnnotations []parser.IDefaultAnnotationContext) map[string][]string {
	defaultAnnoMap := make(map[string][]string)
	for _, defAnn := range defaultAnnotations {
		name := defAnn.ID().GetText()
		var vals []string
		if defAnn.AnnotationArgs() != nil {
			for _, v := range defAnn.AnnotationArgs().AllAnnotationValue() {
				if s := v.STRING(); s != nil {
					raw := s.GetText()
					unquoted, err := strconv.Unquote(raw)
					if err != nil {
						unquoted = raw
					}
					vals = append(vals, unquoted)
				} else if arr := v.ArrayLiteral(); arr != nil {
					var parts []string
					for _, s := range arr.AllSTRING() {
						raw := s.GetText()
						unquoted, err := strconv.Unquote(raw)
						if err != nil {
							unquoted = raw
						}
						parts = append(parts, unquoted)
					}
					vals = append(vals, strings.Join(parts, ","))
				}
			}
		}
		defaultAnnoMap[name] = append(defaultAnnoMap[name], vals...)
	}
	return defaultAnnoMap
}

func processField(field parser.IFieldDefContext, defaultAnnoMap map[string][]string, result *PromptNode) FieldDef {
	name := field.ID().GetText()
	typ := "string"
	if field.Type_().GetText() != "" {
		typ = field.Type_().GetText()
		// log.Println(typ)
	}
	if typ == "[]" {
		typ = "[]string"
	}
	var subfieldlist []FieldDef
	if strings.HasPrefix(typ, "[]struct") {
		// log.Println("is []struct")
		typ = "[]struct"
		typeCtx := field.Type_().Type_()
		// log.Println("typeCtx:",typeCtx)
		for i := 0; i < typeCtx.GetChildCount(); i++ {
			child := typeCtx.GetChild(i)
			// 判断子节点是不是 FieldDef
			if subFieldCtx, ok := child.(parser.IFieldDefContext); ok {
				// fmt.Println("subing:😒",subFieldCtx.ID)
				subfield := processField(subFieldCtx, defaultAnnoMap, result)
				subfieldlist = append(subfieldlist, subfield)
			}
		}
		subfields := Subfield{
			Name:   name,
			Fields: subfieldlist,
		}
		result.SubFields = append(result.SubFields, subfields)
	}
	if strings.HasPrefix(typ, "struct") {
		typ = "struct"
		typeCtx := field.Type_()
		for i := 0; i < typeCtx.GetChildCount(); i++ {
			child := typeCtx.GetChild(i)
			// 判断子节点是不是 FieldDef
			if subFieldCtx, ok := child.(parser.IFieldDefContext); ok {
				// fmt.Println("subing:😒",subFieldCtx.ID)
				subfield := processField(subFieldCtx, defaultAnnoMap, result)
				subfieldlist = append(subfieldlist, subfield)
			}
		}
		subfields := Subfield{
			Name:   name,
			Fields: subfieldlist,
		}
		result.SubFields = append(result.SubFields, subfields)
	}
	// 解析注解
	var annotations []string
	jsonName := name // 默认就是字段名
	//是否是模型需要字段
	var ismodel bool = false
	for _, ann := range field.AllAnnotation() {
		annName := ann.ID().GetText()

		//if annName == "modeloutput"将该字段加入modelfield
		if annName == "modeloutput" || annName == "model" {
			fmt.Println("annName：", annName)
			ismodel = true
		}
		if annName == "derived" {
			fmt.Println("derivedS😢:")
			annotations = append(annotations, "derived")
		}
		if annName == "outignore" {
			fmt.Println("outignore:")
			annotations = append(annotations, "outignore")
		}
		if ann.AnnotationArgs() != nil {
			for _, v := range ann.AnnotationArgs().AllAnnotationValue() {
				var val string

				if s := v.STRING(); s != nil {
					raw := s.GetText()
					unquoted, err := strconv.Unquote(raw)
					if err != nil {
						unquoted = raw
					}
					val = unquoted
				} else if arr := v.ArrayLiteral(); arr != nil {
					// 拼接数组内容
					var parts []string
					for _, s := range arr.AllSTRING() {
						raw := s.GetText()
						unquoted, err := strconv.Unquote(raw)
						if err != nil {
							unquoted = raw
						}
						parts = append(parts, unquoted)
					}
					val = strings.Join(parts, ",") // 或保留原结构
				}

				if annName == "jsonname" || annName == "jn" {
					jsonName = val
				} else {
					annotations = append(annotations, val)
				}

			}
		}
		// 再加 defaultAnnotation 中对应注解的参数（如果有）
		if defVals, ok := defaultAnnoMap[annName]; ok {
			annotations = append(annotations, defVals...)
		}

	}
	pfield := FieldDef{
		Name:        name,
		Type:        typ,
		JsonName:    jsonName,
		Ismodel:     ismodel,
		Annotations: annotations,
		SubFields:   subfieldlist,
	}
	if pfield.Ismodel {
		fmt.Println("fmodeloutput")
		result.ModelFields = append(result.ModelFields, pfield)
	}

	return pfield
}
