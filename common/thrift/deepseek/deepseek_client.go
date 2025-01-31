package deepseek

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
)

const ()

// DeepSeekClient 是 DeepSeek API 的客户端
type DeepSeekClient struct {
	apiKey     string
	baseURL    string
	context    []Message
	client     *http.Client
	clientPool *sync.Pool // 添加对象池
}

// Message 表示对话中的一条消息
type Message struct {
	Role    string `json:"role"`    // 角色：user 或 assistant
	Content string `json:"content"` // 消息内容
}

// NewDeepSeekClient 创建一个新的 DeepSeek 客户端
func NewDeepSeekClient(apiKey, baseURL string) *DeepSeekClient {
	return &DeepSeekClient{
		apiKey:  apiKey,
		baseURL: baseURL,
		context: []Message{},
		clientPool: &sync.Pool{ // 初始化对象池
			New: func() interface{} {
				return &http.Client{}
			},
		},
	}
}

// AddToContext 添加消息到上下文
func (c *DeepSeekClient) AddToContext(role, content string) {
	c.context = append(c.context, Message{Role: role, Content: content})
}

// ClearContext 清空上下文
func (c *DeepSeekClient) ClearContext() {
	c.context = []Message{}
}

// CallAPI 调用 DeepSeek API
func (c *DeepSeekClient) CallAPI(endpoint string, data map[string]interface{}, stream bool) (interface{}, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, endpoint)

	// 将请求数据编码为 JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request data: %v", err)
	}

	// 获取客户端实例
	client := c.clientPool.Get().(*http.Client)
	defer c.clientPool.Put(client) // 使用后放回对象池

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 处理流式响应
	if stream {
		return c.handleStreamResponse(resp.Body), nil
	}

	// 处理普通响应
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}

// handleStreamResponse 处理流式响应
func (c *DeepSeekClient) handleStreamResponse(body io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		scanner := bufio.NewScanner(body)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "data: ") {
				ch <- strings.TrimPrefix(line, "data: ")
			}
		}
	}()
	return ch
}

// Chat 与 DeepSeek 进行对话
func (c *DeepSeekClient) Chat(message string, stream bool) (string, error) {
	// 添加用户消息到上下文
	c.AddToContext("user", message)

	// 构造请求数据
	data := map[string]interface{}{
		"model":    "deepseek-chat",
		"messages": c.context,
		"stream":   stream,
	}

	// 调用 API
	response, err := c.CallAPI("chat/completions", data, stream)
	if err != nil {
		return "", fmt.Errorf("API call failed: %v", err)
	}

	// 处理流式响应
	if stream {
		ch := response.(<-chan string)
		var fullResponse strings.Builder
		for chunk := range ch {
			fullResponse.WriteString(chunk)
			fmt.Print(chunk) // 实时打印流式响应
		}
		assistantMessage := fullResponse.String()
		c.AddToContext("assistant", assistantMessage)
		return assistantMessage, nil
	}

	// 处理普通响应
	result := response.(map[string]interface{})
	assistantMessage := result["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	c.AddToContext("assistant", assistantMessage)
	return assistantMessage, nil
}
