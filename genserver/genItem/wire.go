package genItem

import (
	"fmt"
	"genserver/genserver/charater"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

// WireGenerate WireGenerate
type WireGenerate struct {
}

// PreCheck PreCheck
func (g *WireGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*WireGenerate)(nil)

// GenCode GenCode
func (g WireGenerate) GenCode(env *model.MyEnv) {
	inputFiles := []string{"tmpl/wire.tmpl"}
	filePath := fmt.Sprintf("%v%v/wire.go", env.ClusterPath, charater.LowerFirstChar(env.ServerName))
	gencore.GenProto(filePath, env.ServerName, inputFiles, env)
}
