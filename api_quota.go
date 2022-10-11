package mopler

import (
	"encoding/json"
	"fmt"
	"github.com/766800551/mopler/types"
	"io"
)

type ApiQuota struct {
	SdkContext *SdkContext
}

func NewApiQuota(sdkContext *SdkContext) *ApiQuota {
	return &ApiQuota{
		SdkContext: sdkContext,
	}
}

func (a *ApiQuota) Quota(req *types.QuotaReq) (quota *types.QuotaResp, err error) {
	l := fmt.Sprintf(PanQuota, a.SdkContext.Token.AccessToken)
	if req != nil {
		if req.Checkfree != "" {
			l += fmt.Sprintf("&checkfree=%v", req.Checkfree)
		}
		if req.Checkexpire != "" {
			l += fmt.Sprintf("&checkexpire=%v", req.Checkexpire)
		}
	}

	resp, err := a.SdkContext.Client.Get(l)
	if err != nil {
		return nil, Errcall(err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, Errcall(err, string(b))
	}
	err = json.Unmarshal(b, &quota)
	if err != nil || quota.Errno != 0 {
		return nil, Errcall(err, string(b), quota)
	}
	return quota, nil
}
