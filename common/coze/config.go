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
	Kid       string `json:",optional"`
	Iss       string `json:",optional"`
	PKeyPath  string `json:",optional"`
	DatasetID string
	BotID     string
	OwnToken  string `json:",optional"`
}

type CozeApi struct {
	Doc *CozeDocApi
	Bot *BotApi
}

func NewCozeApi(cfg Config) (api *CozeApi, err error) {
	fmt.Printf("%+v\n", cfg.OwnToken)
	var token AuthResp
	api = &CozeApi{}

	if cfg.PKeyPath != "" {
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

		token = GetAccessToken(param)
		fmt.Printf("%+v\n", token)

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
	}

	if cfg.OwnToken != "" {
		token.AccessToken = cfg.OwnToken
	}

	api.Doc = NewCozeDocApi(token.AccessToken, cfg.DatasetID)
	api.Bot = NewBotApi(token.AccessToken, cfg.BotID)

	return
}
