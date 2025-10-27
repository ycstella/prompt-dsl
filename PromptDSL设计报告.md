# PromptDSL 设计报告

## 1. 引言

随着大模型应用的快速普及，如何高效、稳定、可维护地管理 Prompt 已成为实际应用中的关键问题。当前主流的做法往往是将 Prompt 以字符串的形式硬编码在各类调用脚本中（如 Java、C、Python 等）。这种方式虽然直观，但存在明显的局限性。以下从问题与对应的解决思路两个方面展开说明。

### 1.1 传统硬编码方式的问题

1. **冗余的非核心工作**  
   - 目标：用户只想专注于 Prompt 设计，让模型完成任务。  
   - 现实：除了编写 Prompt，还必须额外处理与调用相关的繁琐事务：  
     - 不同厂商 API 的调用方式差异  
     - 客户端初始化与认证  
     - 异常处理与错误恢复  
     - 输入输出的序列化与解析  
   - 结果：这些工作本身与 Prompt 工程无关，却由于编程环境差异而增加负担。  
   - 影响：对于非编程背景的研究者/业务人员而言，学习与处理这些细节会浪费大量时间和精力。

2. **过度强调过程驱动**  
   - 类型：**命令式编程 (Imperative Programming)**  
   - 特点：程序员必须详细描述“如何做 (How to do)”，通过过程语句逐步得到结果。  
   - 举例：在 C/Java 中，用户编写一系列操作步骤，最终结果只是这些过程执行的副产品。  
   - 局限：适合底层算法编程，但不适合 Prompt 工程这种“以最终结果为目标”的任务。

### 1.2 PromptDSL 的设计理念

1. **面向结果的表达方式**  
   - 类型**声明式编程 (Declarative Programming)**  
   - 特点：用户只需描述“要什么 (What to get)”，至于“如何实现”由系统负责。  
   - 类比：  
     - SQL：用户只写查询条件，数据库负责执行优化。  
     - HTML：用户只写页面结构，浏览器负责渲染显示。

2. **借鉴 HTML 的工作流思想**  
   - HTML 的思路：声明页面的最终结构 → 浏览器渲染 → 用户写的就是结果。  
   - PromptDSL 的思路：声明输入、输出和约束条件 → 系统生成调用并执行 → DSL 脚本本身就是结果结构。  

3. **优势**  
   - 屏蔽底层 API 调用与过程逻辑  
   - 提高可维护性与复用性  
   - 让用户专注于 Prompt 工程核心目标  
   - 降低非编程人员的使用门槛  

### 1.3 小结
PromptDSL 的目标是：**让使用者只需描述“想要什么”，而不必关心“如何实现”。**  
这种“声明式”的方式比传统“命令式”更契合 Prompt 工程的核心需求，从而大幅提升效率与可维护性。

## 2. 总体设计目标

- **聚焦核心目标**：让开发人员能够专注于 Prompt 的设计本身，而非底层调用与环境细节。  
- **简化调用链路**：以统一的抽象屏蔽不同厂商 API 与技术栈差异，降低使用门槛。  
- **减少冗余开销**：避免传统硬编码模式下因模型调用，输入输出处理等带来的额外“噪音”代码。  
- **保持结构清晰**：通过声明式的语法风格，使 Prompt 的输入、输出与约束条件在结构上直观呈现。  
- **结果驱动导向**：DSL 脚本即结果描述，开发者所写内容直接对应任务目标，避免过程化干扰。

## 3. 使用者视角（用户层面）

从用户的角度来看，PromptDSL 的核心目标是**降低上手难度**，让用户只需关心 Prompt 的设计逻辑与结果结构，而无需处理底层调用细节。其使用方式主要分为四个层次：配置文件、单任务定义，完整 workflow 文件，vscode插件。

### 3.1 配置文件

配置文件用于定义全局运行环境与默认参数。  
用户通过 **声明式配置**，即可完成日志、工具函数、模型参数、并发等常见设置

**示例：**  

文件：`default.yaml`

```yaml
log:
  level: "debug"
  file: "logs/server.log"

utils:
  - "utils.go"                         # 用户自定义函数文件
  - ...

models:
  - provider: "dashscope"              # 模型服务提供商（如 openai, qwen, gemini 等）
    api_key: "sk-xxxx"                 # API 鉴权密钥
    base_url: "https://dashscope.aliyuncs.com/compatible-mode/v1"
    model: "qwen-max"                  # 模型名称
    temperature: 0.7                   # 随机性参数
    maxConcurrency: 5                  # 并发上限
    stream: false                      # 是否启用流式输出
    retry: 3                           # 请求失败时的重试次数
  - ...
pools:                                 #workflow任务池
    enable: false                      #是否并发                              
    size: 20                           #最大并发数
```
### 3.2 单任务定义文件

