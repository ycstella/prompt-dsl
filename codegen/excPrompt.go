package codegen

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"github.com/ycstella/prompt-dsl/config"
	"github.com/ycstella/prompt-dsl/service"

)

// 通用的基类，所有步骤都会用到
type BaseSteps struct {
	Config    string
	Modelname string
	Inputfile string
	Data      []byte
	ModelRet  string
}

// 初始化配置和日志
func (b *BaseSteps) Init() {
	if len(os.Args) < 3 {
		log.Fatal("请提供输入文件路径和配置路径作为参数")
	}
	log.Println("[main] 程序启动，等待输入...")
	config.InitConfig(b.Config)
	config.InitLogger()
}

// 读取输入文件
func (b *BaseSteps) LoadInput() {
	var err error
	b.Data, err = os.ReadFile(b.Inputfile)
	if err != nil {
		log.Fatalf("无法读取文件: %v", err)
	}
}

// 解析输入数据，支持单个或多个
func ParseData[T any](b *BaseSteps, out *[]T) {
	if err := json.Unmarshal(b.Data, out); err == nil {
		return
	}
	var single T
	if err := json.Unmarshal(b.Data, &single); err == nil {
		*out = []T{single}
		return
	}
	log.Fatalf("输入 JSON 格式无效")
}


// 泛型接口，确保每个 Steps 都有 BaseSteps
type Task interface {
	Before() error
	AfterProcess() error
    FixProcess() error
    GenUser() string
    GenSys() string
	ValidateInput() error
    Base() *BaseSteps
	SetBaseSteps(base BaseSteps)
}

// 泛型构造函数
func NewTask[T Task]() *T {
	base := BaseSteps{
		Config:    os.Args[2],
		Modelname: os.Args[1],
		Inputfile: os.Args[3],
	}
	var t T
	t.SetBaseSteps(base)
	return &t
}


func RunTask(t Task) error {
    var err error
	err = t.ValidateInput()
	if err != nil {
		return err
	}
	err = t.Before()
	if err != nil {
		log.Println(os.Stderr, "预处理失败 %v", err)
		os.Exit(1)
	}
	sys := t.GenSys(t.currentData)
	user := t.GenUser(t.currentData)
	modelConfig := config.GetModelConfig(p.modelname)
	llm := service.NewLLMClient(*modelConfig)

	if modelConfig.Stream {
		// 流式模式，直接输出，不做 JSON 解析
		_, err := llm.GeneratePromptResponse(sys, user, true)
		if err != nil {
			log.Println("调用大模型失败:", err)
			os.Exit(1)
		}
		return err
	} else {
		// 非流式模式，拿完整结果再处理
		t.modelRet, err = llm.GeneratePromptResponse(sys, user, false)
		if err != nil {
			log.Println("调用大模型失败: ", err)
			os.Exit(1)
		}
		t.FixProcess()
		if err != nil {
			log.Println("解析输入 JSON 失败011111:", err)
			os.Exit(1)
		}
		err = t.AfterProcess()
		if err != nil {
			log.Println("数据后处理失败", err)
		}
		// log.Println("p.afterRet", p.afterRet)
		encoded, err := json.Marshal(t.afterRet)
		if err != nil {
			log.Println("输出编码失败:", err)
			os.Exit(1)
		}
		log.Println(string(encoded))
		return err
	}
}













// ---- 具体实现 ----

// 定义 SplitSolutionSteps
type SplitSolutionSteps struct {
	BaseSteps
	Input       []string
	Output      string
	ModelOutput []string
}

func (s *SplitSolutionSteps) SetBaseSteps(base BaseSteps) {
	s.BaseSteps = base
}



// 定义 OtherSteps
type OtherSteps struct {
	BaseSteps
	Question []string
	Result   string
}

func (o *OtherSteps) SetBaseSteps(base BaseSteps) {
	o.BaseSteps = base
}

// ---- 使用 ----
func main() {
	// 模拟命令行参数
	os.Args = []string{"prog", "gpt-4", "config.yaml", "input.json"}

	split := NewTask[SplitSolutionSteps]()
	other := NewTask[OtherSteps]()

	fmt.Printf("Split: %+v\n", split)
	fmt.Printf("Other: %+v\n", other)
}
