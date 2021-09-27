package gencore

import (
	"bytes"
	"errors"
	"fmt"
	"genserver/genserver/charater"
	"genserver/genserver/model"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type ContentInsertPosition int

const (
	StrPointBegin    ContentInsertPosition = 1 //查找字符串的开始位置
	StrPointEnd      ContentInsertPosition = 2 //查找字符串的结束位置
	StrPointNextLine ContentInsertPosition = 3 //查找字符串的下一行开头位置
	FileEnd          ContentInsertPosition = 4 //文件结尾
)

// 将内容新到新的文件里面去
func InsertContent2NewFile(filePath, content string) error {
	err := ioutil.WriteFile(filePath, []byte(content), 666)
	if err != nil {
		return err
	}
	return nil
}

type InsertContentInput struct {
	FilePath     string
	SearchSubStr string
	Content      string
	PInsertType  ContentInsertPosition
}

func templateGen(str string, data interface{}) string {
	bf := bytes.NewBuffer(nil)
	t, err := template.New("").Funcs(getFuncMap()).Parse(str)
	if err != nil {
		log.Fatalf("template.ParseFiles %v", err)
	}
	err = t.Execute(bf, data)
	if err != nil {
		log.Fatalf("error while Execute %v", err)
	}
	return bf.String()
}

//在文件中查找指定内容的位置 然后插入自己的内容
//首次操作会生成备份 然后会基于备份进行插入内容 即可重复操作
func InsertContent2File(in *InsertContentInput, data interface{}) error {
	filePath := in.FilePath
	searchSubStr := in.SearchSubStr
	content := templateGen(in.Content, data)
	pType := in.PInsertType
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	indx := -1
	fileData := string(bs)
	switch pType {
	case StrPointBegin:
		indx = GetFilePointBeginIndex(fileData, searchSubStr)
	case StrPointEnd:
		indx = GetFilePointEndIndex(fileData, searchSubStr)
	case StrPointNextLine:
		indx = GetFilePointNextLineIndex(fileData, searchSubStr)
	case FileEnd:
		indx = len(fileData)
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
func GetFilePointBeginIndex(fileData string, searchSubStr string) int {
	idx := strings.Index(fileData, searchSubStr)
	return idx
}

//获取文件查找内容的结尾位置
func GetFilePointEndIndex(fileData string, searchSubStr string) int {
	idx := strings.Index(fileData, searchSubStr)
	if idx > 0 {
		return idx + len(searchSubStr)
	}
	return idx
}

//获取文件查找内容的下一行的位置
func GetFilePointNextLineIndex(fileData string, searchSubStr string) int {
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
//find . -name "*.back.txt"  | xargs rm -f
func CopyBackup(prefixPath string) {
	return
	backupFileName := fmt.Sprintf("%v.back.txt", prefixPath)
	if !Exists(backupFileName) {
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
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GenProto(outputFile, tmplName string, inputFiles []string, env *model.MyEnv) {
	tmplName = ""
	if 0 < len(inputFiles) {
		_, fileName := filepath.Split(inputFiles[0])
		tmplName = fileName
	}
	dir, _ := filepath.Split(outputFile)
	_ = os.MkdirAll(dir, 0777)
	fOutput, err := os.Create(outputFile)
	defer fOutput.Close()
	if err != nil {
		log.Fatalf("error while opening %q: %v", outputFile, err)
	}
	t, err := template.New(tmplName).Funcs(getFuncMap()).ParseFiles(inputFiles...)
	if err != nil {
		log.Fatalf("template.ParseFiles %v", err)
	}
	err = t.Execute(fOutput, env)
	if err != nil {
		log.Fatalf("error while Execute %v", err)
	}
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"LowerFirstChar": charater.LowerFirstChar,
		"UpperFirstChar": charater.UpperFirstChar,
	}
}
