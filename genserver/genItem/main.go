package genItem

import (
	"fmt"
	"solarland/backendv2/tools/genserver/charater"
	"solarland/backendv2/tools/genserver/gencore"
	"solarland/backendv2/tools/genserver/model"
)

// MainGenerate MainGenerate
type MainGenerate struct {
}

// PreCheck PreCheck
func (g *MainGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*MainGenerate)(nil)

// GenCode GenCode
func (g MainGenerate) GenCode(env *model.MyEnv) {

	inputFiles := []string{"tmpl/main.tmpl"}
	filePath := fmt.Sprintf("%v%v/main.go", env.ClusterPath, charater.LowerFirstChar(env.ServerName))
	gencore.GenProto(filePath, env.ServerName, inputFiles, env)
}
