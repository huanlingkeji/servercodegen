package core

import (
	"genserver/gqlmodels2pb/convert"
	"genserver/gqlmodels2pb/env"
	"genserver/gqlmodels2pb/model"
	"go/ast"
)

// ast解析model文件的结构
func Visit(myEnv *env.Env, root ast.Node) {
	ast.Inspect(root, func(n ast.Node) bool {
		// 文件
		if file, ok := n.(*ast.File); ok {
			for _, val := range file.Decls {
				Visit(myEnv, val)
			}
		}
		// 模块
		if vnode, ok := n.(*ast.GenDecl); ok {
			for _, val := range vnode.Specs {
				Visit(myEnv, val)
			}
		}
		// type开头的部分
		if vnode, ok := n.(*ast.TypeSpec); ok {
			// struct部分
			if _, ok2 := vnode.Type.(*ast.StructType); ok2 {
				VisitStruct(myEnv, vnode)
			}
			// 别名部分
			if vtai, ok2 := vnode.Type.(*ast.Ident); ok2 {
				myEnv.VariableAliasMap[vnode.Name.Name] = vtai.Name
			}
		}
		return true
	})
}

// 处理结构
func VisitStruct(myEnv *env.Env, node *ast.TypeSpec) {
	structName := node.Name.Name
	sType, ok := node.Type.(*ast.StructType)
	if !ok {
		panic("node.Type.(*ast.StructType)")
	}
	fieldList := sType.Fields.List
	retFieldList := make([]*model.MyField, len(fieldList))
	retFieldMap := make(map[string]*model.MyField, len(fieldList))
	for i, field := range fieldList {
		fieldName := field.Names[0].Name
		myField := &model.MyField{
			IsBaseType: false,
		}
		VisitField(myEnv, myField, field)

		retFieldList[i] = myField
		retFieldMap[fieldName] = myField
	}
	myEnv.ObjectMap[structName] = &model.MyObject{
		Name:     structName,
		FieldArr: retFieldList,
		FieldMap: retFieldMap,
	}
}

// 处理字段
func VisitField(myEnv *env.Env, myField *model.MyField, node ast.Node) interface{} {
	switch node.(type) {
	// 单值类型
	case *ast.Field:
		field := node.(*ast.Field)
		myField.Name = field.Names[0].Name
		return VisitField(myEnv, myField, field.Type)
		// 数组类型
	case *ast.ArrayType:
		myField.IsArr = true
		arrType := node.(*ast.ArrayType)
		return VisitField(myEnv, myField, arrType.Elt)
		// 指针类型
	case *ast.StarExpr:
		myField.IsStar = true
		nextNode := node.(*ast.StarExpr)
		return VisitField(myEnv, myField, nextNode.X)
		// 值类型
	case *ast.Ident:
		nextNode := node.(*ast.Ident)
		myField.Type = nextNode.Name
		myField.IsBaseType = convert.IsBaseType(myEnv.VariableAliasMap, myField.Type)
		return nextNode.Name
		// 表达式类型
	case *ast.SelectorExpr:
		myField.IsSelector = true
		nextNode := node.(*ast.SelectorExpr)
		myField.Type = nextNode.X.(*ast.Ident).Name + "." + nextNode.Sel.Name
		myEnv.SpeStructMap[myField.Type] = struct{}{}
		return myField.Type
		// map类型
	case *ast.MapType:
		myField.IsMap = true
		nextNode := node.(*ast.MapType)
		myField.MKeyType = VisitField(myEnv, myField, nextNode.Key).(string)
		myField.MValueType = VisitField(myEnv, myField, nextNode.Value).(string)
		// 接口类型
	case *ast.InterfaceType:
		myField.Type = "interface{}"
		return "interface{}"
	default:
		panic("unknown type")
	}
	return ""
}
