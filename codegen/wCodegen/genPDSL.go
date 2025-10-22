package wcodegen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
	Wparser "github.com/ycstella/prompt-dsl/codegen/Wparser"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

type PDSLFile struct {
	callname string
	out      string
	in       string
}

type WPDSLGener struct {
	WorkflowContent string
	OutputDir       string
	matches         [][]string
	tree            Wparser.IWorkflowContext
	order           []string
	PDSLFiles        []PDSLFile
}

func NewWPDSLGener(workflowContent, outputDir string) *WPDSLGener {
	return &WPDSLGener{
		WorkflowContent: workflowContent,
		OutputDir:       outputDir,
	}
}
func (g *WPDSLGener) extractCalls() error {

	is := antlr.NewInputStream(g.WorkflowContent)
	lexer := Wparser.NewWorkflowDSLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := Wparser.NewWorkflowDSLParser(stream)
	g.tree = p.Workflow()
	l := NewArrowStructListener("", "")
	antlr.ParseTreeWalkerDefault.Walk(l, g.tree)
	g.matches = l.Matches
	//该正则无法完全满足<-xxx()需要无法嵌套
	// re := regexp.MustCompile(`<-\s*([A-Za-z_][A-Za-z0-9_]*)\s*\(\s*([A-Za-z_][A-Za-z0-9_]*)\s*\)`)
	// g.matches = re.FindAllStringSubmatch(g.WorkflowContent, -1)
	if len(g.matches) == 0 {
		return fmt.Errorf("no model calls found in workflow content")
	}
	return nil
}

func (g *WPDSLGener) extractIntStructFields(name string) string {
	re := regexp.MustCompile(fmt.Sprintf(`%s\s*\{([^{}]*\{[^{}]*\}[^{}]*)*[^{}]*\}`, name))
	match := re.FindString(g.WorkflowContent)
	if match == "" {
		return "in Struct匹配失败"
	}
	// 去掉最外层 name { }
	start := strings.Index(match, "{")
	end := strings.LastIndex(match, "}")
	return strings.TrimSpace(match[start+1 : end])
}

func replaceArrowsInStruct(structStr, funcName string) string {
	var sb strings.Builder
	i := 0
	n := len(structStr)

	for i < n {
		// 找到下一个 "<-"
		idx := strings.Index(structStr[i:], "<-")
		if idx == -1 {
			// 没有更多箭头，写入剩余并结束
			sb.WriteString(structStr[i:])
			break
		}
		idx += i // 转为全局索引

		// 把 "<-" 之前的内容先写入结果
		sb.WriteString(structStr[i:idx])

		j := idx + 2
		// 跳过空白
		for j < n && unicode.IsSpace(rune(structStr[j])) {
			j++
		}

		// 读取函数名（ID）
		startName := j
		for j < n {
			r := structStr[j]
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_' || (r >= '0' && r <= '9' && j > startName) {
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

		// 如果接下来不是 '('，则把 "<- + 名称" 视为普通文本（写回）并继续
		if j >= n || structStr[j] != '(' {
			// 写回原始 "<-"，然后继续扫描位置在 idx+2 避免死循环
			sb.WriteString("<-")
			i = idx + 2
			continue
		}

		// 现在 j 在 '(' 上，开始平衡括号扫描以找到匹配的 ')'
		open := 0
		k := j
		for k < n {
			c := structStr[k]
			if c == '(' {
				open++
			} else if c == ')' {
				open--
				if open == 0 {
					k++ // k 指向匹配右括号后的下一个位置
					break
				}
			}
			k++
		}
		// 如果括号没有配对完整（k >= n 且 open != 0），视为找不到完整调用，写回并继续
		if open != 0 {
			// 写回 "<-" 并推进
			sb.WriteString("<-")
			i = idx + 2
			continue
		}

		// 完整匹配的调用是 structStr[idx:k]
		fullCall := structStr[idx:k]
		fmt.Println("found arrow call:", fullCall)
		// 判断函数名是否等于目标 funcName
		if funcIdent == funcName {
			// 删除（写入空字符串）
			// 如果你想保留一个空格或换行，可以在此添加（目前选择完全删除）
		} else {
			// 用 @ignore 替换（保留前后原有换行/空格会自然保留）
			sb.WriteString("@ignore")
		}

		// 继续扫描，从 k 开始
		i = k
	}

	return sb.String()
}

// 缩进辅助
func indent(s string, spaces int) string {
	pad := strings.Repeat(" ", spaces)
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		lines[i] = pad + strings.TrimSpace(l)
	}
	return strings.Join(lines, "\n")
}

