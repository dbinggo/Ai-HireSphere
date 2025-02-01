package sms

import (
	"Ai-HireSphere/common/zlog"
	"context"
	"testing"
)

// TestSendCaptcha 测试SendCaptcha方法
func TestSendCaptcha(t *testing.T) {

	zlog.Develop()
	// 创建一个模拟的AliyunSMSClient
	mockClient, err := NewAliyunSMSClient("xxxx", "xxxx")
	if err != nil {
		t.Fatalf("Failed to create AliyunSMSClient: %v", err)
	}

	// 调用SendCaptcha方法
	err = mockClient.SendCaptcha(context.Background(), "13131227873", "123456")
	if err != nil {
		t.Errorf("SendCaptcha returned an error: %v", err)
	}
}
