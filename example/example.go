package main

import (
	"github.com/766800551/mopler"
	"log"
)

func main() {

	c := mopler.NewConfiguration(
		"xxxxxxxxxxxxxxxxxxxxxxxxx",
		"xxxxxxxxxxxxxxxxxxxxxxxxx",
		mopler.WithConfigurationAccessToken(
			"xxxxxxxxxxxxxxxxxxxxxxxxx",
		),
	)

	//var err error
	sdk := mopler.NewSdkContext(c)

	//获取用户信息
	log.Println(sdk.ApiUserInfo.UserInfo())

	//获取网盘信息
	//log.Println(sdk.ApiQuota.Quota(nil))

	//获取文件列表
	//list, _ := sdk.ApiFileList.FileList(&types.FileListReq{Dir: "/apps/novel"})
	//mopler.PrintToJSON(list)

	//递归获取文件列表
	//list2, _ := sdk.ApiFileList.FileListRecursion(&types.FileListRecursionReq{Path: "/apps/novel"})
	//mopler.PrintToJSON(list2)

	//根据文件Id获取文件信息
	//metas, err := sdk.ApiFileMeta.FileMeta(&types.FileMetaReq{
	//	Fsids: []int64{219452281350038, 334160726082016, 435417451241411},
	//})
	//if err != nil {
	//	log.Println(err)
	//}
	//mopler.PrintToJSON(metas)

	//根据文件Id下载文件
	//_ = sdk.ApiDownload.DownloadByFsId("temp", []int64{763443200151440}, func(fms types.FileMetaResp) {
	//})

	//根据路径下载文件
	//err = sdk.ApiDownload.DownloadByPath("temp", "/apps/novel", func(fms types.FileMetaResp) {
	//	log.Println(fms.Filename)
	//})
	////
	//if err != nil {
	//	log.Println(err)
	//}

	//根据路径递归下载文件
	//err = sdk.ApiDownload.DownloadByPathRecursion("temp", "/apps/novel", func(fms types.FileMetaResp) {
	//	log.Println(fms.Filename)
	//})
	////
	//if err != nil {
	//	log.Println(err)
	//}

	//递归上传文件，可指定多个目录
	//reqs := []*types.UploadReq{
	//	{
	//		SrcPath:    "C:\\Users\\76680\\Desktop\\test.msi",
	//		RemotePath: "/apps/novel",
	//	},
	//	{
	//		SrcPath:    "C:\\Users\\76680\\Desktop\\e",
	//		RemotePath: "/apps/novel",
	//	},
	//	{
	//		SrcPath:    "./test",
	//		RemotePath: "/apps/novel",
	//	},
	//}
	//
	//_, err = sdk.ApiUpload.Upload(reqs, func(sfs types.SuperfileStatu, count int, complete int) {
	//	log.Printf("正在上传：%v ，上传进度：%.0f%%", sfs.Slicing.Name, float32(complete)/float32(count)*100)
	//})
	//if err != nil {
	//	log.Println(err)
	//}

	//创建文件夹
	//ufr, err := sdk.ApiFileManager.CreateDir("7777", "/apps/novel")
	//if err != nil {
	//	log.Println(err)
	//}
	//mopler.PrintToJSON(ufr)

	//删除文件
	//fd, err := sdk.ApiFileManager.FileDelete(&types.FileManagerReq{
	//	Async: "1",
	//	Filelist: []types.FileManagerInfo{
	//		{
	//			Path: "/apps/novel/aaa",
	//		},
	//	},
	//	Ondup: "",
	//})

	//复制文件
	//fc, err := sdk.ApiFileManager.FileCopy(&types.FileManagerReq{
	//	Async: "1",
	//	Filelist: []types.FileManagerInfo{
	//		{
	//			Path:    "/apps/novel/aaa",
	//			Dest:    "/apps/novel",
	//			Newname: "bbb",
	//			Ondup:   "fail",
	//		},
	//	},
	//	Ondup: "",
	//})

	//重命名
	//fr, err := sdk.ApiFileManager.FileRename(&types.FileManagerReq{
	//	Async: "1",
	//	Filelist: []types.FileManagerInfo{
	//		{
	//			Path:    "/apps/novel/aaa",
	//			Newname: "哈哈",
	//		},
	//	},
	//	Ondup: "",
	//})

	//移动文件
	//fm, err := sdk.ApiFileManager.FileMove(&types.FileManagerReq{
	//	Async: "1",
	//	Filelist: []types.FileManagerInfo{
	//		{
	//			Path:    "/apps/novel/aaa",
	//			Dest:    "/apps/novel/bbb",
	//			Newname: "bbb",
	//			Ondup:   "fail",
	//		},
	//	},
	//	Ondup: "",
	//})

	//if err != nil {
	//	log.Println(err)
	//}

}
