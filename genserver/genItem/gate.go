package genItem

import (
	"fmt"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

//
type GateGenerate struct {
}

func (g *GateGenerate) PreCheck(env *model.MyEnv) {
}

var _ IGenerate = (*GateGenerate)(nil)

func (g GateGenerate) GenCode(env *model.MyEnv) {

	inputFiles := []string{"tmpl/gate.tmpl"}
	for _, v := range env.EntityList {
		filePath := fmt.Sprintf("%vservice/gate%vService.go", env.GatePath, v.ModelName)
		gencore.GenProto(filePath, v.ModelName, inputFiles, env)
	}
}
