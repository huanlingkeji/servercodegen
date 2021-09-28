package genItem

import (
	"fmt"
	"solarland/backendv2/tools/genserver/charater"
	"solarland/backendv2/tools/genserver/gencore"
	"solarland/backendv2/tools/genserver/model"
)

// DeployGenerate DeployGenerate
type DeployGenerate struct {
	kustfile      string
	patchyamlfile string
}

// PreCheck PreCheck
func (g *DeployGenerate) PreCheck(env *model.MyEnv) {
	g.kustfile = fmt.Sprintf("%v%v", env.DeployPath, "kustomization.yaml")
	g.patchyamlfile = fmt.Sprintf("%v%v", env.ProjectBasePath, "deploy/app/local/patch.yaml")
	if !gencore.Exists(g.kustfile) {
		panic("no file")
	}
	if !gencore.Exists(g.patchyamlfile) {
		panic("no file")
	}
	gencore.CopyBackup(g.kustfile)
	gencore.CopyBackup(g.patchyamlfile)
}

var _ IGenerate = (*DeployGenerate)(nil)

// GenCode GenCode
func (g DeployGenerate) GenCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// kustomization.yaml
		{
			FilePath:     g.kustfile,
			SearchSubStr: `resources:`,
			Content: `  - {{ .ServerName | LowerFirstChar }}.yaml
`,
			PInsertType: gencore.StrPointNextLine,
		},
		{
			FilePath:     g.patchyamlfile,
			SearchSubStr: ``,
			Content: `

apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .ServerName | LowerFirstChar }}
spec:
  strategy:
    type: Recreate
  replicas: 1
  template:
    spec:
      containers:
        - name: {{ .ServerName | LowerFirstChar }}
          imagePullPolicy: Never
          resources:
            requests:
              cpu: 1m
              memory: 64Mi
            limits:
              cpu: "1"
              memory: 1Gi

---
`,
			PInsertType: gencore.FileEnd,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v, env))
	}

	inputFiles := []string{"tmpl/deploy.tmpl"}
	filePath := fmt.Sprintf("%v%v.yaml", env.DeployPath, charater.LowerFirstChar(env.ServerName))
	gencore.GenProto(filePath, env.ServerName, inputFiles, env)
}
