package deepseek

// Message 表示对话中的一条消息
type Message struct {
	Role    string `json:"role"`    // 角色：user 或 assistant
	Content string `json:"content"` // 消息内容
}

type StreamResp struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index int `json:"index"`
		Delta struct {
			Content          *string `json:"content"`
			ReasoningContent *string `json:"reasoning_content"`
			Role             string  `json:"role"`
		} `json:"delta"`
		FinishReason         *string `json:"finish_reason"`
		ContentFilterResults struct {
			Hate struct {
				Filtered bool `json:"filtered"`
			} `json:"hate"`
			SelfHarm struct {
				Filtered bool `json:"filtered"`
			} `json:"self_harm"`
			Sexual struct {
				Filtered bool `json:"filtered"`
			} `json:"sexual"`
			Violence struct {
				Filtered bool `json:"filtered"`
			} `json:"violence"`
		} `json:"content_filter_results"`
	} `json:"choices"`
	SystemFingerprint string `json:"system_fingerprint"`
	Usage             struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type NormalResp struct {
	Id                string              `json:"id"`
	Object            string              `json:"object"`
	Created           int                 `json:"created"`
	Model             string              `json:"model"`
	Choices           []NormalRespChoices `json:"choices"`
	Usage             NormalRespUsage     `json:"usage"`
	SystemFingerprint string              `json:"system_fingerprint"`
}
type NormalRespChoices struct {
	Index        int               `json:"index"`
	Message      NormalRespMessage `json:"message"`
	FinishReason string            `json:"finish_reason"`
}
type NormalRespUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
type NormalRespMessage struct {
	Role             string `json:"role"`
	Content          string `json:"content"`
	ReasoningContent string `json:"reasoning_content"`
}
