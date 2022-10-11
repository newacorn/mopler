package mopler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Code struct {
	DeviceCode      string `json:"device_code"`      //设备码，可用于生成单次凭证 Access Token。
	UserCode        string `json:"user_code"`        //用户码。 如果选择让用户输入 user code 方式，来引导用户授权，设备需要展示 user code 给用户。
	VerificationUrl string `json:"verification_url"` //用户输入 user code 进行授权的 url。
	QrcodeUrl       string `json:"qrcode_url"`       //二维码url，用户用手机等智能终端扫描该二维码完成授权。
	ExpiresIn       int    `json:"expires_in"`       //deviceCode 的过期时间，单位：秒。 到期后 deviceCode 不能换 Access Token。
	Interval        int    `json:"interval"`         //deviceCode 换 Access Token 轮询间隔时间，单位：秒。 轮询次数限制小于 expire_in/interval。
}

type ApiDeviceCodeOAuth struct {
	SdkContext *SdkContext
	CodeResp   *Code
}

// NewApiDeviceCodeOAuth 通过设备码获取授权
func NewApiDeviceCodeOAuth(sc *SdkContext) *ApiDeviceCodeOAuth {
	o := &ApiDeviceCodeOAuth{
		SdkContext: sc,
	}
	// 1. 获取设备码、用户码
	o.CodeResp = o.code()

	// 2.引导用户授权
	o.SdkContext.Config.AuthMethod(o.CodeResp)

	//3.用 Device Code 轮询换取 Access Token
	o.SdkContext.Token = o.pollingToken()

	return o
}

// RefreshToken 刷新token
func RefreshToken(o *ApiDeviceCodeOAuth) error {
	resp, err := http.Get(fmt.Sprintf(PanDeviceRefreshToken,
		o.SdkContext.Token.RefreshToken,
		o.SdkContext.Config.ClientId,
		o.SdkContext.Config.ClientSecret))
	if err != nil {
		ErrHandler(1, "", err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		ErrHandler(1, string(b), err)
	}

	var token Token
	err = json.Unmarshal(b, &token)
	if err != nil {
		ErrHandler(1, string(b), err)
	}
	o.SdkContext.Token.AccessToken = token.AccessToken
	o.SdkContext.Token.ExpiresIn = token.ExpiresIn
	o.SdkContext.Token.RefreshToken = token.RefreshToken
	o.SdkContext.Token.Scope = token.Scope
	return nil
}

func (o *ApiDeviceCodeOAuth) code() *Code {
	resp, err := o.SdkContext.Client.Get(fmt.Sprintf(PanDeviceCode, o.SdkContext.Config.ClientId))
	if err != nil {
		ErrHandler(1, "", err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		ErrHandler(1, string(b), err)
	}
	var rc Code
	err = json.Unmarshal(b, &rc)
	if err != nil || rc == (Code{}) {
		ErrHandler(1, string(b), err)
	}
	return &rc
}

func (o *ApiDeviceCodeOAuth) pollingToken() *Token {
	for {
		resp, err := o.SdkContext.Client.Get(fmt.Sprintf(PanDeviceToken,
			o.CodeResp.DeviceCode,
			o.SdkContext.Config.ClientId,
			o.SdkContext.Config.ClientSecret))
		if err != nil {
			ErrHandler(1, "", err)
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			ErrHandler(1, string(b), err)
		}
		//如果获取到了access_token，那么就退出
		if strings.Contains(string(b), "access_token") {
			var token Token
			err = json.Unmarshal(b, &token)
			if err != nil {
				ErrHandler(1, string(b), err)
			}
			resp.Body.Close()
			return &token
		}
		resp.Body.Close()
		//轮询接口必须5秒以上，这里6秒一次
		time.Sleep(time.Second * 6)
	}
}
