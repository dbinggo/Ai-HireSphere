package coze

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Kid       string `json:"kid"`
	Iss       string `json:"iss"`
	Exp       int    `json:"exp"` //单位小时
	PKeyPath  string `json:"pkey_path"`
	DatasetID string `json:"dataset_id"`
}

type CozeApi struct {
	Doc *CozeDocApi
	Bot *BotApi
}

func NewCozeApi(cfg Config) (api *CozeApi, err error) {
	api = &CozeApi{}
	param := JWTParam{
		Begin: time.Now(),
		Exp:   time.Hour * time.Duration(cfg.Exp),
		Iss:   cfg.Iss,
		Kid:   cfg.Kid,
	}

	var privKeyBytes []byte
	privKeyBytes, err = os.ReadFile(cfg.PKeyPath)
	if err != nil {
		return
	}

	var block *pem.Block
	block, _ = pem.Decode(privKeyBytes)

	var pKey interface{}
	pKey, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	param.PKey = pKey.(*rsa.PrivateKey)

	token := GetAccessToken(param)
	fmt.Printf("%+v\n", token)
	api.Doc = NewCozeDocApi(token.AccessToken, cfg.DatasetID)
	api.Bot = NewBotApi(token.AccessToken)
	return
}
