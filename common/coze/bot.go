package coze

import (
	"Ai-HireSphere/common/ssex"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	CreateSessionAPI     = "https://api.coze.cn/v1/conversation/create"
	ChatAPIFormat        = "https://api.coze.cn/v3/chat?conversation_id=%d"
	ChatHistortAPIFormat = " https://api.coze.cn/v1/conversation/message/list?conversation_id=%d"

	ErrTimeout = `{"code":500, "msg":"timeout"}`
)

type BotApi struct {
	Header BotHeader
	BotID  string
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
type BotHistoryResp struct {
	APIResp
	Data []BotMessage `json:"data"`
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

type BotStreamReply struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

func NewBotApi(token string, botID string) *BotApi {
	return &BotApi{
		Header: BotHeader{
			Authorization: "Bearer " + token,
			ContentType:   "application/json",
		},
		BotID: botID,
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
	if err != nil {
		return 0
	}

	defer resp.Body.Close()

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
func (bot *BotApi) Chat(sessionID int64, message string) (ch chan BotStreamReply, err error) {
	c, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	c = context.Background()
	var req *http.Request
	body := new(bytes.Buffer)
	bodyData := BotChatBody{
		BotID:              bot.BotID,
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
	var strCh chan string
	strCh, err = ssex.Connect(req)
	if err != nil {
		return
	}
	ch = make(chan BotStreamReply)
	go func() {
		defer close(ch)
		var reply BotStreamReply
		for reply.Data == "" ||
			!(strings.Contains(reply.Event, "done") || strings.Contains(reply.Event, "conversation.chat.failed")) {
			timer := time.NewTimer(time.Minute * 5)
			select {
			case msg := <-strCh:
				if strings.HasPrefix(msg, "event") {
					reply.Event = strings.TrimPrefix(msg, "event:")
				}
				if strings.HasPrefix(msg, "data") {
					reply.Data = strings.TrimPrefix(msg, "data:")
					ch <- reply
					//清空数据
					reply = BotStreamReply{}
				}
			case <-timer.C:
				reply.Event = "conversation.chat.failed"
				reply.Data = ErrTimeout
				ch <- reply
				return
			}
		}
	}()

	return
}

// 获取会话历史记录
func (bot *BotApi) GetChatHistory(sessionID int64) (history []BotMessage, err error) {
	request, err := http.NewRequest("POST", fmt.Sprintf(ChatHistortAPIFormat, sessionID), nil)
	if err != nil {
		return nil, err
	}

	headerMap := make(map[string]string)
	var headerByte []byte

	headerByte, err = json.Marshal(bot.Header)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(headerByte, &headerMap)
	if err != nil {
		return nil, err
	}

	for key, value := range headerMap {
		request.Header.Set(key, value)
	}

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiData BotHistoryResp
	var dataByte []byte
	dataByte, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataByte, &apiData)
	if err != nil {
		return nil, err
	}
	if apiData.Code != 0 {
		return nil, errors.New(apiData.Msg)
	}

	return apiData.Data, nil
}
