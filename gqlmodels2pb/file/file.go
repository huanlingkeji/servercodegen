package file

import (
	"genserver/gqlmodels2pb/env"
	"genserver/gqlmodels2pb/gen"
	"genserver/gqlmodels2pb/model"
	"io/ioutil"
	"os"
)

// 处理pb文件所有的含有json结构的文件
func RewriteData2ProtoFile(content string, targetFile string, messageInfos []model.PbMessageStructData, myEnv *env.Env) {
	newString := content
	lth := len(messageInfos)
	allDependObjectMap := make(map[string]struct{})
	for i := lth - 1; i >= 0; i-- {
		jsonBeginIdx := messageInfos[i].JsonBeginIndex
		jsonEndIdx := messageInfos[i].EndPos
		genStr, dependObjectMap := gen.GetMessageGenString(messageInfos[i].Name, myEnv, messageInfos[i].RowsNum)
		if genStr == "+" {
			//fmt.Println("in targetFile ", targetFile)
		}
		newString = newString[:jsonBeginIdx] + genStr + newString[jsonEndIdx:]
		for k := range dependObjectMap {
			allDependObjectMap[k] = struct{}{}
		}
	}
	dependStr := ""
	for _, objsMap := range myEnv.AllObjectMap {
		for _, v := range objsMap {
			if _, ok := allDependObjectMap[v.Name]; ok {
				dependStr += gen.GetMessGenStr(v, myEnv.VariableAliasMap)
			}
		}
	}
	newString = newString + dependStr
	err := ioutil.WriteFile(targetFile, []byte(newString), 0600)
	if err != nil {
		panic(err.Error())
	}
}

func CopyBackup(prefixPath string, fileName string) {
	backupFileName := prefixPath + fileName + ".back.txt"
	if !Exists(backupFileName) || env.AlwaysCopy {
		bs, err := ioutil.ReadFile(prefixPath + fileName)
		if err != nil {
			panic(err.Error())
		}
		err = ioutil.WriteFile(backupFileName, bs, 0600)
		if err != nil {
			panic(err.Error())
		}
	}
}

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
