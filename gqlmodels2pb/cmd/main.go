package cmd

import (
	"fmt"
	"genserver/gqlmodels2pb/core"
	"genserver/gqlmodels2pb/env"
	"genserver/gqlmodels2pb/file"
	"genserver/gqlmodels2pb/helper"
	"genserver/gqlmodels2pb/model"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"path/filepath"
)

// 将model_gen的model结构所对应的pb结构生成代码写入到对应文件
func RenderFile(myEnv *env.Env, conf *env.ConfPath) {
	subfix := "models_gen.go"
	if conf.GoFileName != "" {
		subfix = conf.GoFileName
	}
	abspath, err := filepath.Abs(conf.Prefix + subfix)
	if err != nil {
		panic(abspath)
	}
	// 创建用于解析源文件的对象
	fset := token.NewFileSet()
	// 解析源文件，返回ast.File原始文档类型的结构体。
	f, err := parser.ParseFile(fset, abspath, nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	core.Visit(myEnv, f)

	if env.RenderObjects {
		content := []byte(helper.RenderObjects(myEnv))
		err = ioutil.WriteFile(conf.Prefix+"models.proto", content, 0600)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(fmt.Sprintf("./tosearch/%v.models.proto", conf.Name), content, 0600)
		if err != nil {
			panic(err)
		}
	}
}

func Gen() {
	myEnv, err := env.InitEnv()
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	prefix := fmt.Sprintf("%v%v", env.ProjectPath, "deploy/app/base/proto/avatar")
	files, err := ioutil.ReadDir(prefix)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	for _, f := range files {
		backPrefixPath := fmt.Sprintf("%v/%v/", prefix, f.Name())
		protoFileName := "grpc.proto"
		protoFilePath := fmt.Sprintf("%v%v", backPrefixPath, protoFileName)
		if !file.Exists(protoFilePath) {
			continue
		}
		file.CopyBackup(backPrefixPath, protoFileName)
		backupFile := fmt.Sprintf("%v%v.back.txt", backPrefixPath, protoFileName)
		bs, err := ioutil.ReadFile(backupFile)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		messageInfos := model.GetMessageInfos(string(bs))
		file.RewriteData2ProtoFile(string(bs), protoFilePath, messageInfos, myEnv)
	}
}
