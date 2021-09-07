package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//CreateFile 创建文件
func CreateFile(filename, data string)  {
	dir := filepath.Dir(filename)
	e := Mkdir(dir)
	if e != nil {
		log.Panicln(e.Error())
		return
	}

	if IsExist(filename) {
		log.Println(fmt.Sprintf("%s already exists", filename))
		return
	}
	file, err := os.Create(filename)
	if err != nil {
		log.Panicln(e.Error())
		return
	}
	file.WriteString(data)
}

//Mkdir 创建文件夹
func Mkdir(path string) error {
	if !IsExist(path) {
		return os.MkdirAll(path, os.ModeDir)
	}

	return nil
}

//IsExist 判断文件或者目录是否存在
func IsExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func UcWorld(name string) string {
	split := strings.Split(name, "/")
	title := strings.Title(split[len(split)-1])
	split[len(split)-1] = title
	return strings.Join(split, "/")
}

//SnakeString 驼峰转下划线
//参考 https://blog.csdn.net/qq_24909089/article/details/107689446
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

//CamelString 下划线转大驼峰
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}