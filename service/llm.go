package service

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/along416/promptDSL/config"
	openai "github.com/sashabaranov/go-openai"
)

// LLMClientInterface 统一接口
type LLMClientInterface interface {
	GeneratePromptResponse(systemPrompt, userPrompt string, stream bool) (string, error)
}

// LLMClient 封装大模型客户端
type LLMClient struct {
	model       config.ModelConfig
	openaiCli   *openai.Client
	deepseekURL string
	httpClient  *http.Client
}

// NewLLMClient 创建客户端
func NewLLMClient(model config.ModelConfig) *LLMClient {
	httpClient := &http.Client{Timeout: 60 * time.Second}

	if model.Proxy != "" {
		proxyURL, err := url.Parse(model.Proxy)
		if err != nil {
			log.Fatalf("❌ 无效的代理 URL: %v", err)
		}
		httpClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}

	client := &LLMClient{
		model:      model,
		httpClient: httpClient,
	}

	// OpenAI/Gemini 客户端
	if !strings.HasPrefix(strings.ToLower(model.Model), "deepseek") {
		cfg := openai.DefaultConfig(model.ApiKey)
		cfg.BaseURL = model.BaseURL
		cfg.HTTPClient = httpClient
		client.openaiCli = openai.NewClientWithConfig(cfg)
	} else {
		client.deepseekURL = model.BaseURL // DeepSeek API URL
	}

	return client
}

// GeneratePromptResponse 统一生成函数
func (c *LLMClient) GeneratePromptResponse(systemPrompt, userPrompt string, stream bool) (string, error) {
	if strings.HasPrefix(strings.ToLower(c.model.Model), "deepseek") {
		return c.generateDeepSeek(systemPrompt, userPrompt, stream)
	}
	return c.generateOpenAI(systemPrompt, userPrompt, stream)
}

// ---- OpenAI / Gemini 调用 ----
func (c *LLMClient) generateOpenAI(systemPrompt, userPrompt string, stream bool) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: c.model.Model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: systemPrompt},
			{Role: openai.ChatMessageRoleUser, Content: userPrompt},
		},
		Temperature: float32(c.model.Temperature),
	}

	if stream {
		streamer, err := c.openaiCli.CreateChatCompletionStream(context.Background(), req)
		if err != nil {
			return "", fmt.Errorf("OpenAI stream 请求失败: %w", err)
		}
		defer streamer.Close()

		var builder strings.Builder
		for {
			resp, err := streamer.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return "", fmt.Errorf("OpenAI stream 接收失败: %w", err)
			}
			builder.WriteString(resp.Choices[0].Delta.Content)
		}
		return builder.String(), nil
	}

	resp, err := c.openaiCli.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("OpenAI 请求失败: %w", err)
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("OpenAI 返回空响应")
	}
	return strings.TrimSpace(resp.Choices[0].Message.Content), nil
}

// ---- DeepSeek 调用 ----
func (c *LLMClient) generateDeepSeek(systemPrompt, userPrompt string, stream bool) (string, error) {
	// 构建请求体
	reqBody := map[string]interface{}{
		"model": c.model.Model,
		"messages": []map[string]string{
			{"role": "system", "content": systemPrompt},
			{"role": "user", "content": userPrompt},
		},
		"temperature": c.model.Temperature,
		"stream":      stream,
	}

	bodyBytes, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", c.model.BaseURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("创建 DeepSeek 请求失败: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.model.ApiKey)

	// fmt.Println("请求 URL:", c.deepseekURL+"/chat/completions")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("DeepSeek 请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		data, _ := io.ReadAll(resp.Body)
		fmt.Println("HTTP 状态码:", resp.StatusCode)
		fmt.Println("返回内容:", string(data))
		return "", fmt.Errorf("DeepSeek 返回错误: %s", string(data))
	}

	if stream {
		// 流式处理
		reader := bufio.NewReader(resp.Body)
		var builder strings.Builder
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				return "", fmt.Errorf("DeepSeek stream 读取失败: %w", err)
			}

			if !bytes.HasPrefix(line, []byte("data: ")) {
				continue
			}

			chunk := bytes.TrimPrefix(line, []byte("data: "))
			if bytes.Equal(chunk, []byte("[DONE]\n")) {
				break
			}

			var parsed struct {
				Choices []struct {
					Delta struct {
						Content string `json:"content,omitempty"`
					} `json:"delta"`
				} `json:"choices"`
			}

			if err := json.Unmarshal(chunk, &parsed); err != nil {
				continue
			}
			if len(parsed.Choices) > 0 {
				builder.WriteString(parsed.Choices[0].Delta.Content)
			}
		}

		return builder.String(), nil
	}

	// 非流式处理
	data, _ := io.ReadAll(resp.Body)
	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return "", fmt.Errorf("DeepSeek 解析 JSON 失败: %w", err)
	}

	if len(result.Choices) > 0 {
		return result.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("DeepSeek 返回内容为空")
}
