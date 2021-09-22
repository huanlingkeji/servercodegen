package genItem

import (
	"fmt"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

// CmdGenerate CmdGenerate
type CmdGenerate struct {
	cmdgofile string
}

// PreCheck PreCheck
func (g *CmdGenerate) PreCheck(env *model.MyEnv) {
	g.cmdgofile = fmt.Sprintf("%v%v", env.ProjectBasePath, "cmd/main.go")
	if !gencore.Exists(g.cmdgofile) {
		panic("no file")
	}
	gencore.CopyBackup(g.cmdgofile)
}

var _ IGenerate = (*CmdGenerate)(nil)

// GenCode GenCode
func (g CmdGenerate) GenCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// 写入包名
		{
			FilePath:     g.cmdgofile,
			SearchSubStr: `import (`,
			Content: `	"solarland/backendv2/cluster/{{ .ServerName | LowerFirstChar }}"
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// 写入调用
		{
			FilePath:     g.cmdgofile,
			SearchSubStr: `switch cmdName {`,
			Content: `	case "{{ .ServerName | LowerFirstChar }}":
		{{ .ServerName | LowerFirstChar }}.Run(ctx, cfg)
`,
			PInsertType: gencore.StrPointNextLine,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v, env))
	}
}
