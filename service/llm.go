package service

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"
	"config"

	openai "github.com/sashabaranov/go-openai"
)

// LLMClient 封装了 OpenAI 客户端
type LLMClient struct {
	client *openai.Client
}

// NewLLMClient 创建 LLMClient 实例，传入 API Key
func NewLLMClient(model config.ModelConfig) *LLMClient {
	cfg := openai.DefaultConfig(model.ApiKey)
	cfg.BaseURL = model.BaseURL
	return &LLMClient{
		client: openai.NewClientWithConfig(cfg),
	}
}

// GeneratePromptResponse 使用 GPT-3.5 Turbo 生成对话回复
func (c *LLMClient) GeneratePromptResponse(systemPrompt, userPrompt string) (string, error) {
	// fmt.Println("GeneratePromptResponse:")
	req := openai.ChatCompletionRequest{
		Model: "qwen-max", // 你这里替换成你具体的 qwen 模型名字符串
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: userPrompt,
			},
		},
	}
	fmt.Println("发送给模型的内容:", req.Messages)
	resp, err := c.client.CreateChatCompletion(context.Background(), req)
	fmt.Println("AI thingking...")
	// fmt.Println("GeneratePromptResponse:",resp)
	if err != nil {
		return "", fmt.Errorf("调用 OpenAI 失败: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("OpenAI 返回空响应")
	}

	content := strings.TrimSpace(resp.Choices[0].Message.Content)
	if content == "" {
		return "", fmt.Errorf("OpenAI 返回空字符串")
	}
	log.Println("OpenAI 输出:", content)
	jsonPart := extractJSONArray(content)

	if jsonPart == "" {
		return "", fmt.Errorf("未能从模型响应中提取 JSON 数组")
	}
	return jsonPart, nil
}

// 提取 JSON 数组
func extractJSONArray(text string) string {
	re := regexp.MustCompile("(?s)```(json|markdown)\\s*(\\{.*?\\}|\\[.*?\\])\\s*```")
	matches := re.FindStringSubmatch(text)
	if len(matches) > 1 {
		return matches[2] // 第一个子匹配是数组
	}
	return ""
}
