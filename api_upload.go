package mopler

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/766800551/mopler/types"
	"io"
	"io/fs"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type ApiUpload struct {
	SdkContext *SdkContext
}

func NewApiUpload(sdkContext *SdkContext) *ApiUpload {
	return &ApiUpload{
		SdkContext: sdkContext,
	}
}

var (
	oldRemotePath string
	newRemotePath string
)

func (a *ApiUpload) Upload(reqs []*types.UploadReq, superfileStatuFunc types.SuperfileStatuFunc) (respUps []*types.UploadFileResp, err error) {
	respUps = make([]*types.UploadFileResp, 0)
	for _, req := range reqs {
		t := MkdirTemp()
		req.SrcPath, err = filepath.Abs(req.SrcPath)
		if err != nil {
			return nil, Errcall(err, "获取绝对路径失败")
		}
		err = filepath.Walk(req.SrcPath, func(path string, info fs.FileInfo, err error) error {
			path = PathSeparatorFormat(path)
			req.SrcPath = PathSeparatorFormat(req.SrcPath)
			req.RemotePath = PathSeparatorFormat(req.RemotePath)

			//保留用户传入的远程和本地的根路径，防止在递归搜索的时候将其改动了。
			_remote := req.RemotePath
			_path := req.SrcPath

			//查看远程根目录下面的文件
			list, err := a.SdkContext.ApiFileList.FileList(&types.FileListReq{Dir: _remote})
			if err != nil {
				log.Println("查看远程根目录下面的文件失败", err)
			}

			//设置搜索到的路径，路径为远程根路径+本地去掉根路径
			//例如远程路径为：/apps/novel ，本地根路径为c:/desktop/  本地文件路径为：c:/desktop/filefolder，
			//那么新的远程路径就是/apps/novel/filefolder

			req.RemotePath = _remote +
				strings.Replace(path, PathSeparatorFormat(filepath.Dir(req.SrcPath)), "", 1)
			//本地的路径变更为walk递归的路径，确保准确找到文件
			req.SrcPath = path

			//判断当前文件是否是文件夹
			if info.IsDir() {
				//迭代远程根目录下的所有文件的路径
				for _, v := range list.List {

					//如果远程文件的路径有和设置的远程路径一致时，则说明文件夹已经存在，那么就将冲突的文件夹进行临时存储，
					//并且根据时间戳创建一个新的文件夹
					if v.Path == req.RemotePath {
						oldRemotePath = req.RemotePath
						newRemotePath = req.RemotePath + fmt.Sprintf("_%v", time.Now().Unix())
					}
				}
			}

			//将设置的远程冲突的文件夹进行替换为新的文件夹
			req.RemotePath = strings.Replace(req.RemotePath, oldRemotePath, newRemotePath, -1)

			//上传文件
			{
				//如果是目录，就直接创建
				if info.IsDir() {
					req.Isdir = "1"
					up, err := a.create(req, nil, nil)
					if err != nil {
						return Errcall(err, up)
					}

					respUps = append(respUps, up)
				} else {
					//文件分片
					bl, err := SplitFile(path, t.TempName, 1024*1024*4)
					if err != nil {
						return Errcall(err, bl)
					}
					//预上传
					upr, err := a.precreate(req, bl)
					if err != nil {
						return Errcall(err, upr)
					}
					//分片上传
					err = a.superfile(req, bl, upr, superfileStatuFunc)
					if err != nil {
						return Errcall(err, "分片上传失败")
					}
					//创建文件
					up, err := a.create(req, bl, upr)
					if err != nil {
						return Errcall(err, up)
					}
					respUps = append(respUps, up)
				}
			}

			//恢复用户的根路径
			req.RemotePath = _remote
			req.SrcPath = _path
			return err
		})
		if err != nil {
			ErrHandler(1, "路径搜索失败", err)
		}

		t.DeleteMethod()
	}
	return
}

// 预上传
func (a *ApiUpload) precreate(req *types.UploadReq, bl *types.BlockList) (upr *types.UploadPrecreateResp, err error) {
	values := url.Values{}
	if req != nil {
		info, err := os.Stat(req.SrcPath)
		if err != nil {
			return nil, Errcall(err, info)
		}
		req.Size = info.Size()
		if !info.IsDir() {
			req.Isdir = "0"
		} else {
			req.Isdir = "1"
			return nil, err
		}
		values.Add("path", url.QueryEscape(req.RemotePath))
		values.Add("size", fmt.Sprintf("%v", req.Size))
		values.Add("isdir", req.Isdir)
		values.Add("block_list", bl.Block)
		values.Add("autoinit", "1")
		if req.Rtype != "" {
			values.Add("rtype", req.Rtype)
		}
	}
	resp, err := a.SdkContext.Client.PostForm(fmt.Sprintf(PanUploadPrecreate,
		a.SdkContext.Token.AccessToken), values)
	if err != nil {
		return nil, Errcall(err, resp)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, Errcall(err, string(b))
	}
	err = json.Unmarshal(b, &upr)
	if err != nil || upr.Errno != 0 {
		return nil, Errcall(err, string(b), upr)
	}
	return upr, err
}

