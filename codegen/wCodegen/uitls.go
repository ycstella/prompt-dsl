package wcodegen

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

// 提取表达式中第一个括号前的函数名
func extractFuncName(expr string) string {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return ""
	}
	// 找到第一个 "("
	if idx := strings.Index(expr, "("); idx > 0 {
		// 截取 "(" 之前的部分并去掉空格
		return strings.TrimSpace(expr[:idx])
	}
	return ""
}
func getLastPart(expr string) string {
	if expr == "" {
		return ""
	}
	if i := strings.LastIndex(expr, "."); i != -1 && i < len(expr)-1 {
		return expr[i+1:]
	}
	return expr
}
func findVarNameById(id string, nodes []*TaskNode) string {
	for _, n := range nodes {
		if n.Id == id {
			return n.VarName
		}
	}
	return "unknownVar"
}

// template
func renderTemplateWithLookup(nodes []*TaskNode, tmplStr string) string {
	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
		"findVarNameById": func(id string, nodes []*TaskNode) string {
			for _, n := range nodes {
				if n.Id == id {
					return n.VarName
				}
			}
			return "unknownVar"
		},
		"findReNameById": func(id string, nodes []*TaskNode) string {
			for _, n := range nodes {
				if n.Id == id {
					return n.ReturnType
				}
			}
			return "unknownVar"
		},
	}

	tmpl, err := template.New("tmpl").Funcs(funcMap).Parse(tmplStr)
	if err != nil {
		panic(fmt.Errorf("模板解析失败: %v", err))
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, nodes); err != nil {
		panic(fmt.Errorf("模板执行失败: %v", err))
	}
	return buf.String()
}
func logdfine() {
	logFile, err := os.OpenFile("wcodegen.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	// ✅ 将默认日志输出重定向到文件
	log.SetOutput(logFile)
}
func Logdfine() {
	logFile, err := os.OpenFile("wcodegen.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	// ✅ 将默认日志输出重定向到文件
	log.SetOutput(logFile)
}