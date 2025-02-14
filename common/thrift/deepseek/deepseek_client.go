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

const (
	SILICONFLOW_CHAT_URL         = "https://api.siliconflow.cn"
	SILICONFLOW_CHAT_ENDPOINT    = "v1/chat/completions"
	SILICONFLOW_DEEPSEEKMODEL_R1 = "deepseek-ai/DeepSeek-R1" // 硅基流动版本 deepseek r1 模型
	SILICONFLOW_DEEPSEEKMODEL_V3 = "deepseek-ai/DeepSeek-V3" // 硅基流动版本 deepseek v3 模型

)

// DeepSeekClient 是 DeepSeek API 的客户端
type DeepSeekClient struct {
	apiKey   string
	baseURL  string
	endPoint string
	model    string
	context  []Message
	client   *http.Client
	mu       sync.Mutex // 为了维护上下文的整洁 我们需要单个client只支持串行访问
}

// NewDeepSeekClient 创建一个新的 DeepSeek 客户端 同时配置好prompt
func NewDeepSeekClient(apiKey, baseURL string, endpoint string, model string, prompt string) *DeepSeekClient {
	clent := &DeepSeekClient{
		apiKey:   apiKey,
		baseURL:  baseURL,
		endPoint: endpoint,
		context:  []Message{},
		model:    model,
		client:   http.DefaultClient,
	}
	clent.addToContext("system", prompt)
	return clent
}

// AddToContext 添加消息到上下文
func (c *DeepSeekClient) addToContext(role, content string) {
	c.mu.Lock()
	c.context = append(c.context, Message{Role: role, Content: content})
	c.mu.Unlock()
}

// ClearContext 清空上下文
func (c *DeepSeekClient) ClearContext() {
	c.context = []Message{}
}

// CallAPI 调用 DeepSeek API
func (c *DeepSeekClient) callApi(endpoint string, data map[string]interface{}, stream bool) (interface{}, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, endpoint)

	// 将请求数据编码为 JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request data: %v", err)
	}

	// 获取客户端实例

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// 设置请求头
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := c.client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 处理流式响应
	if stream {
		return c.newStream(resp.Body), nil
	} else {
		// 处理普通响应
		var chatResp NormalResp
		if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
			return nil, fmt.Errorf("failed to decode response: %v", err)
		}
		return chatResp, nil
	}

}

// NormalChat 与 DeepSeek 直接对话
func (c *DeepSeekClient) NormalChat(message string) (*NormalResp, error) {

	// 添加用户消息到上下文
	c.addToContext("user", message)

	// 构造请求数据
	data := map[string]interface{}{
		"model":    c.model,
		"messages": c.context,
		"stream":   false,
	}
	// 调用 API
	response, err := c.callApi(c.endPoint, data, false)
	if err != nil {
		return nil, fmt.Errorf("API call failed: %v", err)
	}
	// 处理普通响应
	result := response.(NormalResp)

	return &result, nil
}

func (c *DeepSeekClient) JsonChat(comment string, formatStruct interface{}) (*NormalResp, error) {
	// 添加用户消息到上下文
	c.addToContext("user", comment)

	// 构造请求数据
	data := map[string]interface{}{
		"model":    c.model,
		"messages": c.context,
		"stream":   false,
		//"response_format": formatStruct,
		"temperature": 0.3,
	}
	// 调用 API
	response, err := c.callApi(c.endPoint, data, false)
	if err != nil {
		return nil, fmt.Errorf("API call failed: %v", err)
	}
	// 处理普通响应
	result := response.(NormalResp)

	return &result, nil
}

// StreamChat 与 DeepSeek 流式对话
func (c *DeepSeekClient) StreamChat(message string) (*Stream, error) {

	// 添加用户消息到上下文
	c.addToContext("user", message)

	// 构造请求数据
	data := map[string]interface{}{
		"model":    c.model,
		"messages": c.context,
		"stream":   true,
	}

	// 调用 API
	response, err := c.callApi(c.endPoint, data, true)

	if err != nil {
		return nil, fmt.Errorf("API call failed: %v", err)
	}

	ch := response.(*Stream)
	return ch, nil
}

type Stream struct {
	stream chan StreamResp
	closed chan struct{}
	body   io.ReadCloser
}

// Recv 接收流式响应
func (s *Stream) Recv() <-chan StreamResp {
	return s.stream
}
func (s *Stream) Read() (resp StreamResp, err error) {
	select {
	case <-s.closed:
		return resp, io.EOF
	case resp = <-s.stream:
		return resp, nil
	}
}

// Close 关闭流
func (s *Stream) Close() error {
	close(s.closed)
	return s.body.Close()
}

func (c *DeepSeekClient) newStream(body io.ReadCloser) *Stream {
	s := &Stream{
		stream: make(chan StreamResp, 10),
		closed: make(chan struct{}),
		body:   body,
	}
	go s.handleStreamResponse()
	return s
}

func (s *Stream) handleStreamResponse() {
	scanner := bufio.NewScanner(s.body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "data: ") && line != "data: [DONE]" || line == "" {
			data := strings.TrimPrefix(line, "data: ")
			var resp StreamResp
			err := json.Unmarshal([]byte(data), &resp)
			if err != nil || resp.Choices[0].FinishReason != nil {
				s.Close()
				return
			}
			s.stream <- resp
		} else if line == "data: [DONE]" {
			s.Close()
			return
		}
	}
}
