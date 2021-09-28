package genItem

import (
	"solarland/backendv2/tools/genserver/model"
)

// IGenerate 生成代码的接口
type IGenerate interface {
	PreCheck(env *model.MyEnv) //  前置的检验
	GenCode(env *model.MyEnv)  //  生成代码
}