以下是一个完整的 **单任务定义模板示例**（文件：`ExtractSteps.pdsl`）。  

```dsl
prompt ExtractSteps {
  in {                        # 输入字段定义
    Id: int
    Question: string
    Solution: string
  }

  output {                    # 输出字段定义
    Id: int @model
    Question: string
    Solution: string
    Steps: []struct {
      Step: int               @hint("步骤编号")
      StepContent: string     @hint("步骤内容")
    } @model
  }

  sys {                       # 系统级 Prompt
    - 你是一位精通小学数学的优秀教师，能够很好地将解题过程进行拆分
  }

  user {                      # 用户提示块
    user_title
    user_content
  }

  user_title {                # 用户提示子模板 1
    - 根据所给题目和解题过程进行逻辑拆分，输出步骤列表
  }

  user_content {              # 用户提示子模板 2，可引用 in.*
    - 以下是题目内容：
      in.Question
    - 以下是解题过程：
      in.Solution
    - 输出严格按照 JSON 格式：
      outputspec
  }

  goimport {                  # 引入外部工具/库
    "github.com/summer-camp-k12/homework-coach"
  }

  fix {                       # 模型输出修正逻辑
  }

  after {                     # 后处理逻辑
  }
}

```
该模板由以下几个核心部分组成：  

- **in**  
  输入数据结构体定义。  

- **output**  
  输出数据结构体定义，支持注解：  

- **sys**  
  系统级 Prompt，相当于 LLM 的 **system**，用于设定角色与行为准则。  

- **user / user_title / user_content**  
  用户级 Prompt：  
  - **user**：整体提示块。  
  - **user_title / user_content**：用于任务描述的子模板，可用于承载输入数据，动态构建提示词

- **fix**  
  模型输出的校验与修正逻辑，用于保证结构和语义符合预期。  

- **after**  
  用户自定义后处理逻辑

#### 各部分内部语法详解

1. **in** —— 输入字段定义

定义任务所需输入字段及类型，供 DSL 引擎注入模板。  

- 支持基本类型、数组和嵌套结构。 
- 可在 `user_content` 中通过 `in.*` 引用。

``` dsl
in {
  Id: int
  Name: string
  Scores: []float
  Details: []struct {
    Age: int
    Address: string
  }
}
```
2. **output** —— 输出字段定义

定义任务期望的输出结构体，并可使用注解指导模型生成结果。

- `@model`：由模型生成的字段。
- `@hint("提示")`：提供字段级语义说明，引导模型更好地生成结果。
- `@jn\jsonname`: JSON 映射注解
- `@outignore`: 模型返回，但最终输出不需要字段
- ...
- 支持嵌套结构、数组和复合类型。

```dsl
output {
  Id: int @model
  Name: string
  Scores: []float @model
  Details: struct {
    Age: int @hint("年龄") 
    Address: string @hint("地址")
  } @model
}
```
3. **sys / user** —— 提示词编写部分

`sys` 和 `user` 分别用于系统级与用户级提示词的编写，语法上完全一致，由 `sys` / `user` 及其子结构构成。子结构可选，非必须。

- 提示词文字每行以 `-` 开头，自动换行结束。
- 支持直接插入参数，例如 `in.*`、迭代变量 `it*` 或自定义参数。
- 支持条件结构：`if / else`。(语法与go相同)
- 支持循环结构：`for`。(语法与go相同)
- 支持分支结构：`switch`。(语法与go相同)
- 使用 `outputspec` 可在该位置插入 `output` 定义的结构体格式，指明模型需要返回的内容。
- 可嵌套使用子模板，实现复杂任务的提示词组合。

```dsl
sys {
  - 你是一位专业数学教师
  - 拆分解题步骤时保持逻辑清晰
}

user {
  user_title
  user_content
}

user_title {
  - 将解题过程拆分为步骤列表
}

user_content {
  - 题目: in.Question
  - 解题过程: in.Solution
  - 输出格式: outputspec
}
```

4. **before / fix / after** —— 数据处理部分

`before`用于处理由输入数据到prompt参数的转化`fix` 用于修复模型返回的 JSON，`after` 用于用户自定义的数据后处理逻辑。两者均使用 Go 代码实现，可对模型输出进行校验、修正和持久化处理。

- **内置支持**：提供内置函数帮助处理常用操作，例如 LaTeX 修复、JSON 解析等。  
- **fix 入参**：模型返回的 JSON 字符串。如果不使用 `fix`，系统会自动尝试将模型返回解析为结构体，但可能解析失败，可通过重试机制处理。若使用 `fix`，需编写 Go 代码处理可能的解析失败情况。  
- **生成结构体名字**：该任务名+结构体名/该任务名+(InputContext,OutputContext)
- **任务对象**：每个任务会被封装为一个对象，包含如下成员：