// 分片上传
func (a *ApiUpload) superfile(req *types.UploadReq, bl *types.BlockList, upr *types.UploadPrecreateResp, superfileStatuFunc types.SuperfileStatuFunc) error {
	sfs := make(chan types.SuperfileStatu, len(*bl.Fsc))
	for i := 0; i < len(*bl.Fsc); i++ {
		_i := i
		_slicing := (*bl.Fsc)[i]
		go func() {
			_sfs := &types.SuperfileStatu{
				Current: _i,
				Slicing: &_slicing,
				Err:     nil,
			}
			//注意：！！！单文件小于4MB时，百度预上传返回的BlockList会是空列表，而不是[0]，因此这里拿分片序号做上传，而不是BlockList[i]
			urlSuperfile := fmt.Sprintf(PanUploadSuperfile,
				a.SdkContext.Token.AccessToken,
				url.QueryEscape(req.RemotePath),
				upr.Uploadid,
				fmt.Sprintf("%v", _i),
			)
			method := "POST"
			payload := &bytes.Buffer{}
			writer := multipart.NewWriter(payload)
			file, err := os.Open(_slicing.Path)
			if err != nil {
				_sfs.Err = Errcall(err, "Error opening file", file)
				sfs <- *_sfs
				return
			}
			defer file.Close()
			part1, err := writer.CreateFormFile("file", _slicing.Path)
			if err != nil {
				_sfs.Err = Errcall(err, part1)
				sfs <- *_sfs
				return
			}
			_, err = io.Copy(part1, file)
			if err != nil {
				_sfs.Err = Errcall(err)
				sfs <- *_sfs
				return
			}
			err = writer.Close()
			if err != nil {
				_sfs.Err = Errcall(err)
				sfs <- *_sfs
				return
			}
			req, err := http.NewRequest(method, urlSuperfile, payload)
			if err != nil {
				_sfs.Err = Errcall(err, req)
				sfs <- *_sfs
				return
			}
			req.Header.Set("Content-Type", writer.FormDataContentType())
			resp, err := a.SdkContext.Client.Do(req)
			if err != nil {
				_sfs.Err = Errcall(err, resp, resp)
				sfs <- *_sfs
				return
			}
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				_sfs.Err = Errcall(err, string(b))
				sfs <- *_sfs
				return
			}
			var superfile types.UploadSuperfileResp
			err = json.Unmarshal(b, &superfile)
			if err != nil {
				_sfs.Err = Errcall(err, string(b), superfile)
				sfs <- *_sfs
				return
			}
			sfs <- *_sfs
		}()
	}

	for i := 0; i < len(*bl.Fsc); i++ {
		s := <-sfs
		if superfileStatuFunc != nil {
			superfileStatuFunc(s, len(*bl.Fsc), i+1)
		}
	}
	return nil
}

func (a *ApiUpload) create(req *types.UploadReq, bl *types.BlockList, upr *types.UploadPrecreateResp) (uploadfile *types.UploadFileResp, err error) {
	values := url.Values{}
	//说明是文件夹
	if upr == nil {
		values.Add("size", "0")
	} else {
		//注意：官方文档里面写要对path进行编码，但是经过测试，编码后反而会导致路径存放不正确
		values.Add("path", req.RemotePath)
		values.Add("size", fmt.Sprintf("%v", req.Size))
		values.Add("block_list", bl.Block)
		values.Add("uploadid", upr.Uploadid)
	}
	values.Add("path", req.RemotePath)
	values.Add("isdir", req.Isdir)
	if req.Rtype != "" {
		values.Add("rtype", req.Rtype)
	}
	resp, err := a.SdkContext.Client.PostForm(fmt.Sprintf(PanUploadCreate,
		a.SdkContext.Token.AccessToken), values)
	if err != nil {
		return nil, Errcall(err, resp)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, Errcall(err, string(b))
	}

	err = json.Unmarshal(b, &uploadfile)
	if err != nil {
		return nil, Errcall(err, uploadfile)
	}
	return
}

// SplitFile 分割文件，最小为4MB，最大为32MB
// path：需要分割的文件路径
// tempdir：切片临时存储的目录
// size：安装指定的大小进行分割
func SplitFile(path string, tempdir string, size int64) (*types.BlockList, error) {
	if size < 1024*1024*4 {
		size = 1024 * 1024 * 4
	}
	if size > 1024*1024*32 {
		size = 1024 * 1024 * 32
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, Errcall(err, f)
	}
	defer f.Close()

	info, _ := os.Stat(path)

	buf := make([]byte, size)

	var fsc []types.FileSlicing

	md5s := ""

	for i := int64(0); i <= info.Size()/size; i++ {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, Errcall(err, f)
		}
		//将分片文件写入到临时目录下
		err = os.WriteFile(fmt.Sprintf("%s/%s_%d", tempdir, info.Name(), i+1), buf[:n], os.ModePerm)
		if err != nil {
			return nil, Errcall(err, f)
		}
		fsc = append(fsc, types.FileSlicing{
			Id:   int(i + 1),
			Name: info.Name(),
			Path: fmt.Sprintf("%s/%s_%d", tempdir, info.Name(), i+1),
			Size: int64(n),
			Md5:  fmt.Sprintf("%x", md5.Sum(buf[:n])),
		})

		md5s += "\"" + fsc[i].Md5 + "\","
	}

	if md5s != "" {

		md5s = "[" + md5s[0:len(md5s)-1] + "]"
	}
	bl := &types.BlockList{
		Fsc:   &fsc,
		Block: md5s,
	}
	return bl, err
}
