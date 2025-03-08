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
	"strings"
	"time"
)

const (
	WorkFlowApI       = "https://api.coze.cn/v1/workflow/run"
	WorkFlowStreamApi = "https://api.coze.cn/v1/workflow/stream_run"
)

type workFlowReq struct {
	WorkflowId string                 `json:"workflow_id"`
	Parameters map[string]interface{} `json:"parameters"`
	IsASync    bool                   `json:"is_async"`
}

type workFlowAsyncResp struct {
	Code      int    `json:"code"`
	DebugUrl  string `json:"debug_url"`
	ExecuteId string `json:"execute_id"`
	Msg       string `json:"msg"`
}

type getWorkFlowASyncResultReq struct {
	WorkFlowId string `json:"workflow_id"`
	ExecuteId  string `json:"execute_id"`
}

type getWorkFlowASyncResultResp struct {
	Code int                              `json:"code"`
	Msg  string                           `json:"msg"`
	Data []getWorkFlowASyncResultRespData `json:"data"`
}

type getWorkFlowASyncResultRespData struct {
	UpdateTime    int    `json:"update_time"`
	Cost          string `json:"cost"`
	Output        string `json:"output"`
	BotId         string `json:"bot_id"`
	Token         string `json:"token"`
	ExecuteStatus string `json:"execute_status"`
	ConnectorUid  string `json:"connector_uid"`
	RunMode       int    `json:"run_mode"`
	ConnectorId   string `json:"connector_id"`
	Logid         string `json:"logid"`
	DebugUrl      string `json:"debug_url"`
	ErrorCode     string `json:"error_code"`
	ErrorMessage  string `json:"error_message"`
	ExecuteId     string `json:"execute_id"`
	CreateTime    int    `json:"create_time"`
}

type WorkFlowDoBody struct {
	WorkFlowID string                 `json:"workflow_id"`
	Parameters map[string]interface{} `json:"parameters"`
}

type WorkFlowStreamResp struct {
	ID    int    `json:"id"`
	Event string `json:"event"`
	Data  string `json:"data"`
}

// 异步执行工作流
func (bot *BotApi) UseWorkFlowASync(workflowId string, parameters map[string]interface{}) (string, error) {
	var req = workFlowReq{
		WorkflowId: workflowId,
		Parameters: parameters,
		IsASync:    true,
	}

	var resp workFlowAsyncResp
	err := bot.newRequest("POST", WorkFlowApI, req, &resp)
	if err != nil {
		return "", err
	}
	if resp.Code != 0 {
		return "", errors.New(resp.Msg)
	}
	return resp.ExecuteId, nil
}

// 获得异步执行工作流结果
func (bot *BotApi) GetASyncResult() (string, string, error) {
	var req = getWorkFlowASyncResultReq{
		WorkFlowId: "workflow_id",
		ExecuteId:  "execute_id",
	}
	var resp getWorkFlowASyncResultResp
	err := bot.newRequest("POST", WorkFlowApI, req, &resp)
	if err != nil {
		return "", "", err
	}
	if resp.Code != 0 {
		return "", "", errors.New(resp.Msg)
	}
	return resp.Data[0].ExecuteStatus, resp.Data[0].Output, nil
}

// 快捷请求
func (bot *BotApi) newRequest(method string, url string, req interface{}, respRet interface{}) error {
	headerMap := make(map[string]string)
	var headerByte []byte
	reqBody, _ := json.Marshal(req)
	request, err := http.NewRequest(method, url, strings.NewReader(string(reqBody)))
	if err != nil {
		return err
	}

	headerByte, err = json.Marshal(bot.Header)
	if err != nil {
		return err
	}

	err = json.Unmarshal(headerByte, &headerMap)
	if err != nil {
		return err
	}

	for key, value := range headerMap {
		request.Header.Set(key, value)
	}

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var dataByte []byte
	dataByte, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dataByte, &respRet)
	return err
}

func (bot *BotApi) StreamWorkFlow(workFlowID string, parameters map[string]interface{}) (chan WorkFlowStreamResp, error) {
	c, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	c = context.Background()
	var req *http.Request
	body := new(bytes.Buffer)
	bodyData := WorkFlowDoBody{
		WorkFlowID: workFlowID,
		Parameters: parameters,
	}
	err := json.NewEncoder(body).Encode(bodyData)
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequestWithContext(c, "POST", WorkFlowStreamApi, body)
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
		req.Header.Set(key, value)
	}

	var strCh chan string
	strCh, err = ssex.Connect(req)
	if err != nil {
		return nil, err
	}
	ch := make(chan WorkFlowStreamResp)
	go func() {
		defer func() {
			close(ch)
		}()

		var reply WorkFlowStreamResp
		for reply.Data == "" || (!(strings.Contains(reply.Event, "done") &&
			strings.Contains(reply.Event, "conversation.chat.failed"))) {
			timer := time.NewTimer(time.Minute * 5)
			select {
			case msg, ok := <-strCh:
				if !ok {
					return
				}
				reply = WorkFlowStreamResp{}
				if strings.HasPrefix(msg, "id") {
					_, err = fmt.Sscanf(msg, "id: %d", &reply.ID)
					if err != nil {
						reply.Event = "Error"
						reply.Data = ErrTimeout
						ch <- reply
						return
					}
				}
				if strings.HasPrefix(msg, "event") {
					reply.Event = strings.TrimPrefix(msg, "event:")
				}
				if strings.HasPrefix(msg, "data") {
					reply.Data = strings.TrimPrefix(msg, "data:")
				}
				ch <- reply

			case <-timer.C:
				reply.Event = "Error"
				reply.Data = ErrTimeout
				ch <- reply
				return
			}
		}
	}()

	return ch, nil
}
