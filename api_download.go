package mopler

import (
	"bufio"
	"fmt"
	"github.com/766800551/mopler/types"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ApiDownload struct {
	SdkContext *SdkContext
}

func NewApiDownload(sdkContext *SdkContext) *ApiDownload {
	return &ApiDownload{
		SdkContext: sdkContext,
	}
}

// download 下载指定的文件
// downDir 本地存放下载的目录
// basePath 网盘需要下载的目录，直接使用Download时，可以传空
// fsIds 需要下载的文件的id号
// downloadStatuFunc 下载时的回调
func (a *ApiDownload) download(downDir string, basePath string, fsIds []int64, downloadStatuFunc types.DownloadStatuFunc) (err error) {
	_fm, err := a.SdkContext.ApiFileMeta.FileMeta(&types.FileMetaReq{Fsids: fsIds})
	if err != nil {
		return Errcall(err)
	}

	fms := _fm.List
	if len(fms) == 0 {
		return Errcall(err, "没有需要下载的文件，请检查fsId是否正确！")
	}
	_ = os.MkdirAll(downDir, os.ModePerm)
	errs := make(chan error, len(fms))
	for _, f := range fms {
		go func(f types.FileMetaResp) {
			//如果是目录直接创建，不进行下载操作
			if f.Isdir == 1 {
				err = os.MkdirAll(downDir+strings.Replace(f.Path, basePath, "", 1), os.ModePerm)
				if err != nil {
					errs <- err
					return
				}
				errs <- nil
				return
			}
			req, err := http.NewRequest("GET",
				fmt.Sprintf("%v&access_token=%v", f.Dlink, a.SdkContext.Token.AccessToken), nil)
			if err != nil {
				errs <- err
				return
			}
			req.Header.Set("User-Agent", "pan.baidu.com")
			resp, err := a.SdkContext.Client.Do(req)
			if err != nil {
				errs <- err
				return
			}
			defer resp.Body.Close()

			var file *os.File

			if basePath != "" {
				file, err = os.OpenFile(PathSeparatorFormat(downDir+strings.Replace(f.Path, basePath, "", 1)),
					os.O_CREATE|os.O_RDWR,
					os.ModePerm)
			} else {
				file, err = os.OpenFile(PathSeparatorFormat(downDir+string(filepath.Separator)+f.Filename),
					os.O_CREATE|os.O_RDWR,
					os.ModePerm)
			}
			if err != nil {
				errs <- err
				return
			}
			defer file.Close()
			w := bufio.NewWriterSize(file, 1024*1024*4)
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				errs <- err
				return
			}
			_ = w.Flush()
			errs <- nil
			return
		}(f)
	}

	for i := 0; i < len(fms); i++ {
		err = <-errs
		if err != nil {
			return Errcall(err)
		}
		if downloadStatuFunc != nil {
			downloadStatuFunc(fms[i])
		}
	}
	return nil
}

// DownloadByFsId 根据fsid下载文件
// downDir 本地存放下载的目录
// fsIds 需要下载的文件的id号
// downloadStatuFunc 下载时的回调
func (a *ApiDownload) DownloadByFsId(downDir string, fsIds []int64, downloadStatuFunc types.DownloadStatuFunc) (err error) {
	return a.download(downDir, "", fsIds, downloadStatuFunc)
}

// DownloadByPath 下载当前路径的所有文件，不进行递归， 注意！子目录的文件不会下载
// downDir 本地存放下载的目录
// basePath 网盘需要下载的目录，直接使用Download时，可以传空
// downloadStatuFunc 下载时的回调
func (a *ApiDownload) DownloadByPath(downDir string, path string, downloadStatuFunc types.DownloadStatuFunc) (err error) {
	list, err := a.SdkContext.ApiFileList.FileList(&types.FileListReq{Dir: path})
	if err != nil {
		return Errcall(err, "找不到需要下载的文件")
	}
	var fsIds []int64
	for _, f := range list.List {
		fsIds = append(fsIds, f.FsId)
	}
	err = a.download(downDir, path, fsIds, downloadStatuFunc)
	if err != nil {
		return Errcall(err, "文件下载失败")
	}
	return err
}

// DownloadByPathRecursion 递归下载当前路径的所有文件，子目录的也会被下载。
// downDir 本地存放下载的目录
// basePath 网盘需要下载的目录，直接使用Download时，可以传空
// downloadStatuFunc 下载时的回调
func (a *ApiDownload) DownloadByPathRecursion(downDir string, path string, downloadStatuFunc types.DownloadStatuFunc) (err error) {
	list, err := a.SdkContext.ApiFileList.FileListRecursion(&types.FileListRecursionReq{Path: path})
	if err != nil {
		return Errcall(err, "找不到需要下载的文件")
	}
	var fsIds []int64
	for _, f := range list.List {
		fsIds = append(fsIds, f.FsId)
	}
	err = a.download(downDir, path, fsIds, downloadStatuFunc)
	if err != nil {
		return Errcall(err, "文件下载失败")
	}
	return err
}
