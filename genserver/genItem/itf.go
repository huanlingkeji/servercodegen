package genItem

import "genserver/genserver/model"

// 生成代码的接口
type IGenerate interface {
	PreCheck(env *model.MyEnv)	// 前置的检验
	GenCode(env *model.MyEnv)	// 生成代码
}