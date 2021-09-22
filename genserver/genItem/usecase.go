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
type UsecaseGenerate struct {
}

func (g *UsecaseGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*UsecaseGenerate)(nil)

func (g UsecaseGenerate) GenCode(env *model.MyEnv) {
	funcMap := map[string]interface{}{}
	inputFiles := []string{"tmpl/usecase.tmpl"}
	for _, v := range env.EntityList {
		filePath := fmt.Sprintf("%v%v/internal/usecase/%v.go", env.ClusterPath,
			charater.LowerFirstChar(env.ServerName), charater.LowerFirstChar(v.ModelName))
		GenUsecase(filePath, v.ModelName, funcMap, inputFiles, v)
	}
}

func GenUsecase(outputFile, tmplName string, funcMap template.FuncMap, inputFiles []string, entity *model.MyEntity) {
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
