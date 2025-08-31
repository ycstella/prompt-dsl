package codegen

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"unicode"

	"regexp"

	// "path/filepath"
	"strconv"

	"github.com/ycstella/prompt-dsl/codegen/parser"

	// "runtime"

	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// 构建输出规范文本，用于根据字段列表生成 JSON 字符串表示。
func BuildModelOutputSpecLines(fields []FieldDef, asArray bool) []string {
	var lines []string
	if asArray {
		lines = append(lines, "```json")
		lines = append(lines, "[")
		lines = append(lines, "  {")
	} else {
		lines = append(lines, "```json")
		lines = append(lines, "{")
	}

	for i, f := range fields {
		lines = append(lines, writefield(f, fields, i, "    ")...)
	}

	if asArray {
		lines = append(lines, "  }")
		lines = append(lines, "]")
	} else {
		lines = append(lines, "}")
	}
	lines = append(lines, "```")

	return lines
}
func writefield(f FieldDef, fields []FieldDef, i int, indent string) []string {
	var lines []string

	// struct 类型特殊处理
	if f.Type == "struct" {
		// 普通结构体
		line := fmt.Sprintf("%s\"%s\": {  // %s", indent, f.JsonName, strings.Join(f.Annotations, ","))
		lines = append(lines, line)

		for idx, sub := range f.SubFields {
			lines = append(lines, writefield(sub, f.SubFields, idx, indent+"    ")...)
		}

		closing := indent + "}"
		if i < len(fields)-1 {
			closing += ","
		}
		lines = append(lines, closing)
		return lines
	}

	// []struct 特殊处理
	if strings.HasPrefix(f.Type, "[]") && strings.Contains(f.Type, "struct") {
		line := fmt.Sprintf("%s\"%s\": [  // %s", indent, f.JsonName, strings.Join(f.Annotations, ","))
		lines = append(lines, line)

		// 生成示例元素
		lines = append(lines, indent+"    {")
		for idx, sub := range f.SubFields {
			lines = append(lines, writefield(sub, f.SubFields, idx, indent+"        ")...)
		}
		lines = append(lines, indent+"    }")

		closing := indent + "]"
		if i < len(fields)-1 {
			closing += ","
		}
		lines = append(lines, closing)
		return lines
	}

	// 普通类型
	var example string
	switch f.Type {
	case "string":
		example = "\"\""
	case "int":
		example = "0"
	case "float":
		example = "0.0"
	case "bool":
		example = "false"
	default:
		if strings.HasPrefix(f.Type, "[]") {
			elemType := f.Type[2:]
			switch elemType {
			case "string":
				example = "[\"\"]"
			case "int":
				example = "[0]"
			case "float":
				example = "[0.0]"
			case "bool":
				example = "[false]"
			default:
				example = "[]"
			}
		} else {
			example = "{}"
		}
	}

	// 普通字段行
	line := fmt.Sprintf("%s\"%s\": %s  // %s", indent, f.JsonName, example, strings.Join(f.Annotations, ","))
	if i < len(fields)-1 {
		line += ","
	}
	lines = append(lines, line)
	return lines
}

func getCurrentPackageName() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "main"
	}
	// 获取路径中的目录名作为“包名”
	dir := filepath.Base(filepath.Dir(file))
	return dir
}

// after.fix node处理
func extractRawText(ctx antlr.ParserRuleContext, tokens *antlr.CommonTokenStream) []string {

	startIdx := ctx.GetStart().GetTokenIndex()
	stopIdx := ctx.GetStop().GetTokenIndex()

	// 获取完整 token 列表
	allTokens := tokens.GetAllTokens()

	// 切片截取 ctx 范围内的 token
	if startIdx < 0 || stopIdx >= len(allTokens) || startIdx > stopIdx {
		return nil
	}
	var builder strings.Builder
	for _, tok := range allTokens[startIdx : stopIdx+1] {
		builder.WriteString(tok.GetText())
	}
	code := builder.String()
	rePrefix := regexp.MustCompile(`(?i)^\s*(after|fix|before)\s*{`)
	code = rePrefix.ReplaceAllString(code, "")
	code = strings.TrimSuffix(code, "}")
	return []string{code}
}

// 字段首字母大写
func capitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// 判断code中是否有包构造import
func inferImportsFromCode(code string) []string {
	importSet := map[string]struct{}{}

	for prefix, pkg := range symbolToImport {
		if strings.Contains(code, prefix) {
			importSet[pkg] = struct{}{}
		}
	}

	var imports []string
	for pkg := range importSet {
		imports = append(imports, pkg)
	}
	// sort.Strings(imports) // 可选：让 import 有序
	return imports
}

