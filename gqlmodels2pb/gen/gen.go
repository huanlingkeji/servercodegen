package gen

import (
	"fmt"
	"genserver/gqlmodels2pb/convert"
	"genserver/gqlmodels2pb/data"
	"genserver/gqlmodels2pb/env"
	"genserver/gqlmodels2pb/model"
	"genserver/gqlmodels2pb/tool"
	"strings"
)

// 获取model生成的pb结构
func GetMessGenStr(v *model.MyObject, aliasMap env.VariableAliasMap) string {
	retStr := ""
	objStr := "message " + v.Name + " {\n"
	for i, fil := range v.FieldArr {
		objStr += fmt.Sprintf("  %v %v = %v;\n", convert.TransType(aliasMap, fil), tool.Ns(fil.Name), i+2)
	}
	retStr += objStr + "}\n\n"
	return retStr
}

// 获取pb结构json替换生成的结构/并获取依赖的结构
func GetMessageGenString(messageName string, myEnv *env.Env, extraRows int) (string, map[string]struct{}) {
	flag := ""
	origMessageName := messageName
	dependObjectMap := make(map[string]struct{})
	if strings.Contains(messageName, "Response") {
		flag = "Response"
		messageName = strings.Replace(messageName, "Response", "", 1)
	} else if strings.Contains(messageName, "Request") {
		flag = "Request"
		messageName = strings.Replace(messageName, "Request", "", 1)
	}
	if flag == "" {
		panic("no Response message " + messageName)
	}
	if val, ok := data.TransMessage2Object[messageName]; ok {
		messageName = val
	}
	// TODO list 类型 待检测
	if val, ok := data.TransMessageListObject[messageName]; ok {
		return fmt.Sprintf("  string deprecated_json = %v;\n  repeated %v list = %v ;\n", 1+extraRows, val, 2+extraRows), map[string]struct{}{val: {}}
	}
	// TODO 检测用的名字不对
	if _, ok := data.NotSameObjectMap[messageName]; ok {
		//fmt.Println("in notSameObjectMap ", messageName)
	}
	for _, objectMap := range myEnv.AllObjectMap {
		var val *model.MyObject
		if flag == "Response" {
			val = getObjectFromMapRes(messageName, objectMap)
		} else if flag == "Request" {
			val = getObjectFromMapIntp(messageName, objectMap)
		}
		if val != nil {
			ret := fmt.Sprintf("  string deprecated_json = %v;\n", 1+extraRows)
			for i, fil := range val.FieldArr {
				ret += fmt.Sprintf("  %v %v = %v;\n", convert.TransType(myEnv.VariableAliasMap, fil), tool.Ns(fil.Name), i+2+extraRows)
				if !fil.IsBaseType {
					dependObjectMap[fil.Type] = struct{}{}
				}
			}
			return ret, dependObjectMap
		}
	}
	fmt.Println("message cant find target object messageName ", messageName, " origMessageName ", origMessageName)
	return fmt.Sprintf("  string deprecated_json = %v;\n", extraRows+1), dependObjectMap
}

// 获取pb结构名字所对应的response model obj
func getObjectFromMapRes(messageName string, objectMap env.ObjectMap) *model.MyObject {
	if val, ok := objectMap[messageName]; ok {
		return val
	}
	if val, ok := objectMap[messageName+"Payload"]; ok {
		return val
	}
	if val, ok := objectMap[messageName+"Response"]; ok {
		return val
	}
	if val, ok := objectMap[messageName+"Result"]; ok {
		return val
	}
	if val, ok := objectMap[messageName+"Reply"]; ok {
		return val
	}
	//if val, ok := objectMap["rpcmsg."+messageName]; ok {
	//	return val
	//}
	return nil
}

// 获取pb结构名字所对应input model obj
func getObjectFromMapIntp(messageName string, objectMap env.ObjectMap) *model.MyObject {
	if val, ok := objectMap[messageName]; ok {
		return val
	}
	if val, ok := objectMap[messageName+"Input"]; ok {
		return val
	}
	if val, ok := objectMap[messageName+"Request"]; ok {
		return val
	}
	//if val, ok := objectMap["rpcmsg."+messageName]; ok {
	//	return val
	//}
	return nil
}
