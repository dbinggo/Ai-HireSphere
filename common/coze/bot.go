package coze

import (
	"Ai-HireSphere/common/ssex"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	CreateSessionAPI = "https://api.coze.cn/v1/conversation/create"
	ChatAPIFormat    = "https://api.coze.cn/v3/chat?conversation_id=%d"
	botID            = "7463508586472423465"
)

type BotApi struct {
	Header BotHeader
}

type BotHeader struct {
	Authorization string `json:"Authorization"`
	ContentType   string `json:"Content-Type"`
}

type BotSessionData struct {
	APIResp
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

type BotChatBody struct {
	BotID              string       `json:"bot_id"`
	UserID             string       `json:"user_id"`
	Stream             bool         `json:"stream"`
	AdditionalMessages []BotMessage `json:"additional_messages"`
}

type BotMessage struct {
	Role        string `json:"role"`
	ContentType string `json:"content_type"`
	Content     string `json:"content"`
}

func NewBotApi(token string) *BotApi {
	return &BotApi{
		Header: BotHeader{
			Authorization: "Bearer " + token,
			ContentType:   "application/json",
		},
	}
}

// 创建会话
func (bot *BotApi) CreateSession() int64 {
	request, err := http.NewRequest("POST", CreateSessionAPI, nil)
	if err != nil {
		return 0
	}

	headerMap := make(map[string]string)
	var headerByte []byte

	headerByte, err = json.Marshal(bot.Header)
	if err != nil {
		return 0
	}

	err = json.Unmarshal(headerByte, &headerMap)
	if err != nil {
		return 0
	}

	for key, value := range headerMap {
		request.Header.Set(key, value)
	}

	client := http.Client{}
	resp, err := client.Do(request)
	defer resp.Body.Close()

	if err != nil {
		return 0
	}

	var apiData BotSessionData
	var dataByte []byte
	dataByte, err = io.ReadAll(resp.Body)
	if err != nil {
		return 0
	}
	err = json.Unmarshal(dataByte, &apiData)
	if err != nil {
		return 0
	}
	if apiData.Code != 0 {
		return 0
	}

	//data, ok := apiData.Data.(BotSessionData)
	fmt.Printf("create session: %+v", apiData.Data.ID)
	id, err := strconv.Atoi(apiData.Data.ID)
	if err != nil {
		return 0
	}

	return int64(id)
}

// 流式对话
func (bot *BotApi) Chat(sessionID int64, message string) (ch chan string, err error) {
	c, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	c = context.Background()
	var req *http.Request
	body := new(bytes.Buffer)
	bodyData := BotChatBody{
		BotID:              botID,
		UserID:             "1",
		Stream:             true,
		AdditionalMessages: []BotMessage{{Role: "user", ContentType: "text", Content: message}},
	}
	err = json.NewEncoder(body).Encode(bodyData)
	if err != nil {
		return
	}

	req, err = http.NewRequestWithContext(c, "POST", fmt.Sprintf(ChatAPIFormat, sessionID), body)
	if err != nil {
		return
	}

	headerMap := make(map[string]string)
	var headerByte []byte

	headerByte, err = json.Marshal(bot.Header)
	if err != nil {
		return
	}

	err = json.Unmarshal(headerByte, &headerMap)
	if err != nil {
		return
	}

	for key, value := range headerMap {
		req.Header.Set(key, value)
	}

	//connection := ssex.Client.NewConnection(req)
	//connection.SubscribeToAll(func(msg sse.Event) {
	//	logrus.Debugf("sse message: %v", msg)
	//	ch <- msg.Data
	//})
	//go func() {
	//	err = connection.Connect()
	//	if !errors.Is(err, context.Canceled) {
	//		// 结束了会返回 EOF
	//		logrus.Debugf("sse finish: %v", err)
	//	}
	//	logrus.Debugf("sse finish: %v", err)
	//	close(ch)
	//}()

	ch, err = ssex.Connect(req)
	if err != nil {
		return
	}

	return
}