// import（带别名）
func renderImportSectionWithAlias(goimports []goimport, pkgs []string) string {
	var b strings.Builder
	b.WriteString("import (\n")

	// 先把带别名的 goimports 写进去
	for _, imp := range goimports {
		if imp.Alias != "" {
			b.WriteString(fmt.Sprintf("\t%s \"%s\"\n", imp.Alias, imp.Path))
		} else {
			b.WriteString(fmt.Sprintf("\t\"%s\"\n", imp.Path))
		}
	}

	// 把纯路径 pkgs 里没有在 goimports 里出现的路径补上（无别名）
	exist := map[string]bool{}
	for _, imp := range goimports {
		exist[imp.Path] = true
	}
	for _, pkg := range pkgs {
		if !exist[pkg] {
			b.WriteString(fmt.Sprintf("\t\"%s\"\n", pkg))
		}
	}

	b.WriteString(")\n\n")
	return b.String()
}
func extractFieldDef(field parser.IFieldDefContext, defaultAnnoMap map[string][]string) FieldDef {
	name := field.ID().GetText()
	typ := field.Type_().GetText()
	jsonName := name
	var annotations []string
	var subFields []FieldDef

	// typeCtx := field.Type_()

	// // 判断是否为 struct 类型
	// if structType := typeCtx.GetChild(0); structType != nil {
	// 	if structKeyword, ok := structType.(antlr.TerminalNode); ok && structKeyword.GetText() == "struct" {
	// 		// 手动访问 struct 的字段定义
	// 		// 遍历 Type_() 的子节点，找出 fieldDef
	// 		for i := 0; i < typeCtx.GetChildCount(); i++ {
	// 			child := typeCtx.GetChild(i)
	// 			if fdCtx, ok := child.(*parser.FieldDefContext); ok {
	// 				subFields = append(subFields, extractFieldDef(fdCtx, defaultAnnoMap))
	// 			}
	// 		}
	// 	}
	// }

	// 注解处理（不变）
	for _, ann := range field.AllAnnotation() {
		annName := ann.ID().GetText()
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
					var parts []string
					for _, s := range arr.AllSTRING() {
						raw := s.GetText()
						unquoted, err := strconv.Unquote(raw)
						if err != nil {
							unquoted = raw
						}
						parts = append(parts, unquoted)
					}
					val = strings.Join(parts, ",")
				}
				if annName == "jsonname" {
					jsonName = val
				} else {
					annotations = append(annotations, val)
				}
			}
		}
		if defVals, ok := defaultAnnoMap[annName]; ok {
			annotations = append(annotations, defVals...)
		}
	}

	return FieldDef{
		Name:        name,
		Type:        typ,
		JsonName:    jsonName,
		Annotations: annotations,
		SubFields:   subFields,
	}
}

type Range struct {
	start int
	end   int
}

// Function to extract code blocks from the token stream
func extractCodeBlocks(tokens *antlr.CommonTokenStream, typ string) Range {

	// Get all tokens
	allTokens := tokens.GetAllTokens()

	ret := Range{-1, -1}

	t := parser.PromptDSLLexerFIX
	if typ == "after" {
		t = parser.PromptDSLLexerAFTER
	}
	fmt.Println("t:", t)
	// Track brace nesting level
	braceLevel := 0

	for i, token := range allTokens {
		typel := token.GetTokenType()
		fmt.Println("typel:", typel)
		if token.GetTokenType() == t {
			fmt.Println("😮")
			ret.start = i
			braceLevel = 0
			// Find the opening brace
			for j := i + 1; j < len(allTokens); j++ {
				if allTokens[j].GetTokenType() == parser.PromptDSLParserLBRACE {
					// Find the matching closing brace
					for k := j + 1; k < len(allTokens); k++ {
						if allTokens[k].GetTokenType() == parser.PromptDSLParserLBRACE {
							braceLevel++
						} else if allTokens[k].GetTokenType() == parser.PromptDSLParserRBRACE {
							if braceLevel == 0 {
								ret.end = k
								break
							}
							braceLevel--
						}
					}
					break
				}
			}
		}
	}

	return ret
}
func slej() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// 调用 go get 安装额外的依赖
func installGoImports(goimports []goimport, workDir string) error {
	for _, pkg := range goimports {
		// 标准库如 "fmt"、"os" 不需要 go get
		if isStandardPackage(pkg.Path) {
			continue
		}

		cmd := exec.Command("go", "get", pkg.Path)
		cmd.Dir = workDir
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to go get %s: %v\nOutput: %s", pkg, err, string(out))
		}
	}
	return nil
}

