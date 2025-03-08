package coze

import (
	"fmt"
)

type Config struct {
	Kid       string `json:",optional"`
	Iss       string `json:",optional"`
	PKeyPath  string `json:",optional"`
	DatasetID string
	OwnToken  string `json:",optional"`
}

type CozeApi struct {
	token string `json:"token"`
}

func NewCozeApi(cfg Config) (api CozeApi, err error) {
	fmt.Printf("%+v\n", cfg.OwnToken)
	api = CozeApi{}

	//if cfg.PKeyPath != "" {
	//	param := JWTParam{
	//		Begin: time.Now(),
	//		Exp:   time.Hour * 3,
	//		Iss:   cfg.Iss,
	//		Kid:   cfg.Kid,
	//	}
	//	var privKeyBytes []byte
	//	privKeyBytes, err = os.ReadFile(cfg.PKeyPath)
	//	if err != nil {
	//		return
	//	}
	//
	//	var block *pem.Block
	//	block, _ = pem.Decode(privKeyBytes)
	//
	//	var pKey interface{}
	//	pKey, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	//	if err != nil {
	//		return
	//	}
	//
	//	param.PKey = pKey.(*rsa.PrivateKey)
	//
	//	token = GetAccessToken(param)
	//	fmt.Printf("%+v\n", token)
	//
	//	go func() {
	//		for {
	//			timer := time.NewTimer(time.Hour * 2)
	//			<-timer.C
	//			api, err = NewCozeApi(cfg)
	//			if err != nil {
	//				return
	//			}
	//		}
	//	}()
	//}

	api.token = cfg.OwnToken
	return api, nil
}
func (c *CozeApi) GetToken() string {
	return c.token
}
