package genItem

import (
	"fmt"
	"solarland/backendv2/tools/genserver/model"
	"testing"
)

// GenCode GenCode
type FuncGenCode func(env *model.MyEnv)

// GenCode GenCode
func (f FuncGenCode) GenCode(env *model.MyEnv) {
	f(env)
}

// PreCheck PreCheck
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
