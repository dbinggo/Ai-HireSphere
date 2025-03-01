package coze

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

const (
	WorkFlowApI = "https://api.coze.cn/v1/workflow/run"
)

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
