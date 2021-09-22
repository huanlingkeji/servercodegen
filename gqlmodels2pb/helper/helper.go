package helper

import (
	"fmt"
	"genserver/gqlmodels2pb/env"
	"genserver/gqlmodels2pb/gen"
	"genserver/gqlmodels2pb/model"
)

// 输出重复的结构 返回不重复的结构
func PrintRepetitionObjects(allObjectMap env.AllObjectMap) map[string]struct{} {
	repetitionObjectsMap := make(map[string]map[string]*model.MyObject)
	for busName, objectMap := range allObjectMap {
		for k, v := range objectMap {
			if _, ok := repetitionObjectsMap[k]; !ok {
				repetitionObjectsMap[k] = make(map[string]*model.MyObject)
			}
			repetitionObjectsMap[k][busName] = v
		}
	}
	notSameObjectMap := make(map[string]struct{})
	for k, repObjectMap := range repetitionObjectsMap {
		if len(repObjectMap) > 1 {
			for busName := range repObjectMap {
				fmt.Println("存在重复的结构:業務", busName, k)
			}
			retObjName := comprOjects(repObjectMap)
			if retObjName != "" {
				notSameObjectMap[retObjName] = struct{}{}
			}
			fmt.Println()
		}
	}
	return notSameObjectMap
}

// 对比同名结构的差异
func comprOjects(objectMap map[string]*model.MyObject) string {
	a := objectMap["bms"]
	b := objectMap["gate"]
	if a == nil || b == nil {
		return ""
	}
	same := true
	for k, v := range a.FieldMap {
		if _, ok := b.FieldMap[k]; !ok {
			fmt.Printf("gate 少的字段 %v 具体:%+v\n", k, v)
			same = false
		}
	}
	for k, v := range b.FieldMap {
		if _, ok := a.FieldMap[k]; !ok {
			fmt.Printf("bms 少的字段 %v 具体:%+v\n", k, v)
			same = false
		}
	}
	if !same {
		return a.Name
	}
	return ""
}

func RenderObjects(myEnv *env.Env) string {
	retStr := "//<框出特殊的结构 待处理>\n"
	for k := range myEnv.SpeStructMap {
		retStr += fmt.Sprintf("//%v\n", k)
	}
	retStr += "//<前面的是特殊的结构>\n\n\n\n\n"
	for _, v := range myEnv.ObjectMap {
		retStr += gen.GetMessGenStr(v, myEnv.VariableAliasMap)
	}
	return retStr
}
