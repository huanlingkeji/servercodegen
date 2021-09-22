package genItem

import (
	"fmt"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

//
type ConfigGenerate struct {
	configgofile   string
	configyamlfile string
}

func (g *ConfigGenerate) PreCheck(env *model.MyEnv) {
	g.configgofile = fmt.Sprintf("%v%v", env.ConfigPath, "config.go")
	g.configyamlfile = fmt.Sprintf("%v%v", env.ConfigPath, "config.yaml")
	if !gencore.Exists(g.configgofile) {
		panic("no file")
	}
	if !gencore.Exists(g.configyamlfile) {
		panic("no file")
	}
}

var _ IGenerate = (*ConfigGenerate)(nil)

func (g ConfigGenerate) GenCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// config.go 提供接口
		{
			FilePath:     g.configgofile,
			SearchSubStr: ``,
			Content: `// EmailGRPCPort Email server port
func EmailGRPCPort() string {
	return cfg.GetString("service.email.port.grpc")
}
`,
			PInsertType: gencore.FileEnd,
		},
		//  yaml 提供接口
		{
			FilePath:     g.configyamlfile,
			SearchSubStr: `service:`,
			Content: `  email:
    port:
      grpc: 9233
`,
			PInsertType: gencore.StrPointNextLine,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v, env))
	}
}
