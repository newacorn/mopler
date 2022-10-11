package mopler

import (
	"encoding/json"
	"fmt"
	"github.com/766800551/mopler/types"
	"io"
	"net/url"
	"os"
	"path/filepath"
)

type ApiFileManager struct {
	SdkContext *SdkContext
}

func NewApiFileManager(sdkContext *SdkContext) *ApiFileManager {
	return &ApiFileManager{
		SdkContext: sdkContext,
	}
}

// CreateDir 创建文件夹
func (a *ApiFileManager) CreateDir(dirname string, remotePath string) (respUps []*types.UploadFileResp, err error) {
	//创建一个临时文件夹，将需要新增的文件夹放到临时文件夹下面，便于后续删除
	temp := MkdirTemp()
	defer temp.DeleteMethod()
	err = os.Mkdir(temp.TempName+string(filepath.Separator)+dirname, os.ModePerm)
	if err != nil {
		return nil, Errcall(err, "文件夹创建失败")
	}

	absPath, err := filepath.Abs(temp.TempName + string(filepath.Separator) + dirname)
	if err != nil {
		return nil, Errcall(err, "获取绝对路径失败")
	}
	respUps, err = a.SdkContext.ApiUpload.Upload([]*types.UploadReq{
		{
			SrcPath:    absPath,
			RemotePath: remotePath,
		},
	}, nil)
	if err != nil {
		return nil, Errcall(err, "文件夹上传失败", respUps)
	}
	return respUps, nil
}

// FileCopy 复制文件
func (a *ApiFileManager) FileCopy(req *types.FileManagerReq) (fmr *types.FileManagerResp, err error) {
	values := url.Values{}
	if req == nil {
		return nil, Errcall(nil, "缺少入参", req)
	}
	if req.Async != "" {
		values.Add("async", req.Async)
	}
	if req.Ondup != "" {
		values.Add("ondup", req.Ondup)
	}
	if len(req.Filelist) > 0 {
		b, err := json.Marshal(req.Filelist)
		if err != nil {
			return nil, Errcall(err, "入参序列化失败")
		}
		values.Add("filelist", string(b))
	}
	l := fmt.Sprintf(PanFileManager, a.SdkContext.Token.AccessToken, "copy")
	resp, err := a.SdkContext.Client.PostForm(l, values)
	if err != nil {
		return nil, Errcall(err, resp)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, Errcall(err, string(b))
	}
	err = json.Unmarshal(b, &fmr)
	if err != nil || fmr.Errno != 0 {
		return nil, Errcall(err, string(b), fmr)
	}
	return fmr, nil
}

// FileRename 重名名
func (a *ApiFileManager) FileRename(req *types.FileManagerReq) (fmr *types.FileManagerResp, err error) {
	values := url.Values{}
	if req == nil {
		return nil, Errcall(nil, "缺少入参", req)
	}
	if req.Async != "" {
		values.Add("async", req.Async)
	}
	if req.Ondup != "" {
		values.Add("ondup", req.Ondup)
	}
	if len(req.Filelist) > 0 {
		b, err := json.Marshal(req.Filelist)
		if err != nil {
			return nil, Errcall(err, "入参序列化失败")
		}
		values.Add("filelist", string(b))
	}
	l := fmt.Sprintf(PanFileManager, a.SdkContext.Token.AccessToken, "rename")
	resp, err := a.SdkContext.Client.PostForm(l, values)
	if err != nil {
		return nil, Errcall(err, resp)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, Errcall(err, string(b))
	}
	err = json.Unmarshal(b, &fmr)
	if err != nil || fmr.Errno != 0 {
		return nil, Errcall(err, string(b), fmr)
	}
	return fmr, nil
}

// FileMove 移动文件
func (a *ApiFileManager) FileMove(req *types.FileManagerReq) (fmr *types.FileManagerResp, err error) {
	values := url.Values{}
	if req == nil {
		return nil, Errcall(nil, "缺少入参", req)
	}
	if req.Async != "" {
		values.Add("async", req.Async)
	}
	if req.Ondup != "" {
		values.Add("ondup", req.Ondup)
	}
	if len(req.Filelist) > 0 {
		b, err := json.Marshal(req.Filelist)
		if err != nil {
			return nil, Errcall(err, "入参序列化失败")
		}
		values.Add("filelist", string(b))
	}
	l := fmt.Sprintf(PanFileManager, a.SdkContext.Token.AccessToken, "move")
	resp, err := a.SdkContext.Client.PostForm(l, values)
	if err != nil {
		return nil, Errcall(err, resp)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, Errcall(err, string(b))
	}
	err = json.Unmarshal(b, &fmr)
	if err != nil || fmr.Errno != 0 {
		return nil, Errcall(err, string(b), fmr)
	}
	return fmr, nil
}

// FileDelete 删除文件
func (a *ApiFileManager) FileDelete(req *types.FileManagerReq) (fmr *types.FileManagerResp, err error) {
	values := url.Values{}
	if req == nil {
		return nil, Errcall(nil, "缺少入参", req)
	}
	if req.Async != "" {
		values.Add("async", req.Async)
	}
	if req.Ondup != "" {
		values.Add("ondup", req.Ondup)
	}
	if len(req.Filelist) > 0 {
		b, err := json.Marshal(req.Filelist)
		if err != nil {
			return nil, Errcall(err, "入参序列化失败")
		}
		values.Add("filelist", string(b))
	}
	l := fmt.Sprintf(PanFileManager, a.SdkContext.Token.AccessToken, "delete")
	resp, err := a.SdkContext.Client.PostForm(l, values)
	if err != nil {
		return nil, Errcall(err, resp)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, Errcall(err, string(b))
	}
	err = json.Unmarshal(b, &fmr)
	if err != nil || fmr.Errno != 0 {
		return nil, Errcall(err, string(b), fmr)
	}
	return fmr, nil
}
