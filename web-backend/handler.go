package main

import (
	"codegen" // 引用核心模块
	"encoding/json"
	"io"
	"net/http"
	"service"
	"strings"
)

// 请求结构
type GenGuideRequest struct {
	Input string `json:"input"`
}

// 响应结构
type GenGuideResponse struct {
	Prompt      string `json:"prompt"`
	ModelOutput string `json:"model_output"`
}

// 全局 LLM 客户端（也可以用 DI 方式传入）
var llm *service.LLMClient

// POST /api/genGuide
func HandleGenGuide(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "读取请求体失败", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req GenGuideRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "JSON 解析失败", http.StatusBadRequest)
		return
	}

	// 调用核心模块处理 DSL
	prompt, err := codegen.RunPromptDSL(req.Input, "")
	if err != nil {
		http.Error(w, "生成失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 步骤 2：拆分为 system + user
	promptcontent := *prompt
	systemPart := promptcontent.Sys
	userPart := promptcontent.User
	systemText := strings.Join(systemPart, "\n")
	userText := strings.Join(userPart, "\n")
	// 步骤 3：调用大模型
	result, err := llm.GeneratePromptResponse(systemText, userText)
	if err != nil {
		http.Error(w, "调用大模型失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := GenGuideResponse{
		Prompt:      systemText + userText,
		ModelOutput: result,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
	resp = GenGuideResponse{Prompt: systemText + userText}

}

// 提取 block 内容（例如提取 "system:" 后的内容）
func extractBlock(fullPrompt, blockName string) string {
	start := strings.Index(fullPrompt, blockName)
	if start == -1 {
		return ""
	}
	end := strings.Index(fullPrompt[start+len(blockName):], "\n\n")
	if end == -1 {
		return strings.TrimSpace(fullPrompt[start+len(blockName):])
	}
	return strings.TrimSpace(fullPrompt[start+len(blockName) : start+len(blockName)+end])
}