**示例：**
```go
type GetAnswer struct {
    Input         []GetAnswerInputContext      // 全部输入数据
    Output        GetAnswerOutputContext       // 单次输出数据
    config        string                        // 配置文件
    modelname     string                        // 调用的模型
    inputfile     string                        // 输入数据文件名
    modelRet      string                        // 模型返回 JSON 字符串
    currentData   GetAnswerInputContext         // 循环处理当前数据
    data          []byte                        // 插入参数时的 JSON 转换变量
    fixRet        GetAnswerModelOutputContext  // 修复后的模型输出
    afterRet      GetAnswerOutputContext       // 数据后处理后的结果
    results       []any                         // 最终输出到 JSON 文件的数据
    // 单条数据内部迭代参数
    path          string                        // 迭代路径
    it            OnlStepInfoGainStepAConditons // 最小迭代分支
    it_idx        []int                           // 迭代下标
}
```
- **调用方式**：对象成员可通过 self.* 访问，在 fix 或 after 中进行操作。
**示例**
```dsl
after{
    self.afterRet.Id=self.currentData.Id
    self.afterRet.Question=self.currentData.Question
    self.afterRet.Solution=self.currentData.Solution
    self.afterRet.Process.Id=self.fixRet.Id
    self.afterRet.Process.Steps=self.fixRet.Process.Steps
    return nil
}
```

5. **其他可支持功能**

   1. **looprange[a,b]**  
      - 用于指定循环输入数据中的某个范围。  
      - 支持的用法：  
        - `[a,b]`：表示从第 **a** 条到第 **b** 条数据。  
        - `[a,a]`：表示仅处理第 **a** 条单条数据。  
        - `[a]`：表示从第 **1** 条到第 **a** 条数据。  

   1. **goimport{}**  
      - 在 **fix/after 阶段**引入所需的 Go 包。  
      - 支持别名引入，例如：  
        ```go
        import (
          fmt "fmt"
          json "encoding/json"
        )
        ```  

   1. **it_ctx("Process.Steps")**  
      - 可通过指定路径对内部上下文进行循环访问。  
      - 示例：  
        ```dsl
        it_ctx("Process.Steps")
        ```  
      - 表示迭代访问输入 JSON 中 `Process.Steps` 的内容。  





### 3.3 workflow 文件

`workflow` 文件用于定义完整的任务流，支持将多个 **单任务定义**（如 `ExtractSteps`）以及用户自定义函数组合在一起，实现复杂的数据处理与调用逻辑。  

以下是一个 **workflow 文件模板示例**：  
文件：`workflow.wkfl`
```dsl
workflow {
  in1 {
    // 输入字段定义，可直接作为任务流的全局输入
    Id: string
    Question: string
  }
  []in2 {
    //同in1可自定义，插件选择json，也可路径指向json
  }

  out {
    Id: string                              = in1.Id
    Question: string                        = in1.Question

    // 字段可通过调用单任务、函数以及逻辑组合生成
    stepAly: [                              <- ExtractSteps(in1)    @Retry()    // 调用函数生成字段
      {
        id:                                 

        stepcontext:                        = in1.stepcontext                  // 可直接赋值或由子任务生成

        path: {                                               // 嵌套结构支持
          value:                            <- getPath(ExtractSteps(in1), [$sum(in1.price)*(p2.quantity)])
        }

        num:                                <- [$sum(ExtractSteps().b1) + (p1().a1)]  // 支持运算和链式引用

        field:                              <- [if p1.a1>=100 && p1.b1<50] p2(p1()) [else] p3(p1())    // 支持条件分支调用
                                              
        []field2: {                         <- p4
          // 可继续嵌套结构
        }
      }
    ]
  }

  models {
    default: 0/auto              // 没指定时的默认模型(数字对应yaml中配置的模型)
    getId: 1              
    ExtractSteps: 2                  
    p1: 3               
    p2: 4                      
    p3: 5
    ...                   
  }
}
```
#### 语义说明

##### 输入（in 部分）
- 用于定义输入 JSON 文件的结构体。  
- 用户既可以 **手动定义**，也可以通过 **插件选择文件**，**路径指引文件**，由系统自动解析并渲染。  
- `[]` 表示输入数据是数组：  
  - **无 `[]`** → 单条数据逐条循环引用，直到所有数据完成；  
  - **有 `[]`** → 单次引用整个数据列表。  

