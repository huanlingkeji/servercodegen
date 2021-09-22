package genItem

import (
	"fmt"
	"genserver/genserver/charater"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

// UsecaseGenerate UsecaseGenerate
type UsecaseGenerate struct {
}

// PreCheck PreCheck
func (g *UsecaseGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*UsecaseGenerate)(nil)

// GenCode GenCode
func (g UsecaseGenerate) GenCode(env *model.MyEnv) {

	inputFiles := []string{"tmpl/usecase.tmpl"}
	for _, v := range env.EntityList {
		filePath := fmt.Sprintf("%v%v/internal/usecase/%v.go", env.ClusterPath,
			charater.LowerFirstChar(env.ServerName), charater.LowerFirstChar(v.ModelName))
		gencore.GenProto(filePath, v.ModelName, inputFiles, env)
	}
}
