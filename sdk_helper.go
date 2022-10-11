package mopler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// PrintToJSON 将对象打印成json字符串
func PrintToJSON(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, b, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf.String())
}

func ErrHandler(skip int, msg string, err error) {
	_, file, line, _ := runtime.Caller(skip)
	panic(fmt.Sprintf("\n====================\n"+
		"文件名：%v\n"+
		"行数：%v\n"+
		"网页错误信息：%v\n"+
		"堆栈错误信息：%v\n"+
		"====================\n",
		file, line, msg, err))
}

func Int64SliceToStringSlice(val []int64) []string {
	strs := make([]string, 0)
	for _, v := range val {
		strs = append(strs, fmt.Sprintf("%v", v))
	}
	return strs
}

// PathSeparatorFormat 将所有的\替换为/
func PathSeparatorFormat(s string) string {
	return strings.Replace(s, "\\", "/", -1)
}

func Errcall(err error, anys ...any) error {
	msg := ""
	_, file, line, _ := runtime.Caller(1)
	for i, a := range anys {
		b, e := json.Marshal(a)
		if e != nil {
			msg += fmt.Sprintf("【%v】%v\n", i+1, a)
			continue
		}
		msg += string(b)
	}

	return errors.New(
		fmt.Sprintf("\n====================\n"+
			"文件名：%v\n"+
			"行数：%v\n"+
			"自定义错误：%v\n"+
			"error接口：%v\n"+
			"====================\n",
			file, line, msg, fmt.Sprintf("%v", err)),
	)
}

type TempDir struct {
	DeleteMethod func() //临时文件夹删除操作
	TempName     string //临时文件夹名称
	Err          error  //删除文件时产生的错误
}

// MkdirTemp 返回一个临时文件结构体，实现了删除方法和临时文件夹路径
func MkdirTemp() *TempDir {
	temp, _ := os.MkdirTemp(".", "")
	td := &TempDir{TempName: temp}
	td.DeleteMethod = func() {
		err := os.RemoveAll(temp)
		if err != nil {
			td.Err = Errcall(err, "Failed to remove")
		}
	}
	return td
}
