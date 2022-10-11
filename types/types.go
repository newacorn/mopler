package types

type UserInfoResp struct {
	BaiduName   string `json:"baidu_name"`   //百度帐号
	NetdiskName string `json:"netdisk_name"` //网盘帐号
	AvatarUrl   string `json:"avatar_url"`   //头像地址
	VipType     int    `json:"vip_type"`     //会员类型，0普通用户、1普通会员、2超级会员
	Uk          int    `json:"uk"`           //用户ID
	Errno       int    `json:"errno"`
	RequestId   any    `json:"request_id"`
	Errmsg      any    `json:"errmsg"`
}

type FileListReq struct {
	//需要list的目录，以/开头的绝对路径, 默认为/路径包含中文时需要UrlEncode编码给出的示例的路径是/测试目录的UrlEncode编码
	Dir string
	/*
		排序字段：默认为name；
			time表示先按文件类型排序，后按修改时间排序；
			name表示先按文件类型排序，后按文件名称排序；
			size表示先按文件类型排序，后按文件大小排序。
	*/
	Order     string
	Desc      string //  默认为升序，设置为1实现降序 （注：排序的对象是当前目录下所有文件，不是当前分页下的文件）
	Start     string //  起始位置，从0开始
	Limit     string //    查询数目，默认为1000，建议最大不超过1000
	Web       string // 值为1时，返回dir_empty属性和缩略图数据
	Folder    string //  是否只返回文件夹，0 返回所有，1 只返回文件夹，且属性只返回path字段
	Showempty string //   是否返回dir_empty属性，0 不返回，1 返回
}

type FileListRecursionReq struct {
	//需要list的目录，以/开头的绝对路径, 默认为/路径包含中文时需要UrlEncode编码给出的示例的路径是/测试目录的UrlEncode编码
	Path string
	/*
		排序字段：默认为name；
			time表示先按文件类型排序，后按修改时间排序；
			name表示先按文件类型排序，后按文件名称排序；
			size表示先按文件类型排序，后按文件大小排序。
	*/
	Order string
	Desc  string //  默认为升序，设置为1实现降序 （注：排序的对象是当前目录下所有文件，不是当前分页下的文件）
	Start string //  起始位置，从0开始
	Limit string //    查询数目，默认为1000，建议最大不超过1000
	Web   string // 值为1时，返回dir_empty属性和缩略图数据
}

type FileListResp struct {
	Errno     int    `json:"errno"`
	GuidInfo  string `json:"guid_info"`
	RequestId int64  `json:"request_id"`
	Errmsg    string `json:"errmsg"`
	Guid      int64  `json:"guid"`
	HasMore   int    `json:"has_more"`
	Cursor    int    `json:"cursor"`
	List      []struct {
		FsId           int64               `json:"fs_id"`           //文件在云端的唯一标识ID
		Path           string              `json:"path"`            //文件的绝对路径
		ServerFilename string              `json:"server_filename"` //文件名称
		Size           int64               `json:"size"`            //文件大小，单位B
		ServerMtime    int64               `json:"server_mtime"`    //文件在服务器修改时间
		ServerCtime    int64               `json:"server_ctime"`    //文件在服务器创建时间
		LocalMtime     int64               `json:"local_mtime"`     //文件在客户端修改时间
		LocalCtime     int64               `json:"local_ctime"`     //文件在客户端创建时间
		Isdir          int                 `json:"isdir"`           //是否为目录，0 文件、1 目录
		Category       int                 `json:"category"`        //文件类型，1 视频、2 音频、3 图片、4 文档、5 应用、6 其他、7 种子
		Md5            string              `json:"md_5"`            //云端哈希（非文件真实MD5），只有是文件类型时，该字段才存在
		DirEmpty       int                 `json:"dir_empty"`       //该目录是否存在子目录，只有请求参数web=1且该条目为目录时，该字段才存在， 0为存在， 1为不存在
		Thumbs         []map[string]string `json:"thumbs"`          //只有请求参数web=1且该条目分类为图片时，该字段才存在，包含三个尺寸的缩略图URL
		Privacy        string              `json:"privacy"`
		Unlist         int64               `json:"unlist"`
		ServerAtime    int64               `json:"server_atime"`
		Share          int64               `json:"share"`
		Empty          int64               `json:"empty"`
		OperId         int64               `json:"oper_id"`
	} `json:"list"`
}

