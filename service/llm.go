package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/along416/promptDSL/config"
	openai "github.com/sashabaranov/go-openai"
)

// LLMClient 封装了 OpenAI 客户端
type LLMClient struct {
	client *openai.Client
	model  config.ModelConfig
}

// NewLLMClient 创建 LLMClient 实例，传入 API Key
func NewLLMClient(model config.ModelConfig) *LLMClient {
	cfg := openai.DefaultConfig(model.ApiKey)
	cfg.BaseURL = model.BaseURL
	proxy:=model.Proxy
	// 如果有 proxy 设置，则配置 HTTPClient
	if proxy != "" {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			log.Fatalf("❌ 无效的代理 URL: %v", err)
		}
		cfg.HTTPClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		}
	}

	return &LLMClient{
		client:   openai.NewClientWithConfig(cfg),
		model:    model,
	}
}

// GeneratePromptResponse 使用 GPT-3.5 Turbo 生成对话回复
func (c *LLMClient) GeneratePromptResponse(systemPrompt, userPrompt string) (string, error) {
	// fmt.Println("GeneratePromptResponse:")
	// Gemini 特殊处理
	if strings.HasPrefix(strings.ToLower(c.model.Model), "gemini") {
		return c.generateGeminiResponse(systemPrompt, userPrompt)
	}
	req := openai.ChatCompletionRequest{
		Model: c.model.Model,
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


// generateGeminiResponse 专门处理 Gemini 调用
func (c *LLMClient) generateGeminiResponse(systemPrompt, userPrompt string) (string, error) {
	key := c.model.ApiKey
	if key == "" {
		return "", fmt.Errorf("❌ Gemini API Key 未设置")
	}

	// 使用代理
	httpClient := &http.Client{}
	if c.model.Proxy != "" {
		proxyURL, _ := url.Parse(c.model.Proxy)
		httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			},
		}
	}

	cfg := openai.DefaultConfig(key)
	cfg.BaseURL = c.model.BaseURL
	cfg.HTTPClient = httpClient
	client := openai.NewClientWithConfig(cfg)

	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: systemPrompt},
		{Role: openai.ChatMessageRoleUser, Content: userPrompt},
	}

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:    c.model.Model,
		Messages: messages,
	})
	if err != nil {
		return "", fmt.Errorf("❌ Gemini API 请求失败: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("Gemini 返回空响应")
	}

	content := strings.TrimSpace(resp.Choices[0].Message.Content)
	log.Println("Gemini 输出:", content)
	jsonPart := extractJSONArray(content)
	if jsonPart == "" {
		return "", fmt.Errorf("未能从 Gemini 响应中提取 JSON 数组")
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
