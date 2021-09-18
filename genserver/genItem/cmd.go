package genItem

import (
	"fmt"
	"genserver/genserver/charater"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

var cmdgofile =`/cmd/main.go`

//
type CmdGenerate struct {
}

func (g CmdGenerate) PreCheck(env *model.MyEnv) {
	if !gencore.Exists(fmt.Sprintf("%v%v", env.ProjectBasePath, cmdgofile)) {
		panic("no file")
	}
}

var _ IGenerate = (*CmdGenerate)(nil)

func (g CmdGenerate) GenCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// 写入包名
		{
			FilePath:     fmt.Sprintf("%v%v", env.ProjectBasePath, cmdgofile),
			SearchSubStr: `import (`,
			Content: fmt.Sprintf(`	"solarland/backendv2/cluster/%v"
`, charater.LowerFirstChar(env.ServerName)),
			PInsertType: gencore.StrPointNextLine,
		},
		// 写入调用
		{
			FilePath:     fmt.Sprintf("%v%v", env.ProjectBasePath, cmdgofile),
			SearchSubStr: `switch cmdName {`,
			Content: fmt.Sprintf(`	case "%v":
		email.Run(ctx, cfg)
`, charater.LowerFirstChar(env.ServerName)),
			PInsertType: gencore.StrPointNextLine,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v))
	}
}