type FileMetaReq struct {
	Fsids     []int64 //	[414244021542671,633507813519281]	文件id数组，数组中元素是uint64类型，数组大小上限是：100
	Dlink     string  //			是否需要下载地址，0为否，1为是，默认为0。获取到dlink后，参考下载文档进行下载操作
	Path      string  //	/123-571234	URL参数	查询共享目录或专属空间内文件时需要。 共享目录格式： /uk-fsid  其中uk为共享目录创建者id， fsid对应共享目录的fsid 专属空间格式：/_pcs_.appdata/xpan/
	Thumb     string  //		是否需要缩略图地址，0为否，1为是，默认为0
	Extra     string  //		图片是否需要拍摄时间、原图分辨率等其他信息，0 否、1 是，默认0
	Needmedia string  //	视频是否需要展示时长信息，0 否、1 是，默认0
}

type FileMetasResp struct {
	Errmsg string         `json:"errmsg"`
	Errno  int            `json:"errno"`
	List   []FileMetaResp `json:"list"`
	Names  struct {
	} `json:"names"`
	RequestId string `json:"request_id"`
}

type FileMetaResp struct {
	Category    int    `json:"category"`
	Filename    string `json:"filename"`
	FsId        int64  `json:"fs_id"`
	Isdir       int    `json:"isdir"`
	Md5         string `json:"md5"`
	OperId      int    `json:"oper_id"`
	Path        string `json:"path"`
	RverCtime   int    `json:"rver_ctime"`
	ServerMtime int    `json:"server_mtime"`
	Size        int    `json:"size"`
	Dlink       string `json:"dlink"`
}

type FileSearchReq struct {
	Key       string `json:"key"`       //	搜索关键字
	Dir       string `json:"dir"`       //	搜索目录，默认根目录
	Page      string `json:"page"`      //	页数，从1开始，缺省则返回所有条目
	Num       string `json:"num"`       //	默认为500，不能修改
	Recursion string `json:"recursion"` //	是否递归搜索子目录 1:是，0:否（默认）
	Web       string `json:"web"`       //	默认0，为1时返回缩略图信息
}

type FileSearchResp struct {
	Errno       int      `json:"errno"`
	RequestId   any      `json:"request_id"`
	Errmsg      any      `json:"errmsg"`
	Contentlist []string `json:"contentlist"`
	HasMore     int      `json:"has_more"`
	List        []struct {
		FsID           int64               `json:"fs_id"`
		Path           string              `json:"path"`
		ServerFilename string              `json:"server_filename"`
		Size           int                 `json:"size"`
		ServerMtime    int                 `json:"server_mtime"`
		ServerCtime    int                 `json:"server_ctime"`
		LocalMtime     int                 `json:"local_mtime"`
		LocalCtime     int                 `json:"local_ctime"`
		Isdir          int                 `json:"isdir"`
		Category       int                 `json:"category"`
		Share          int                 `json:"share"`
		OperID         int                 `json:"oper_id"`
		ExtentTinyint1 int                 `json:"extent_tinyint_1"`
		Md5            string              `json:"md_5"`
		Thumbs         []map[string]string `json:"thumbs"`
		DeleteType     int                 `json:"delete_type"`
		OwnerId        int                 `json:"owner_id"`
		Wpfile         int                 `json:"wpfile"`
	} `json:"list"`
}

type QuotaReq struct {
	//	是否检查免费信息，0为不查，1为查，默认为0
	Checkfree string `json:"checkfree"`
	//		是否检查过期信息，0为不查，1为查，默认为0
	Checkexpire string `json:"checkexpire"`
}

type QuotaResp struct {
	Total     int  `json:"total"`  //总空间大小，单位B
	Expire    bool `json:"expire"` //7天内是否有容量到期
	Used      int  `json:"used"`   //已使用大小，单位B
	Free      int  `json:"free"`   //剩余大小，单位B
	Errno     int  `json:"errno"`
	RequestId any  `json:"request_id"`
	Errmsg    any  `json:"errmsg"`
}