// 简单判断是不是标准库
func isStandardPackage(pkg string) bool {
	stdPkgs := map[string]bool{
		"fmt": true, "os": true, "io": true, "strings": true, "time": true, "bytes": true,
	}
	return stdPkgs[pkg]
}
func FixLatex[T any](response string) (T, error) {
	// 用strings.Builder手动替换单反斜杠
	fmt.Println("response:", response)
	var results T
	err := json.Unmarshal([]byte(response), &results)
	if err != nil {
		var buf strings.Builder
		for i := 0; i < len(response); i++ {
			if response[i] == '\\' {
				// 判断是否有下一个字符
				if i+1 < len(response) {
					next := response[i+1]
					// 如果是两个连续的反斜杠
					if next == '\\' {
						// 再判断第三个字符是否存在，且不是字母或反斜杠
						if i+2 >= len(response) || !(unicode.IsLetter(rune(response[i+2])) || response[i+2] == '\\') {
							// 变成 4 个斜杠
							buf.WriteString(`\\\\`)
						} else {
							// 保留原样 2 个斜杠
							buf.WriteString(`\\`)
							if response[i+3] == '\\' {
								i++
							}
						}
						i++ // 跳过下一个斜杠
					} else if next == '"' {
						// 保留一个反斜杠
						buf.WriteByte('\\')
						i++ // 跳过 "
					} else {
						// fmt.Println("last char is \\")
						// 单独的 \，不是合法转义，变成两个
						buf.WriteString(`\\`)
					}
				} else {
					// 最后一个字符是反斜杠，补一个
					buf.WriteString(`\`)
				}
			} else {
				buf.WriteByte(response[i])
			}
		}
		fixed := buf.String()
		fmt.Println("fixed:", fixed)

		err := json.Unmarshal([]byte(fixed), &results)
		if err != nil {
			return results, err // 返回T的零值和错误
		}
		return results, nil
	}
	return results, nil
}

// 默认直接unmarshall
func FixAuto[T any](response string) (T, error) {
	var results T
	err := json.Unmarshal([]byte(response), &results)
	log.Println("FixAuto:解析模型回复失败：", err)
	return results, err
}

//不同名同结构的结构体转换
func CopyStructRecursive(src, dest interface{}) {
	srcVal := reflect.ValueOf(src)
	destVal := reflect.ValueOf(dest)

	if srcVal.Kind() != reflect.Ptr || srcVal.IsNil() {
		log.Fatal("src 必须是非 nil 指针")
	}
	if destVal.Kind() != reflect.Ptr || destVal.IsNil() {
		log.Fatal("dest 必须是非 nil 指针")
	}

	copyValue(srcVal.Elem(), destVal.Elem())
}

func copyValue(srcVal, destVal reflect.Value) {
	switch srcVal.Kind() {
	case reflect.Struct:
		for i := 0; i < srcVal.NumField(); i++ {
			srcField := srcVal.Field(i)
			fieldName := srcVal.Type().Field(i).Name
			destField := destVal.FieldByName(fieldName)
			if !destField.IsValid() || !destField.CanSet() {
				continue
			}
			copyValue(srcField, destField)
		}
	case reflect.Slice:
		if srcVal.IsNil() {
			destVal.Set(reflect.Zero(destVal.Type()))
			return
		}
		newSlice := reflect.MakeSlice(destVal.Type(), srcVal.Len(), srcVal.Len())
		for i := 0; i < srcVal.Len(); i++ {
			copyValue(srcVal.Index(i), newSlice.Index(i))
		}
		destVal.Set(newSlice)
	case reflect.Ptr:
		if srcVal.IsNil() {
			destVal.Set(reflect.Zero(destVal.Type()))
			return
		}
		if destVal.IsNil() {
			destVal.Set(reflect.New(destVal.Type().Elem()))
		}
		copyValue(srcVal.Elem(), destVal.Elem())
	default:
		if srcVal.Type().AssignableTo(destVal.Type()) {
			destVal.Set(srcVal)
		} else {
			// 基础类型不同但可以转换，可以手动扩展
			// log.Printf("⚠️ 类型不匹配: %s -> %s\n", srcVal.Type(), destVal.Type())
		}
	}
}