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
type NeventGenerate struct {
}

func (g *NeventGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*NeventGenerate)(nil)

func (g NeventGenerate) GenCode(env *model.MyEnv) {
	funcMap := map[string]interface{}{}
	inputFiles := []string{"tmpl/nevent.tmpl"}
	filePath := fmt.Sprintf("%v%v/nevent.proto", env.ProtoPath, charater.LowerFirstChar(env.ServerName))
	GenNevent(filePath, env.ServerName, funcMap, inputFiles, env)
}

func GenNevent(outputFile, tmplName string, funcMap template.FuncMap, inputFiles []string, entity *model.MyEnv) {
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
