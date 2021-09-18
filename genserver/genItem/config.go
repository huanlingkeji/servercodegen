package genItem

import (
	"fmt"
	"genserver/genserver/charater"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

var configgofile = `/config.go`
var configyamlfile = `/config.yaml`

//
type ConfigGenerate struct {
}

func (g ConfigGenerate) PreCheck(env *model.MyEnv) {
	if !gencore.Exists(fmt.Sprintf("%v%v%v", env.ProjectBasePath, env.ConfigPath, configyamlfile)) {
		panic("no file")
	}
	if !gencore.Exists(fmt.Sprintf("%v%v%v", env.ProjectBasePath, env.ConfigPath, configgofile)) {
		panic("no file")
	}
}

var _ IGenerate = (*ConfigGenerate)(nil)

func (g ConfigGenerate) GenCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		{
			FilePath:     fmt.Sprintf("%v%v%v", env.ProjectBasePath, env.ConfigPath, configgofile),
			SearchSubStr: ``,
			Content: fmt.Sprintf(`// %vGRPCPort %v server port
func %vGRPCPort() string {
	return cfg.GetString("service.%v.port.grpc")
}
`, env.ServerName, env.ServerName, env.ServerName, charater.LowerFirstChar(env.ServerName)),
			PInsertType: gencore.FileEnd,
		},

		{
			FilePath:     fmt.Sprintf("%v%v%v", env.ProjectBasePath, env.ConfigPath, configyamlfile),
			SearchSubStr: `service:`,
			Content: fmt.Sprintf(`  %v:
    port:
      grpc: %v
`, charater.LowerFirstChar(env.ServerName), env.UsePort),
			PInsertType: gencore.StrPointNextLine,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v))
	}
}