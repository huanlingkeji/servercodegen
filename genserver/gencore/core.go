package gencore

import (
	"bytes"
	"errors"
	"genserver/genserver/charater"
	"genserver/genserver/model"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// ContentInsertPosition ContentInsertPosition
type ContentInsertPosition int

// ContentInsertPosition ContentInsertPositions
const (
	StrPointBegin    ContentInsertPosition = 1 // 查找字符串的开始位置
	StrPointEnd      ContentInsertPosition = 2 // 查找字符串的结束位置
	StrPointNextLine ContentInsertPosition = 3 // 查找字符串的下一行开头位置
	FileEnd          ContentInsertPosition = 4 // 文件结尾
)

// InsertContent2NewFile 将内容新到新的文件里面去
func InsertContent2NewFile(filePath, content string) error {
	err := ioutil.WriteFile(filePath, []byte(content), 0600)
	if err != nil {
		return err
	}
	return nil
}

// InsertContentInput InsertContentInput
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

// InsertContent2File 在文件中查找指定内容的位置 然后插入自己的内容
// 首次操作会生成备份 然后会基于备份进行插入内容 即可重复操作
func InsertContent2File(in *InsertContentInput, env *model.MyEnv) error {
	filePath := in.FilePath
	searchSubStr := in.SearchSubStr
	content := templateGen(in.Content, env)
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
	{
		cmprFile := strings.Replace(filePath, env.ProjectBasePath, "./cmpr/", 1)
		dir, _ := filepath.Split(cmprFile)
		_ = os.MkdirAll(dir, 0777)
		fOutput, err := os.Create(cmprFile)
		if err != nil {
			return err
		}
		fOutput.Close()
		err = ioutil.WriteFile(cmprFile, []byte(newData), 0600)
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(filePath, []byte(newData), 0600)
	if err != nil {
		return err
	}
	return nil
}

// GetFilePointBeginIndex 获取文件查找内容的开始位置
func GetFilePointBeginIndex(fileData string, searchSubStr string) int {
	idx := strings.Index(fileData, searchSubStr)
	return idx
}

// GetFilePointEndIndex 获取文件查找内容的结尾位置
func GetFilePointEndIndex(fileData string, searchSubStr string) int {
	idx := strings.Index(fileData, searchSubStr)
	if idx > 0 {
		return idx + len(searchSubStr)
	}
	return idx
}

// GetFilePointNextLineIndex 获取文件查找内容的下一行的位置
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

// CopyBackup 如果文件还没有备份则备份
// find . -name "*.back.txt"  | xargs rm -f
func CopyBackup(prefixPath string) {
	// return
	// backupFileName := fmt.Sprintf("%v.back.txt", prefixPath)
	// if !Exists(backupFileName) {
	//	bs, err := ioutil.ReadFile(prefixPath)
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	err = ioutil.WriteFile(backupFileName, bs, 0600)
	//	if err != nil {
	//		panic(err.Error())
	//	}
	// }
}

// Exists 判断文件路径是存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// CheckErr CheckErr
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// GenProto GenProto
func GenProto(outputFile, tmplName string, inputFiles []string, env *model.MyEnv) {
	genProto(outputFile, tmplName, inputFiles, env)
	cmprFile := strings.Replace(outputFile, env.ProjectBasePath, "./cmpr/", 1)
	genProto(cmprFile, tmplName, inputFiles, env)
}

func genProto(outputFile, tmplName string, inputFiles []string, env *model.MyEnv) {
	if len(inputFiles) > 0 {
		_, fileName := filepath.Split(inputFiles[0])
		tmplName = fileName
	}
	dir, _ := filepath.Split(outputFile)
	_ = os.MkdirAll(dir, 0777)
	fOutput, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("error while opening %q: %v", outputFile, err)
	}
	t, err := template.New(tmplName).Funcs(getFuncMap()).ParseFiles(inputFiles...)
	if err != nil {
		_ = fOutput.Close()
		log.Fatalf("template.ParseFiles %v", err)
	}
	err = t.Execute(fOutput, env)
	if err != nil {
		_ = fOutput.Close()
		log.Fatalf("error while Execute %v", err)
	}
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"LowerFirstChar": charater.LowerFirstChar,
		"UpperFirstChar": charater.UpperFirstChar,
		"UpperAllChar":   charater.UpperAllChar,
	}
}
