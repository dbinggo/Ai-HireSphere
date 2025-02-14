package deepseekV2

import "os"

var (
	DeepSeekClient *Client
)

func Init() {
	DeepSeekClient = NewClient(os.Getenv("DEEPSEEK_API_KEY"))
}