// 根据调用生成 prompt 文件
func (g *WPDSLGener) generatePromptFile(allin,funcName, inputName string) (PDSLFile, error) {
	filename := filepath.Join(g.OutputDir, fmt.Sprintf("%s.pdsl", funcName))

	inBlock := g.extractIntStructFields(inputName)
	for _, value := range g.PDSLFiles {
		log.Println("inputName:",inputName)
		log.Println("value.callname:",value.callname)
	   if value.callname==inputName{
		
		inBlock=value.out
	   }
	}
	listener := NewArrowStructListener(funcName, g.WorkflowContent)
	antlr.ParseTreeWalkerDefault.Walk(listener, g.tree)
	outBlock := listener.result
	pdsfl:=PDSLFile{
		in: inBlock,
		out: outBlock,
		callname: allin,
	}
	content := fmt.Sprintf(`prompt %s {
    in {
%s
    }
	//内部循环
	//loop[,]
    output {
%s
    }

    before{

    }

    sys {

    }

    user{

    }

	// goimport{

    // }

    fix{

    }

    after{

    }
}
`, funcName, inBlock, outBlock)

	return pdsfl,os.WriteFile(filename, []byte(content), 0644)
}
func (g *WPDSLGener) genPDSLFile() (err error) {
	for _, funcName := range g.order {
		// 找到对应的 match
		for _, call := range g.matches {
			allin := call[0]
			fname := call[1]
			inputName := call[2]

			if allin == funcName {
				log.Println("functioncall:", allin)
				PDSLFile,err := g.generatePromptFile(allin,fname, inputName)
				if err != nil {
					return err
				}
				g.PDSLFiles = append(g.PDSLFiles, PDSLFile)
				fmt.Printf("✅ Generated prompt file: %s\n", fname)
				break
			}
		}
	}
	return nil
}
func (g *WPDSLGener) topsortpdsl() ([]string, error) {
	gD := simple.NewDirectedGraph()
	nodesMap := map[string]int64{}
	var id int64 = 1

	for _, m := range g.matches {
		for _, name := range []string{m[0], m[2]} { // 目标和依赖
			if _, ok := nodesMap[name]; !ok {
				nodesMap[name] = id
				gD.AddNode(simple.Node(id))
				id++
			}
		}
	}

	for _, m := range g.matches {
		from := nodesMap[m[2]] // 依赖
		to := nodesMap[m[0]]   // 目标
		gD.SetEdge(gD.NewEdge(simple.Node(from), simple.Node(to)))
	}

	sortedNodes, err := topo.Sort(gD)
	if err != nil {
		return nil, err // 有环依赖
	}

	idToName := map[int64]string{}
	for k, v := range nodesMap {
		idToName[v] = k
	}
	result := make([]string, len(sortedNodes))
	for i, n := range sortedNodes {
		result[i] = idToName[n.ID()]
	}

	return result, nil
}

func (g *WPDSLGener) Execute() (err error) {

	//提取functioncall和文件内容
	err = g.extractCalls()
	if err != nil {
		log.Println("提取functioncall失败", err)
		return err
	}
	log.Println("g.matchs:", g.matches)
	//top排序（决定生成顺序）
	g.order, err = g.topsortpdsl()
	if err != nil {
		log.Println("拓扑排序失败:", err)
		return err
	}
	log.Println("生成顺序:", g.order)
	//生成结构体列表包含functioncall，instr，outstr，以functioncall为索引，为了嵌套结构的input

	//生成文件
	err = g.genPDSLFile()
	if err != nil {
		log.Println("生成pdsl失败", err)
		return err
	}
	return nil
}
