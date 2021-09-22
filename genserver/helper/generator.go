package helper

import (
	"fmt"
	"genserver/genserver/genItem"
	"genserver/genserver/model"
)

func (m *Generator) PreCheck(env *model.MyEnv) {
	for _, v := range m.GenItemList {
		fmt.Printf("%T\n", v)
		v.PreCheck(env)
	}
}

func (m *Generator) GenAll(env *model.MyEnv) {
	for _, v := range m.GenItemList {
		fmt.Printf("%T\n", v)
		v.GenCode(env)
	}
}

type Generator struct {
	GenItemList []genItem.IGenerate
}

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
		},
	}
}
