package genItem

import (
	"fmt"
	"genserver/genserver/charater"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

// EntityGenerate EntityGenerate
type EntityGenerate struct {
}

// PreCheck PreCheck
func (g *EntityGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*EntityGenerate)(nil)

// GenCode GenCode
func (g EntityGenerate) GenCode(env *model.MyEnv) {

	inputFiles := []string{"tmpl/entity.tmpl"}
	for _, v := range env.EntityList {
		filePath := fmt.Sprintf("%v%v/internal/domain/entity/%v.go", env.ClusterPath,
			charater.LowerFirstChar(env.ServerName), charater.LowerFirstChar(v.ModelName))
		gencore.GenProto(filePath, v.ModelName, inputFiles, env)
	}
}
