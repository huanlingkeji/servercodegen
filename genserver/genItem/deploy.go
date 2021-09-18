package genItem

import (
	"fmt"
	"genserver/genserver/charater"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

var kustfile = `/kustomization.yaml`

//
type DeployGenerate struct {
}

func (g DeployGenerate) PreCheck(env *model.MyEnv) {
	if !gencore.Exists(fmt.Sprintf("%v%v%v", env.ProjectBasePath, env.DeployPath, kustfile)) {
		panic("no file")
	}
}

var _ IGenerate = (*DeployGenerate)(nil)

func (g DeployGenerate) GenCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		{
			FilePath:     fmt.Sprintf("%v%v%v", env.ProjectBasePath, env.DeployPath, kustfile),
			SearchSubStr: `resources:`,
			Content: fmt.Sprintf(`  - %v.yaml
`, charater.LowerFirstChar(env.ServerName)),
			PInsertType: gencore.StrPointNextLine,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v))
	}

	funcMap := map[string]interface{}{}
	inputFiles := []string{"tmpl/deploy.tmpl"}
	filePath := fmt.Sprintf("%v%v/%v.yaml", env.ProjectBasePath, env.DeployPath, charater.LowerFirstChar(env.ServerName))
	GenDeploy(filePath, env.ServerName, funcMap, inputFiles, env)
}

func GenDeploy(outputFile, tmplName string, funcMap template.FuncMap, inputFiles []string, entity *model.MyEnv) {
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
