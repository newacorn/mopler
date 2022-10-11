package mopler

/*
鉴权相关
*/
const (
	DeviceCode            = "DeviceCode"        //设备码
	ImplicitGrant         = "ImplicitGrant"     //简化
	AuthorizationCode     = "AuthorizationCode" //授权码
	PanDeviceCode         = "https://openapi.baidu.com/oauth/2.0/device/code?response_type=device_code&client_id=%v&scope=basic,netdisk"
	PanDeviceToken        = "https://openapi.baidu.com/oauth/2.0/token?grant_type=device_token&code=%v&client_id=%v&client_secret=%v"
	PanDeviceRefreshToken = "https://openapi.baidu.com/oauth/2.0/token?grant_type=refresh_token&refresh_token=%v&client_id=%v&client_secret=%v"
)

/*
网盘基础服务
*/
const (
	// PanUinfo 获取用户信息
	PanUinfo = "https://pan.baidu.com/rest/2.0/xpan/nas?method=uinfo&access_token=%v"
	// PanQuota 查询网盘信息
	PanQuota = "https://pan.baidu.com/api/quota?access_token=%v"
	// PanFileList 获取文件列表
	PanFileList = "https://pan.baidu.com/rest/2.0/xpan/file?method=list&access_token=%v"
	// PanFileListRecursion 递归获取文件列表
	PanFileListRecursion = "https://pan.baidu.com/rest/2.0/xpan/file?method=listall&access_token=%v&recursion=1"
	// PanUploadPrecreate 文件预上传
	PanUploadPrecreate = "http://pan.baidu.com/rest/2.0/xpan/file?method=precreate&access_token=%v"
	// PanUploadSuperfile 分片上传
	PanUploadSuperfile = "https://d.pcs.baidu.com/rest/2.0/pcs/superfile2?method=upload&access_token=%v" +
		"&type=tmpfile&path=%v&uploadid=%v&partseq=%v"
	// PanUploadCreate 文件创建
	PanUploadCreate = "https://pan.baidu.com/rest/2.0/xpan/file?method=create&access_token=%v"
	// PanFileSearch 文件搜索
	PanFileSearch = "http://pan.baidu.com/rest/2.0/xpan/file?access_token=%v&method=search&key=%v"
	// PanFileMeta 文件信息
	PanFileMeta = "http://pan.baidu.com/rest/2.0/xpan/multimedia?access_token=%v&method=filemetas&dlink=%v&fsids=%v"
	// PanFileManager 文件管理
	PanFileManager = "http://pan.baidu.com/rest/2.0/xpan/file?method=filemanager&access_token=%v&opera=%v"
)
