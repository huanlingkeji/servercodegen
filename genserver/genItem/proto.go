package genItem

import (
	"fmt"
	"genserver/genserver/charater"
	"genserver/genserver/model"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

//
type ProtoGenerate struct {
}

func (g *ProtoGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*ProtoGenerate)(nil)

func (g ProtoGenerate) GenCode(env *model.MyEnv) {
	funcMap := map[string]interface{}{}
	inputFiles := []string{"tmpl/proto.tmpl"}
	if len(env.EntityList) > 0 {
		filePath := fmt.Sprintf("%v%vgrpc.proto", env.ProtoPath, charater.LowerFirstChar(env.ServerName))
		GenProto(filePath, env.ServerName, funcMap, inputFiles, env.EntityList[0])

		inputFiles = []string{"tmpl/types.proto.tmpl"}
		filePath = fmt.Sprintf("%v%vtypes.proto", env.ProtoPath, charater.LowerFirstChar(env.ServerName))
		GenProto(filePath, env.ServerName, funcMap, inputFiles, env.EntityList[0])
	}
}

func GenProto(outputFile, tmplName string, funcMap template.FuncMap, inputFiles []string, entity *model.MyEntity) {
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
	t, err := template.New(tmplName).Funcs(funcMap).ParseFiles(inputFiles...)
	if err != nil {
		log.Fatalf("template.ParseFiles %v", err)
	}
	err = t.Execute(fOutput, entity)
	if err != nil {
		log.Fatalf("error while Execute %v", err)
	}
}
