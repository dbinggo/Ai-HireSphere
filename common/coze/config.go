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
	Kid       string
	Iss       string
	PKeyPath  string
	DatasetID string
}

type CozeApi struct {
	Doc *CozeDocApi
	Bot *BotApi
}

func NewCozeApi(cfg Config) (api *CozeApi, err error) {
	api = &CozeApi{}
	param := JWTParam{
		Begin: time.Now(),
		Exp:   time.Hour * 3,
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

	go func() {
		for {
			timer := time.NewTimer(time.Hour * 2)
			<-timer.C
			api, err = NewCozeApi(cfg)
			if err != nil {
				return
			}
		}
	}()
	return
}
