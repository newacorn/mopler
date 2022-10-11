package mopler

import (
	"encoding/json"
	"fmt"
	"github.com/766800551/mopler/types"
	"io"
	"net/url"
)

type ApiFileList struct {
	SdkContext *SdkContext
}

func NewApiFileList(sdkContext *SdkContext) *ApiFileList {
	return &ApiFileList{
		SdkContext: sdkContext,
	}
}

func (a *ApiFileList) FileList(req *types.FileListReq) (filelist *types.FileListResp, err error) {
	l := fmt.Sprintf(PanFileList, a.SdkContext.Token.AccessToken)
	if req != nil {
		if req.Folder != "" {
			l += fmt.Sprintf("&folder=%v", req.Folder)
		}
		if req.Order != "" {
			l += fmt.Sprintf("&order=%v", req.Order)
		}
		if req.Showempty != "" {
			l += fmt.Sprintf("&showempty=%v", req.Showempty)
		}
		if req.Desc != "" {
			l += fmt.Sprintf("&desc=%v", req.Desc)
		}
		if req.Dir != "" {
			l += fmt.Sprintf("&dir=%v", url.QueryEscape(req.Dir))
		}
		if req.Limit != "" {
			l += fmt.Sprintf("&limit=%v", req.Limit)
		}
		if req.Start != "" {
			l += fmt.Sprintf("&start=%v", req.Start)
		}
		if req.Web != "" {
			l += fmt.Sprintf("&web=%v", req.Web)
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
	err = json.Unmarshal(b, &filelist)
	if err != nil || filelist.Errno != 0 {
		return nil, Errcall(err, string(b), filelist)
	}
	return filelist, nil
}

func (a *ApiFileList) FileListRecursion(req *types.FileListRecursionReq) (filelist *types.FileListResp, err error) {
	l := fmt.Sprintf(PanFileListRecursion, a.SdkContext.Token.AccessToken)
	if req != nil {
		if req.Order != "" {
			l += fmt.Sprintf("&order=%v", req.Order)
		}
		if req.Desc != "" {
			l += fmt.Sprintf("&desc=%v", req.Desc)
		}
		if req.Path != "" {
			l += fmt.Sprintf("&dir=%v", url.QueryEscape(req.Path))
		}
		if req.Limit != "" {
			l += fmt.Sprintf("&limit=%v", req.Limit)
		}
		if req.Start != "" {
			l += fmt.Sprintf("&start=%v", req.Start)
		}
		if req.Web != "" {
			l += fmt.Sprintf("&web=%v", req.Web)
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
	err = json.Unmarshal(b, &filelist)
	if err != nil || filelist.Errno != 0 {
		return nil, Errcall(err, string(b), filelist)
	}
	return filelist, nil
}
