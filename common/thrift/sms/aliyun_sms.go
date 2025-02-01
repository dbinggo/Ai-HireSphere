package sms

import (
	"Ai-HireSphere/common/zlog"
	"context"
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

const (
	TemplateCaptchaCode = "SMS_296340473"
)

// AliyunSMSClient 阿里云SMS客户端结构体
type AliyunSMSClient struct {
	client *dysmsapi.Client
}

type AliyunSMSConfig struct {
	AccessKeyId     string
	AccessKeySecret string
}

// NewAliyunSMSClient 初始化阿里云SMS客户端
func NewAliyunSMSClient(accessKeyId, accessKeySecret string) (*AliyunSMSClient, error) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}

	c := &AliyunSMSClient{
		client: client,
	}

	return c, nil
}

func MustNewAliyunSMSClient(accessKeyId, accessKeySecret string) *AliyunSMSClient {
	client, err := NewAliyunSMSClient(accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	return client
}

func (c *AliyunSMSClient) SendCaptcha(ctx context.Context, phoneNumber, code string) error {

	bizId, err := c.SendSMS(phoneNumber, "竞赛云界小程序", TemplateCaptchaCode, map[string]string{"code": code})
	zlog.InfofCtx(ctx, "send sms to %s, bizId: %s", phoneNumber, bizId)
	return err
}

// SendSMS 发送短信
func (c *AliyunSMSClient) SendSMS(phoneNumber, signName, templateCode string, templateParam map[string]string) (string, error) {

	param, _ := json.Marshal(templateParam)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phoneNumber
	request.SignName = signName
	request.TemplateCode = templateCode
	request.TemplateParam = string(param)

	response, err := c.client.SendSms(request)
	if err != nil {
		return "", err
	}

	return response.BizId, nil
}
