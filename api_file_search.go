package mopler

import (
	"encoding/json"
	"fmt"
	"github.com/766800551/mopler/types"
	"io"
	"net/url"
)

type ApiFileSearch struct {
	SdkContext *SdkContext
}

func NewAPiFileSearch(sdk *SdkContext) *ApiFileSearch {
	return &ApiFileSearch{
		SdkContext: sdk,
	}
}

func (a *ApiFileSearch) Search(req *types.FileSearchReq) (fsr *types.FileSearchResp, err error) {
	l := fmt.Sprintf(PanFileSearch, a.SdkContext.Token.AccessToken, url.QueryEscape(req.Key))
	if req != nil {
		if req.Web != "" {
			l += fmt.Sprintf("&web=%v", req.Web)
		}
		if req.Dir != "" {
			l += fmt.Sprintf("&dir=%v", url.QueryEscape(req.Dir))
		}
		if req.Num != "" {
			l += fmt.Sprintf("&num=%v", req.Num)
		}
		if req.Page != "" {
			l += fmt.Sprintf("&page=%v", req.Page)
		}
		if req.Recursion != "" {
			l += fmt.Sprintf("&recursion=%v", req.Recursion)
		}
	}

	resp, err := a.SdkContext.Client.Get(l)
	if err != nil {
		return nil, Errcall(err, resp)
	}

	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, Errcall(err, string(b))
	}

	err = json.Unmarshal(b, &fsr)
	if err != nil || fsr.Errno != 0 {
		return nil, Errcall(err, string(b), fsr)
	}
	return fsr, nil
}