##### 输出（out 部分）
- 定义当前工作流的 **最终目标结构**。  
- 可以是单个结构体，也可以是结构体数组。  
- `out` 中的字段可以通过以下方式生成：  
  - **直接赋值**（如 `in1.stepcontext`）  
  - **单任务调用**（如 `p1()`、`ExtractSteps()`）  
  - **函数调用**（如 `getId()`、`getPath()`）  
  - **逻辑组合**（运算、条件分支等） 

##### 模型选择

- 用于 **任务与模型的绑定关系**，不同任务侧重点不同，偏向速度或质量，需要不同的模型。  
- DSL 内部只引用任务名（如 `getId()`、`p1()`），不直接写模型名称。  
- 每个任务名对应一个编号，编号再映射到外部配置文件中的具体模型。  
- **default** 表示默认模型，任务未指定时使用该模型。  

#### 语法规则
- **箭头 `<-`** ：表示赋值，箭尾指向要执行的操作，箭头指向目标字段。  
- **`*()`** ：表示单任务的执行结果；`*().*` 可引用其内部字段。  
- **`[$*]`** ：语法参考 **JSONata**，用于前处理（如聚合、计算）或后处理（与核心 prompt 逻辑无关的结果整理）。  
- **`if ... else`** ：支持条件分支逻辑，可嵌套复杂的判断表达式。 

### 3.4 运行插件

#### 单个任务
1. **新建pdsl文件**
   -新建文件改后缀名为.pdsl，生成文件自动渲染如下结构
   ```dsl
   in{

   }
   out{
    
   }
   before{

   }
   sys{
    
   }
   user{
    
   }
   fix{
    
   }
   after{
    
   }
   ```
1. **输入解析**  
   - 当鼠标焦点位于当前 `.pdsl` 文件时，按下 `Ctrl + Alt + I`  
   - 选择与该 `.pdsl` 文件同级目录下的 `.json` 文件作为输入  
   - 系统会根据 `.json` 文件自动解析结构体，并以 DSL 语法渲染到 `in {}` 中  

2. **运行**  
   - 当鼠标焦点位于当前 `.pdsl` 文件时，按下 `Ctrl + Alt + P`  
   - 选择同级目录下的 `.json` 文件作为输入  
   - 选择同级目录下 `.yaml` 文件中已配置的模型  
   - 系统开始运行，生成临时目录 `generated_code/`，并在 `.pdsl` 文件的父级目录下生成 `log/` 和 `output/` 文件  

3. **Prompt 编写**  
   - 在编写 `prompt` 部分时，系统会在行首自动填充 `- `  

---

#### 工作流（workflow）
1. **输入解析**  
   - 当鼠标焦点位于当前 `.wkfl` 文件时，按下 `Ctrl + Alt + I`  
   - 选择与该 `.wkfl` 文件同级目录下的 `.json` 文件作为输入  
   - 系统会根据 `.json` 文件解析结构体，并以 DSL 语法渲染到 `in {}` 中  

2. **单任务编写**  
   - 在工作流中声明任务，例如：  
     ```dsl
     id: <- getId(in)
     ```
   - 声明完成后，`Ctrl + 单击` 任务名（如 `getId(in)`）  
   - 系统会在 `.wkfl` 文件所在目录下新建一个与任务同名的 `.pdsl` 文件，包含以下部分：  
     - `in`  
     - `out`  
     - `before`  
     - `fix`  
     - `after`  
   - 若任务声明了参数，且能找到依赖项，系统会自动解析依赖结构，并以 DSL 语法渲染到新建 `.pdsl` 文件的 `in {}` 中。
   - 多个依赖项为多个结构体，或者按具体依赖关系进行融合  

3. **运行**  
   - 当鼠标焦点位于当前 `.wkfl` 文件时，按下 `Ctrl + Alt + E`  
   - 选择与该 `.wkfl` 文件同级目录下的 `.json` 文件作为输入  
   - 系统即可运行整个工作流  

#### 外部需求

1. **自定义函数文件**  
   - 当运行 `Ctrl + Alt + p/e` 时，系统会自动复制当前文件的同级 Go 文件，供生成后的代码调用。  
   - 用户可在该文件中编写自定义函数。  
   - 需要在 `yaml` 配置文件中声明文件名，系统才会加载。  

2. **函数参数类型声明**  
   - 在 `.pdsl` 文件尚未运行之前，其内部声明的结构体是不存在的，这会导致编写自定义函数时缺乏类型支持。  
   - 为解决此问题，系统采用 **提前编译** 的方式：  
     - 如果当前焦点在 Go 文件上，系统会检查上一个文件是否为 `.pdsl` 文件；  
     - 若是，则会提前生成 `.pdsl` 中声明的结构体定义，供 Go 文件调用时使用。  









