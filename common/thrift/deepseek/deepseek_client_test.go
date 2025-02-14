package deepseek

import (
	"io"
	"os"
	"testing"
)

// TestNewDeepSeekClient 测试 NewDeepSeekClient 函数
func TestNewDeepSeekClient(t *testing.T) {
	apiKey := os.Getenv("DEEP_SEEEK_API_KEY")
	client := NewDeepSeekClient(apiKey, SILICONFLOW_CHAT_URL, SILICONFLOW_CHAT_ENDPOINT, SILICONFLOW_DEEPSEEKMODEL_V3, "You are a helpful assistant.")
	go func() {
		for {

			a := 1 + 1
			a = a - 1
			a = a
		}
	}()
	ans, _ := client.NormalChat("hello")
	t.Logf("ans: %v", ans)
	st, _ := client.StreamChat("请重复我上次说的话，测试上下文api调用能力")
	for {
		data, err := st.Read()
		if err != nil {
			t.Logf("err: %v", err)
			if err == io.EOF {
				break
			}
		}
		t.Log(data)
	}

}
