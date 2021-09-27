package genItem

import (
	"fmt"
	"genserver/genserver/charater"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

//
type NeventGenerate struct {
}

func (g *NeventGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*NeventGenerate)(nil)

func (g NeventGenerate) GenCode(env *model.MyEnv) {

	inputFiles := []string{"tmpl/nevent.tmpl"}
	filePath := fmt.Sprintf("%v%v/nevent.proto", env.ProtoPath, charater.LowerFirstChar(env.ServerName))
	gencore.GenProto(filePath, env.ServerName, inputFiles, env)
}
