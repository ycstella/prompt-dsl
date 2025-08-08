// promptdsl-core/astbuilder.go
package codegen

import (
	"codegen/parser"
	"fmt"

	// "log"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func ConvertASTtoPrompt(parseTree *parser.PromptFileContext, stream *antlr.CommonTokenStream, inputStream *antlr.InputStream) *PromptNode {
	result := &PromptNode{
		Vars:             make(map[string]interface{}), // 初始化空map
		SysNodes:         []Node{},
		UserNodes:        []Node{},
		ModuleDefs:       map[string][]Node{}, // 初始化 map
		InFields:         []FieldDef{},
		OutFields:        []FieldDef{},
		ModelFields:      []FieldDef{},
		BeforeCode:       "",
		AfterCode:        []string{},
		FixCode:          []string{},
		BeforeNodes:      []Node{},
		Goimport:         []goimport{},
		outputspectNodes: OutputSpecNode{},
	}
	fmt.Println("Building AST...")
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
			// modName := b.ID().GetText()
			// var contentNodes []Node
			// for _, mc := range b.AllModuleContent() {
			// 	mcCtx := mc.(*parser.ModuleContentContext)

			// 	for i := 0; i < mcCtx.GetChildCount(); i++ {
			// 		child := mcCtx.GetChild(i)
			// 		switch sub := child.(type) {
			// 		case *parser.TextLineContext:
			// 			contentNodes = append(contentNodes, &StringNode{Val: cleanQuotes(sub.GetText())})
			// 		case *parser.ParamPathContext:
			// 			contentNodes = append(contentNodes, &ParamNode{Path: cleanQuotes(sub.GetText())})
			// 		case *parser.IfStatementContext:
			// 			contentNodes = append(contentNodes, buildIfNode(sub))
			// 		case *parser.ExprContext:
			// 			contentNodes = append(contentNodes, &StringNode{Val: cleanQuotes(sub.GetText())}) //表达式
			// 		default:

			// 		}
			// 	}
			// }
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
			for _, param := range b.AllFieldDef() {
				name := param.ID().GetText()
				typ := "string"
				if param.Type_().GetText() != "" {
					fmt.Println("😊空字符")
					typ = param.Type_().GetText()
				}
				if typ == "[]" {
					typ = "[]string"
				}

				// 解析注解
				var annotations []string
				for _, ann := range param.AllAnnotation() {
					if ann.AnnotationArgs() != nil {
						for _, v := range ann.AnnotationArgs().AllAnnotationValue() {
							if s := v.STRING(); s != nil {
								annotations = append(annotations, strings.Trim(s.GetText(), "\""))
							} else if arr := v.ArrayLiteral(); arr != nil {
								for _, s := range arr.AllSTRING() {
									annotations = append(annotations, strings.Trim(s.GetText(), "\""))
								}
							}
						}
					}
				}

				result.InFields = append(result.InFields, FieldDef{
					Name:        name,
					Type:        typ,
					JsonName:    name,
					Annotations: annotations,
				})
			}
			fmt.Println("😅inNode:", result.InFields)

		case *parser.OutputSectionContext:

			// 先构建 defaultAnnotation map，方便查找
			defaultAnnoMap := buildDefaultAnnotationMap(b)

			// 检查是哪种 output 类型
			if structCtx := b.OutputStruct(); structCtx != nil {
				// var fields []FieldDef
				for _, field := range structCtx.AllFieldDef() {
					name := field.ID().GetText()
					typ := "string"
					if field.Type_().GetText() != "" {
						typ = field.Type_().GetText()
					}
					if typ == "[]" {
						typ = "[]string"
					}
					// 解析注解
					var annotations []string
					jsonName := name // 默认就是字段名
					//是否是模型需要字段
					var ismodel bool = false
					for _, ann := range field.AllAnnotation() {
						annName := ann.ID().GetText()
						//if annName == "modeloutput"将该字段加入modelfield
						if annName == "modeloutput"||annName == "model" {
							fmt.Println("annName：", annName)
							ismodel = true
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

								if annName == "jsonname"||annName == "jn" {
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

					field := FieldDef{
						Name:        name,
						Type:        typ,
						JsonName:    jsonName,
						Annotations: annotations,
					}

					result.OutFields = append(result.OutFields, field)
					if ismodel {
						fmt.Println("fmodeloutput")
						result.ModelFields = append(result.ModelFields, field)
					}
				}
			} else if mdCtx := b.OutputMarkdown(); mdCtx != nil {
				text := mdCtx.MARKDOWN().GetText()
				mdNode := &MarkdownNode{Content: cleanQuotes(text)}
				result.SysNodes = append(result.SysNodes, mdNode)
			}

		case *parser.BeforeSectionContext:
			// for _, bc := range b.AllBeforeContent() {
			// 	node := buildNodeFromBeforeContent(bc) // 你自定义的函数，返回 Node 接口实现
			// 	result.BeforeNodes = append(result.BeforeNodes, node)
			// }
		case *parser.AfterSectionContext:
			fmt.Println("😊AfterSection", extractRawText(b, stream))

			result.AfterCode = extractRawText(b, stream)
		case *parser.FixSectionContext:
			result.FixCode = extractRawText(b, stream)
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
	// fix := extractCodeBlocks(stream, "fix")
	// after := extractCodeBlocks(stream, "after")
	// fixPos := Range{stream.Get(fix.start).GetStart(), stream.Get(fix.end).GetStop() + 1}
	// afterPos := Range{stream.Get(after.start).GetStart(), stream.Get(after.end).GetStop() + 1}
	// result.FixCode =inputStream.GetText(fixPos.start, fixPos.end)
	// result.AfterCode =inputStream.GetText(afterPos.start, afterPos.end)

	return result
}

func BuildSysNodes(root antlr.Tree) []Node {
	moduleContentMap := make(map[string][]Node)

	// Step 1: 收集所有 moduleDef
	var collectModules func(antlr.Tree)
	collectModules = func(node antlr.Tree) {
		switch ctx := node.(type) {
		case *parser.ModuleDefContext:
			moduleName := ctx.ID().GetText()
			var nodes []Node
			for _, mc := range ctx.AllModuleContent() {
				for i := 0; i < mc.(antlr.Tree).GetChildCount(); i++ {
					child := mc.(antlr.Tree).GetChild(i)
					switch t := child.(type) {
					case *parser.TextLineContext:
						nodes = append(nodes, &StringNode{Val: cleanQuotes(t.GetText())})
					}
				}
			}
			moduleContentMap[moduleName] = nodes
		}
		for i := 0; i < node.GetChildCount(); i++ {
			collectModules(node.GetChild(i))
		}
	}
	collectModules(root)
	// Step 2: 提取 sysSection 中的模块名，展开为 AST 节点
	var result []Node
	var expandSys func(antlr.Tree)
	expandSys = func(node antlr.Tree) {
		switch ctx := node.(type) {
		case *parser.SystemSectionContext:
			for _, id := range ctx.AllID() {
				modName := id.GetText()
				if nodes, ok := moduleContentMap[modName]; ok {
					result = append(result, nodes...) // 展开合并
				} else {
					// 未定义模块也可以加提示节点
					result = append(result, &StringNode{Val: fmt.Sprintf("[Missing module: %s]", modName)})
				}
			}
		}
		for i := 0; i < node.GetChildCount(); i++ {
			expandSys(node.GetChild(i))
		}
	}
	expandSys(root)
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

// 构建一个 userContent 的 Node
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
			// fmt.Println("😊param path:", sub.GetText())
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
func buildExpr(exprCtx parser.IExprContext) *Expr {
	switch expr := exprCtx.(type) {
	case *parser.ExprContext:
		if param := expr.ParamPath(); param != nil {
			paramName := getParamName(param)
			return &Expr{Leaf: &paramName}
		}
		if str := expr.DASH_STRING(); str != nil {
			s := strings.Trim(str.GetText(), "-")
			return &Expr{Leaf: &s}
		}
		if num := expr.NUMBER(); num != nil {
			n := num.GetText()
			return &Expr{Leaf: &n}
		}
		if b := expr.BOOL(); b != nil {
			val := b.GetText()
			return &Expr{Leaf: &val}
		}
	}
	return &Expr{Op: ExprOp_None}
}

func cleanQuotes(s string) string {
	// 去掉开头和结尾的双引号（如果真的是配对的）
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		s = strings.TrimPrefix(s, "\"")
		s = strings.TrimSuffix(s, "\"")
	}
	s = strings.TrimPrefix(s, "-")
	return s
}

func getParamName(p parser.IParamPathContext) string {
	parts := []string{}
	for _, id := range p.AllID() {
		name := id.GetText()
		if name != "in" {
			parts = append(parts, name)
		}
	}
	return strings.Join(parts, ".")
}
func buildDefaultAnnotationMap(b *parser.OutputSectionContext) map[string][]string {
	defaultAnnotations := b.AllDefaultAnnotation()
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
