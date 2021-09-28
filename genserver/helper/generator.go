package helper

import (
	"fmt"
	"solarland/backendv2/tools/genserver/genItem"
	"solarland/backendv2/tools/genserver/model"
)

// PreCheck PreCheck
func (m *Generator) PreCheck(env *model.MyEnv) {
	for _, v := range m.GenItemList {
		fmt.Printf("%T\n", v)
		v.PreCheck(env)
	}
}

// GenAll GenAll
func (m *Generator) GenAll(env *model.MyEnv) {
	for _, v := range m.GenItemList {
		fmt.Printf("%T\n", v)
		v.GenCode(env)
	}
}

// Generator Generator
type Generator struct {
	GenItemList []genItem.IGenerate
}

// MakeGenerator MakeGenerator
func MakeGenerator() *Generator {
	return &Generator{
		GenItemList: []genItem.IGenerate{
			&genItem.BundleGenerate{},
			&genItem.CmdGenerate{},
			&genItem.ConfigGenerate{},
			&genItem.EntityGenerate{},
			&genItem.GateGenerate{},
			&genItem.DeployGenerate{},
			&genItem.MainGenerate{},
			&genItem.ProtoGenerate{},
			&genItem.RepositoryGenerate{},
			&genItem.ServiceGenerate{},
			&genItem.UsecaseGenerate{},
			&genItem.WireGenerate{},
			&genItem.NeventGenerate{},
			&genItem.ConvertGenerate{},
			&genItem.SchemaGenerate{},
		},
	}
}
