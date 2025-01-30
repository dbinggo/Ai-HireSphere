package deepseek

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestNewDeepSeekClient 测试 NewDeepSeekClient 函数
func TestNewDeepSeekClient(t *testing.T) {
	apiKey := "test-api-key"
	baseURL := "http://test-url.com"
	client := NewDeepSeekClient(apiKey, baseURL)
	if client.apiKey != apiKey {
		t.Errorf("Expected apiKey to be %s, got %s", apiKey, client.apiKey)
	}
	if client.baseURL != baseURL {
		t.Errorf("Expected baseURL to be %s, got %s", baseURL, client.baseURL)
	}
	if len(client.context) != 0 {
		t.Errorf("Expected context to be empty, got %v", client.context)
	}
	if client.clientPool == nil {
		t.Errorf("Expected clientPool to be initialized, got nil")
	}
}

// TestAddToContext 测试 AddToContext 方法
func TestAddToContext(t *testing.T) {
	client := NewDeepSeekClient("test-api-key", "http://test-url.com")
	client.AddToContext("user", "Hello")
	if len(client.context) != 1 {
		t.Errorf("Expected context length to be 1, got %d", len(client.context))
	}
	if client.context[0].Role != "user" || client.context[0].Content != "Hello" {
		t.Errorf("Expected context to contain {Role: user, Content: Hello}, got %v", client.context[0])
	}
}

// TestClearContext 测试 ClearContext 方法
func TestClearContext(t *testing.T) {
	client := NewDeepSeekClient("test-api-key", "http://test-url.com")
	client.AddToContext("user", "Hello")
	client.ClearContext()
	if len(client.context) != 0 {
		t.Errorf("Expected context to be empty, got %v", client.context)
	}
}

// TestCallAPI 测试 CallAPI 方法
func TestCallAPI(t *testing.T) {
	// 使用 httptest 包来模拟 HTTP 服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.Header.Get("Authorization") != "Bearer test-api-key" {
			t.Errorf("Expected Authorization header to be Bearer test-api-key, got %s", r.Header.Get("Authorization"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type header to be application/json, got %s", r.Header.Get("Content-Type"))
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"choices":[{"message":{"content":"Hello, World!"}}]}`)
	}))
	defer server.Close()

	client := NewDeepSeekClient("test-api-key", server.URL)
	data := map[string]interface{}{
		"model":    "deepseek-chat",
		"messages": []Message{{Role: "user", Content: "Hello"}},
		"stream":   false,
	}
	response, err := client.CallAPI("chat/completions", data, false)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	result := response.(map[string]interface{})
	assistantMessage := result["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	if assistantMessage != "Hello, World!" {
		t.Errorf("Expected assistant message to be Hello, World!, got %s", assistantMessage)
	}
}

// TestChat 测试 Chat 方法
func TestChat(t *testing.T) {
	// 使用 httptest 包来模拟 HTTP 服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.Header.Get("Authorization") != "Bearer test-api-key" {
			t.Errorf("Expected Authorization header to be Bearer test-api-key, got %s", r.Header.Get("Authorization"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type header to be application/json, got %s", r.Header.Get("Content-Type"))
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"choices":[{"message":{"content":"Hello, World!"}}]}`)
	}))
	defer server.Close()

	client := NewDeepSeekClient("test-api-key", server.URL)
	message, err := client.Chat("Hello", false)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if message != "Hello, World!" {
		t.Errorf("Expected message to be Hello, World!, got %s", message)
	}
	if len(client.context) != 2 {
		t.Errorf("Expected context length to be 2, got %d", len(client.context))
	}
	if client.context[0].Role != "user" || client.context[0].Content != "Hello" {
		t.Errorf("Expected context to contain {Role: user, Content: Hello}, got %v", client.context[0])
	}
	if client.context[1].Role != "assistant" || client.context[1].Content != "Hello, World!" {
		t.Errorf("Expected context to contain {Role: assistant, Content: Hello, World!}, got %v", client.context[1])
	}
}
