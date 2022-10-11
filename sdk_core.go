package mopler

import (
	"net/http"
)

type Token struct {
	AccessToken  string `json:"access_token"`  //获取到的Access Token，Access Token是调用网盘开放API访问用户授权资源的凭证。
	ExpiresIn    int    `json:"expires_in"`    //Access Token的有效期，单位为秒。
	RefreshToken string `json:"refresh_token"` //用于刷新 Access Token, 有效期为10年。
	Scope        string `json:"scope"`         //Access Token 最终的访问权限，即用户的实际授权列表。
}

type SdkContext struct {
	Config         Configuration
	Client         *http.Client
	OAuth          *ApiDeviceCodeOAuth
	Token          *Token
	ApiUserInfo    *ApiUserInfo
	ApiQuota       *ApiQuota
	ApiFileList    *ApiFileList
	ApiUpload      *ApiUpload
	APiFileSearch  *ApiFileSearch
	ApiFileMeta    *ApiFileMeta
	ApiDownload    *ApiDownload
	ApiFileManager *ApiFileManager
}

func NewSdkContext(config Configuration) *SdkContext {
	sc := &SdkContext{
		Client: &http.Client{},
		Token: &Token{
			AccessToken: config.AccessToken,
		},
		Config: config,
	}

	//如果token为空，那么就进行鉴权，如果不为空，那么就直接使用
	if *sc.Token == (Token{}) {
		//鉴权
		switch config.TokenMode {
		case AuthorizationCode:
			//	todo
		case ImplicitGrant:
			//	todo
		case DeviceCode:
			sc.OAuth = NewApiDeviceCodeOAuth(sc)
		default:
			sc.OAuth = NewApiDeviceCodeOAuth(sc)
		}
	}

	//初始化api
	sc.ApiUserInfo = NewApiUserInfo(sc)
	sc.ApiQuota = NewApiQuota(sc)
	sc.ApiFileList = NewApiFileList(sc)
	sc.ApiUpload = NewApiUpload(sc)
	sc.APiFileSearch = NewAPiFileSearch(sc)
	sc.ApiFileMeta = NewApiFileMeta(sc)
	sc.ApiDownload = NewApiDownload(sc)
	sc.ApiFileManager = NewApiFileManager(sc)
	return sc
}
