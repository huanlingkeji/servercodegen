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

//
type DeployGenerate struct {
	kustfile      string
	patchyamlfile string
}

func (g *DeployGenerate) PreCheck(env *model.MyEnv) {
	g.kustfile = fmt.Sprintf("%v%v", env.DeployPath, "kustomization.yaml")
	g.patchyamlfile = fmt.Sprintf("%v%v", env.ProtoPath, "deploy/app/local/patch.yaml")
	if !gencore.Exists(g.kustfile) {
		panic("no file")
	}
}

var _ IGenerate = (*DeployGenerate)(nil)

func (g DeployGenerate) GenCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// kustomization.yaml
		{
			FilePath:     g.kustfile,
			SearchSubStr: `resources:`,
			Content:      `  - email.yaml`,
			PInsertType:  gencore.StrPointNextLine,
		},
		{
			FilePath:     g.kustfile,
			SearchSubStr: `resources:`,
			Content: `apiVersion: apps/v1
kind: Deployment
metadata:
  name: email
spec:
  strategy:
    type: Recreate
  replicas: 1
  template:
    spec:
      containers:
        - name: email
          imagePullPolicy: Never
          resources:
            requests:
              cpu: 1m
              memory: 64Mi
            limits:
              cpu: "1"
              memory: 1Gi

---`,
			PInsertType: gencore.FileEnd,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v, env))
	}

	funcMap := map[string]interface{}{}
	inputFiles := []string{"tmpl/deploy.tmpl"}
	filePath := fmt.Sprintf("%v%v.yaml", env.DeployPath, charater.LowerFirstChar(env.ServerName))
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
