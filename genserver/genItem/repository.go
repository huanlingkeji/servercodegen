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

type RepositoryGenerate struct {
}

func (g *RepositoryGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*RepositoryGenerate)(nil)

func (g RepositoryGenerate) GenCode(env *model.MyEnv) {
	funcMap := map[string]interface{}{}
	inputFiles := []string{"tmpl/repository.tmpl"}
	for _, v := range env.EntityList {
		filePath := fmt.Sprintf("%v%v/internal/repository/%v.go", env.ClusterPath,
			charater.LowerFirstChar(env.ServerName), charater.LowerFirstChar(v.ModelName))
		GenRepository(filePath, v.ModelName, funcMap, inputFiles, v)
	}
}

func GenRepository(outputFile, tmplName string, funcMap template.FuncMap, inputFiles []string, entity *model.MyEntity) {
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
