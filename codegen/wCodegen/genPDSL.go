package wcodegen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

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
	PDSLFiles       []PDSLFile
	ignoreRanges []Range
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
func (g *WPDSLGener) generatePromptFile(allin, funcName, inputName string) (PDSLFile, error) {
	filename := filepath.Join(g.OutputDir, fmt.Sprintf("%s.pdsl", funcName))

	inBlock := g.extractIntStructFields(inputName)
	for _, value := range g.PDSLFiles {
		log.Println("inputName:", inputName)
		log.Println("value.callname:", value.callname)
		if value.callname == inputName {

			inBlock = value.out
		}
	}
	// listener := NewArrowStructListener(funcName, g.WorkflowContent)
	// antlr.ParseTreeWalkerDefault.Walk(listener, g.tree)

	outBlock := extractTargetContent(g.WorkflowContent, funcName, g.ignoreRanges)
	pdsfl := PDSLFile{
		in:       inBlock,
		out:      outBlock,
		callname: allin,
	}
	content := fmt.Sprintf(`prompt %s {
    in {
%s
    }

	//外部循环
	//loop[,]
	
	//内部循环
	//it_ctx("path")

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

	return pdsfl, os.WriteFile(filename, []byte(content), 0644)
}
func (g *WPDSLGener) genPDSLFile() (err error) {
	// fmt.Println("genPDSLFile-----------")
	listener := NewArrowStructListener("a", g.WorkflowContent)
	antlr.ParseTreeWalkerDefault.Walk(listener, g.tree)
	g.ignoreRanges=listener.ignoreRanges
	for _, funcName := range g.order {
		// 找到对应的 match
		for _, call := range g.matches {
			allin := call[0]
			fname := call[1]
			inputName := call[2]

			if allin == funcName {
				log.Println("functioncall:", allin)
				PDSLFile, err := g.generatePromptFile(allin, fname, inputName)
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
		from := nodesMap[m[2]] 
		to := nodesMap[m[0]]   
		gD.SetEdge(gD.NewEdge(simple.Node(from), simple.Node(to)))
	}

	sortedNodes, err := topo.Sort(gD)
	if err != nil {
		return nil, err 
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
