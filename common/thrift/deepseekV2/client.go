package deepseekV2

import (
	"net/http"
)

const (
	baseURL = "https://api.deepseek.com/chat/completions"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
	baseUrl    string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
		baseUrl:    baseURL,
	}
}