type UploadReq struct {
	SrcPath   string //本地文件的路径
	BlockSize int64
	//	/apps/appName/filename.jpg	RequestBody参数	上传后使用的文件绝对路径，需要urlencode
	RemotePath string
	//	4096	RequestBody参数	文件和目录两种情况：上传文件时，表示文件的大小，单位B；上传目录时，表示目录的大小，目录的话大小默认为0
	Size int64
	//	0	RequestBody参数	是否为目录，0 文件，1 目录
	Isdir string
	/*
		["98d02a0f54781a93e354b1fc85caf488", "ca5273571daefb8ea01a42bfa5d02220"]
		RequestBody参数
		文件各分片MD5数组的json串。block_list的含义如下，如果上传的文件小于4MB，其md5值（32位小写）即为block_list字符串数组的唯一元素；
		如果上传的文件大于4MB，需要将上传的文件按照4MB大小在本地切分成分片，
		不足4MB的分片自动成为最后一个分片，所有分片的md5值（32位小写）组成的字符串数组即为block_list。
	*/
	BlockList string
	//	1	RequestBody参数	固定值1
	Autoinit string
	//	1	RequestBody参数	文件命名策略。
	/*
		1 表示当path冲突时，进行重命名
		2 表示当path冲突且block_list不同时，进行重命名
		3 当云端存在同名文件时，对该文件进行覆盖
	*/
	Rtype string

	//    P1-MTAuMjI4LjQzLjMxOjE1OTU4NTg==    RequestBody参数    上传ID
	Uploadid string
	//	content-md5  b20f8ac80063505f264e5f6fc187e69a    RequestBody参数    文件MD5，32位小写
	ContentMd5 string
	// 9aa0aa691s5c0257c5ab04dd7eddaa47    RequestBody参数    文件校验段的MD5，32位小写，校验段对应文件前256KB
	SliceMd5 string
	//1595919297    RequestBody参数    客户端创建时间， 默认为当前时间戳
	LocalCtime string
	// 1595919297    RequestBody参数    客户端修改时间，默认为当前时间戳
	LocalMtime string
}

type UploadPrecreateResp struct {
	Errno      int    `json:"errno"`       //错误码
	Path       string `json:"path"`        //文件的绝对路径
	Uploadid   string `json:"uploadid"`    //上传唯一ID标识此上传任务
	ReturnType int    `json:"return_type"` //返回类型，系统内部状态字段
	BlockList  []int  `json:"block_list"`  //需要上传的分片序号列表，索引从0开始
}

type UploadSuperfileResp struct {
	Errno int    `json:"errno"` //错误码
	Md5   string `json:"md5"`   //文件切片云端md5
}

type UploadFileResp struct {
	Errno          int    `json:"errno"`           //错误码
	FsId           int64  `json:"fs_id"`           //文件在云端的唯一标识ID
	Md5            string `json:"md_5"`            //文件的MD5，只有提交文件时才返回，提交目录时没有该值
	ServerFilename string `json:"server_filename"` //文件名
	Category       int    `json:"category"`        //分类类型, 1 视频 2 音频 3 图片 4 文档 5 应用 6 其他 7 种子
	Path           string `json:"path"`            //上传后使用的文件绝对路径
	Size           int64  `json:"size"`            //文件大小，单位B
	Ctime          int64  `json:"ctime"`           //文件创建时间
	Mtime          int64  `json:"mtime"`           //文件修改时间
	Isdir          int    `json:"isdir"`           //是否目录，0 文件、1 目录
	Name           string `json:"name"`
	FromType       int    `json:"from_type"`
}

type FileSlicing struct {
	Id   int    `json:"id"`   //切片的序号
	Name string `json:"name"` //文件名
	Path string `json:"path"` //切片的路径
	Size int64  `json:"size"` //切片的大小
	Md5  string `json:"md5"`  //切片的md5值
}

type BlockList struct {
	Fsc   *[]FileSlicing `json:"fsc"`
	Block string         `json:"block"`
}

// SuperfileStatuFunc 分片上传完成的回调
// sfs 分片信息
// current 正在上传第几个分片
// count 总共有多少个分片
type SuperfileStatuFunc func(sfs SuperfileStatu, count int, complete int)

// DownloadStatuFunc 下载完成的回调
type DownloadStatuFunc func(fms FileMetaResp)

type SuperfileStatu struct {
	Current int          //正在上传第几个分片
	Slicing *FileSlicing //当前分片的信息
	Err     error        //上传失败的错误信息
}

type FileManagerInfo struct {
	Path    string `json:"path,omitempty"`
	Dest    string `json:"dest,omitempty"`
	Newname string `json:"newname,omitempty"`
	Ondup   string `json:"ondup,omitempty"`
}

type FileManagerReq struct {
	Async    string            `json:"async,omitempty"`
	Filelist []FileManagerInfo `json:"filelist,omitempty"`
	Ondup    string            `json:"ondup,omitempty"`
}

type FileManagerResp struct {
	Errno     int `json:"errno"`
	RequestId any `json:"request_id"`
	Errmsg    any `json:"errmsg"`
	Info      []struct {
		Errno int    `json:"errno"`
		Path  string `json:"path"`
	} //文件信息
	Taskid uint64 //异步任务id, 当async=2时返回
}
