package mopler

import (
	"encoding/json"
	"fmt"
	"github.com/766800551/mopler/types"
	"io"
	"strings"
)

type ApiFileMeta struct {
	SdkContext *SdkContext
}

func NewApiFileMeta(sdk *SdkContext) *ApiFileMeta {
	return &ApiFileMeta{
		SdkContext: sdk,
	}
}

func (a *ApiFileMeta) FileMeta(req *types.FileMetaReq) (fm *types.FileMetasResp, err error) {
	l := fmt.Sprintf(PanFileMeta, a.SdkContext.Token.AccessToken, "1",
		fmt.Sprintf("[%v]", strings.Join(Int64SliceToStringSlice(req.Fsids), ",")))
	if req != nil {
		if req.Path != "" {
			l += fmt.Sprintf("&path=%v", req.Path)
		}
		if req.Extra != "" {
			l += fmt.Sprintf("&extra=%v", req.Extra)
		}
		if req.Thumb != "" {
			l += fmt.Sprintf("&thumb=%v", req.Thumb)
		}
		if req.Needmedia != "" {
			l += fmt.Sprintf("&needmedia=%v", req.Needmedia)
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
	err = json.Unmarshal(b, &fm)
	if err != nil || fm.Errno != 0 {
		return nil, Errcall(err, string(b), fm)
	}
	return fm, nil
}
