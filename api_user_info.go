package mopler

import (
	"encoding/json"
	"fmt"
	"github.com/766800551/mopler/types"
	"io"
)

type ApiUserInfo struct {
	SdkContext *SdkContext
}

func NewApiUserInfo(sc *SdkContext) *ApiUserInfo {
	return &ApiUserInfo{
		SdkContext: sc,
	}
}

// UserInfo 获取用户信息
func (a *ApiUserInfo) UserInfo() (userInfo *types.UserInfoResp, err error) {
	resp, err := a.SdkContext.Client.Get(fmt.Sprintf(PanUinfo, a.SdkContext.Token.AccessToken))
	if err != nil {
		return nil, Errcall(err, resp)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, Errcall(err, string(b))
	}
	err = json.Unmarshal(b, &userInfo)
	if err != nil || userInfo.Errno != 0 {
		return nil, Errcall(err, string(b), userInfo)
	}
	return
}
