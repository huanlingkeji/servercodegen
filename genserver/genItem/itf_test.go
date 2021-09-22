package genItem

import (
	"fmt"
	"genserver/genserver/model"
	"testing"
)

type FuncGenCode func(env *model.MyEnv)

func (f FuncGenCode) GenCode(env *model.MyEnv) {
	f(env)
}

func (f FuncGenCode) PreCheck(env *model.MyEnv) {
	f(env)
}

type AA struct {
}

func (a *AA) GenCode(env *model.MyEnv) {
	fmt.Println(env)
}

func TestExample(t *testing.T) {
	var funcGenCode FuncGenCode = func(env *model.MyEnv) {
		fmt.Println(env)
	}

	list := []IGenerate{
		funcGenCode,
		FuncGenCode((*AA)(nil).GenCode),
	}
	env := &model.MyEnv{
		ServerName: "env name",
	}
	for _, v := range list {
		v.GenCode(env)
	}
}
