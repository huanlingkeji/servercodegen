package gencore

import (
	"errors"
	"fmt"
	"genserver/genserver/model"
	"io/ioutil"
	"os"
	"strings"
)

type FilePosition int

const (
	PBegin    FilePosition = 1
	PEnd      FilePosition = 2
	PNextLine FilePosition = 3
)

//在文件中查找指定内容的位置 然后插入自己的内容
//首次操作会生成备份 然后会基于备份进行插入内容 即可重复操作
func InsertContent2File(filePath, searchSubStr, content string, pType FilePosition) error {
	openFile := filePath
	if exists(fmt.Sprintf("%v.back.txt", filePath)) {
		openFile = fmt.Sprintf("%v.back.txt", filePath)
	} else {
		copyBackup(filePath)
	}
	bs, err := ioutil.ReadFile(openFile)
	if err != nil {
		return err
	}
	indx := -1
	fileData := string(bs)
	switch pType {
	case PBegin:
		indx = getFilePointBeginIndex(fileData, searchSubStr)
	case PEnd:
		indx = getFilePointEndIndex(fileData, searchSubStr)
	case PNextLine:
		indx = getFilePointNextLineIndex(fileData, searchSubStr)
	}
	if indx < 0 {
		return errors.New("can not find position")
	}
	newData := fileData[:indx] + content + fileData[indx:]
	err = ioutil.WriteFile(filePath, []byte(newData), 666)
	if err != nil {
		return err
	}
	return nil
}

//获取文件查找内容的开始位置
func getFilePointBeginIndex(fileData string, searchSubStr string) int {
	idx := strings.Index(fileData, searchSubStr)
	return idx
}

//获取文件查找内容的结尾位置
func getFilePointEndIndex(fileData string, searchSubStr string) int {
	idx := strings.Index(fileData, searchSubStr)
	if idx > 0 {
		return idx + len(searchSubStr)
	}
	return idx
}

//获取文件查找内容的下一行的位置
func getFilePointNextLineIndex(fileData string, searchSubStr string) int {
	idx := strings.Index(fileData, searchSubStr)
	if idx < 0 {
		return -1
	}
	idx2 := strings.Index(fileData[idx:], "\n")
	if idx2 >= 0 {
		return idx + idx2 + 1
	}
	return -1
}

// 如果文件还没有备份则备份
func copyBackup(prefixPath string) {
	backupFileName := fmt.Sprintf("%v.back.txt", prefixPath)
	if !exists(backupFileName) {
		bs, err := ioutil.ReadFile(prefixPath)
		if err != nil {
			panic(err.Error())
		}
		err = ioutil.WriteFile(backupFileName, bs, 0600)
		if err != nil {
			panic(err.Error())
		}
	}
}

//判断文件路径是存在
func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//检测文件路径是否都存在
func CheckPath(m *model.MyEnv) bool {
	if !exists(fmt.Sprintf("%v%v", m.ProjectBasePath, m.ClusterPath)) {
		fmt.Println("m.ClusterPath not found")
		return false
	}
	if !exists(fmt.Sprintf("%v%v", m.ProjectBasePath, m.DeployPath)) {
		fmt.Println("m.DeployPath not found")
		return false
	}
	if !exists(fmt.Sprintf("%v%v", m.ProjectBasePath, m.ProtoPath)) {
		fmt.Println("m.ProtoPath not found")
		return false
	}
	if !exists(fmt.Sprintf("%v%v", m.ProjectBasePath, m.GraphqlPath)) {
		fmt.Println("m.GraphqlPath not found")
		return false
	}
	if !exists(fmt.Sprintf("%v%v", m.ProjectBasePath, m.GatePath)) {
		fmt.Println("m.GatePath not found")
		return false
	}
	if !exists(fmt.Sprintf("%v%v", m.ProjectBasePath, m.ConfigPath)) {
		fmt.Println("m.ConfigPath not found")
		return false
	}
	return true
}
