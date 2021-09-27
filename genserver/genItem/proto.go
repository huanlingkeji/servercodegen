package genItem

import (
	"fmt"
	"genserver/genserver/charater"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

//
type ProtoGenerate struct {
}

func (g *ProtoGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*ProtoGenerate)(nil)

func (g ProtoGenerate) GenCode(env *model.MyEnv) {

	inputFiles := []string{"tmpl/proto.tmpl"}
	if len(env.EntityList) > 0 {
		filePath := fmt.Sprintf("%v%v/grpc.proto", env.ProtoPath, charater.LowerFirstChar(env.ServerName))
		gencore.GenProto(filePath, env.ServerName, inputFiles, env)

		inputFiles = []string{"tmpl/types.proto.tmpl"}
		filePath = fmt.Sprintf("%v%v/types.proto", env.ProtoPath, charater.LowerFirstChar(env.ServerName))
		gencore.GenProto(filePath, env.ServerName, inputFiles, env)
	}
}
