package test

import (
	"Ai-HireSphere/common/coze"
	"Ai-HireSphere/utils"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	param := coze.JWTParam{
		Kid:   "xgs534BFxEVpf91pMPo-Tk8o4Ov4hR5GPX9a7jKPknU",
		Iss:   "1166857449405",
		Begin: time.Now(),
		Exp:   time.Minute * 5,
	}

	privKeyBytes, err := os.ReadFile("../private/private_key.pem")
	if err != nil {
		fmt.Println(err)
		return
	}
	block, _ := pem.Decode(privKeyBytes)
	if block == nil {
		fmt.Println(123)
		return
	}

	pKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}

	param.PKey = pKey.(*rsa.PrivateKey)
	token := coze.GetAccessToken(param)
	api := coze.NewCozeDocApi(token.AccessToken, "7454197247615008820")
	var files []utils.FileBase
	files = append(files, utils.FileBase{
		FileByte: []byte("fbfvbndhjakcfvbdshjcndsjkbncvdsjkbfghasfrs"),
		Filename: "test01",
		Ext:      "pdf",
	})
	err = api.CreateDoc(files)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
}
