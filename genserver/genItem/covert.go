package genItem

import (
	"fmt"
	"solarland/backendv2/tools/genserver/charater"
	"solarland/backendv2/tools/genserver/gencore"
	"solarland/backendv2/tools/genserver/model"
)

// ConvertGenerate ConvertGenerate
type ConvertGenerate struct {
}

// PreCheck PreCheck
func (g *ConvertGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*ConvertGenerate)(nil)

// GenCode GenCode
func (g ConvertGenerate) GenCode(env *model.MyEnv) {
	inputFiles := []string{"tmpl/convert_entity2pb.tmpl"}
	filePath := fmt.Sprintf("%v%v/internal/convert/protobuf.go", env.ClusterPath, charater.LowerFirstChar(env.ServerName))
	gencore.GenProto(filePath, env.ServerName, inputFiles, env)
	inputFiles = []string{"tmpl/convert_pb2gen.tmpl"}
	filePath = fmt.Sprintf("%vconvert/pb2gql_%v.go", env.GatePath, charater.LowerFirstChar(env.ServerName))
	gencore.GenProto(filePath, env.ServerName, inputFiles, env)
}
