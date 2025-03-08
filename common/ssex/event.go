package ssex

import (
	"encoding/json"
	"strconv"
)

type SseEventType string

const (
	SseEventTypeId    = "id: "
	SseEventTypeRetry = "retry: "
	SseEventTypeData  = "data: "
	SseEventTypeEvent = "event: "
	sseEnd            = "\n\n"
	sseN              = "\n"
)

var ()

type SseEvent struct {
	Retry int         `json:"retry"`
	Id    string      `json:"id"`
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func (s SseEvent) build() string {
	var ret string
	if s.Retry > 0 {
		ret += SseEventTypeRetry + strconv.Itoa(s.Retry) + sseN
	}
	if s.Id != "" && s.Id != "0" {
		ret += SseEventTypeId + s.Id + sseN
	}
	if s.Event != "" {
		ret += SseEventTypeEvent + s.Event + sseN
	}
	if s.Data != "" {
		b, _ := json.Marshal(s.Data)
		ret += SseEventTypeData + string(b) + sseEnd
	}
	return ret
}
