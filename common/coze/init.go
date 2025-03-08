package coze

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stringx"
	"io"
	"net/http"
	"strings"
	"time"
)

// 用于获取通用字段值,例如token

const (
	defaultAlg       = "RS256"
	defaultTyp       = "JWT"
	defaultAud       = "api.coze.cn"
	defaultGrantType = "urn:ietf:params:oauth:grant-type:jwt-bearer"
)

type JWTHeader struct {
	Alg string `json:"alg"` // 固定为RS256
	Typ string `json:"typ"` // 固定为 JWT
	Kid string `json:"kid"` // OAuth 应用的公钥指纹
}

type JWTPayload struct {
	Iss string `json:"iss"`
	Aud string `json:"aud"`
	Iat int64  `json:"iat"` //Unixtime 时间戳格式，精确到秒
	Exp int64  `json:"exp"`
	Jti string `json:"jti"`
}

type JWTSignature struct {
	Kid   string          `json:"kid"`
	Iss   string          `json:"iss"`
	Begin time.Time       `json:"begin"`
	End   time.Time       `json:"end"`
	PKey  *rsa.PrivateKey `json:"PKey"`
}

type JWTParam struct {
	Kid   string          `json:"kid"`
	Iss   string          `json:"iss"`
	Begin time.Time       `json:"begin"`
	Exp   time.Duration   `json:"exp"`
	PKey  *rsa.PrivateKey `json:"PKey"`
}

type AuthResp struct {
	AccessToken string `json:"access_token"`
}

type APIResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 开始签名
func GetJWT(param JWTSignature) string {
	header := getJWTHeader(param.Kid)
	payload := getJWTPayload(param.Iss, param.Begin, param.End)
	signature := getSignature(param.PKey, header, payload)
	jwt := header + "." + payload + "." + signature
	return jwt
}

// 签署header
func getJWTHeader(kid string) string {
	header := JWTHeader{
		Alg: defaultAlg,
		Typ: defaultTyp,
		Kid: kid,
	}

	headerByte, _ := json.Marshal(header)
	b64 := base64UrlEncode(headerByte)
	return b64
}

// 签署payload
func getJWTPayload(iss string, begin, end time.Time) string {
	payload := JWTPayload{
		Iss: iss,
		Aud: defaultAud,
		Iat: begin.Unix(),
		Exp: end.Unix(),
	}

	jti := stringx.Randn(32)
	payload.Jti = jti
	payloadByte, _ := json.Marshal(payload)
	b64 := base64UrlEncode(payloadByte)
	return b64
}

// 生成签名
func getSignature(privateKey *rsa.PrivateKey, header, payload string) string {
	message := header + "." + payload

	// 对拼接后的字符串计算SHA256哈希
	hash := sha256.New()
	hash.Write([]byte(message))
	hashed := hash.Sum(nil)

	// 使用私钥签名（RS256）
	signature, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return ""
	}

	// 返回签名的Base64Url编码
	return base64UrlEncode(signature)
}

func base64UrlEncode(data []byte) string {
	encoded := base64.StdEncoding.EncodeToString(data)
	encoded = strings.TrimRight(encoded, "=") // 去掉`=`
	encoded = strings.ReplaceAll(encoded, "+", "-")
	encoded = strings.ReplaceAll(encoded, "/", "_")
	return encoded
}

func GetAccessToken(param JWTParam) AuthResp {
	type Header struct {
		ContentType   string `json:"Content-Type"`
		Authorization string `json:"Authorization"`
	}

	type Data struct {
		Seconds   int64  `json:"duration_seconds"`
		GrantType string `json:"grant_type"`
	}

	sign := JWTSignature{
		Kid:   param.Kid,
		Iss:   param.Iss,
		Begin: param.Begin,
		End:   param.Begin.Add(param.Exp),
		PKey:  param.PKey,
	}

	header := Header{
		ContentType:   "application/json",
		Authorization: fmt.Sprintf("Bearer %s", GetJWT(sign)),
	}

	fmt.Println(header)

	data := Data{
		GrantType: defaultGrantType,
		Seconds:   int64(param.Exp.Seconds()),
	}

	fmt.Println(data)

	url := "https://api.coze.cn/api/permission/oauth2/token"

	headerByte, _ := json.Marshal(header)
	headerMap := make(map[string]string)
	json.Unmarshal(headerByte, &headerMap)
	dataByte, _ := json.Marshal(data)

	var client http.Client
	request, err := http.NewRequest("POST", url, strings.NewReader(string(dataByte)))
	if err != nil {
		return AuthResp{}
	}
	for key, value := range headerMap {
		request.Header.Set(key, value)
	}

	resp, err := client.Do(request)
	defer resp.Body.Close()

	respByte, err := io.ReadAll(resp.Body)
	//fmt.Println(string(respByte))

	var respData AuthResp
	err = json.Unmarshal(respByte, &respData)
	if err != nil {
		return AuthResp{}
	}

	return respData
}
