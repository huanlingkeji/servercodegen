package convert

import (
	"fmt"
	"genserver/gqlmodels2pb/env"
	"genserver/gqlmodels2pb/model"
)

func TransType(aliasMap env.VariableAliasMap, field *model.MyField) string {
	ret := ""
	if field.IsMap {
		return fmt.Sprintf("map<%v,%v>", field.MKeyType, TransType(aliasMap, &model.MyField{
			Name:       field.Name,
			Type:       field.Type,
			IsBaseType: false,
			IsArr:      field.IsArr,
			IsSelector: field.IsSelector,
			IsStar:     field.IsStar,
			IsMap:      false,
			MKeyType:   "",
			MValueType: "",
		}))
	}
	if field.IsArr {
		ret += "repeated "
	}
	filType := ToBaseType(aliasMap, field.Type)
	star := ""
	if field.IsStar {
		star = "*"
	}
	switch star + filType {
	case "*int":
		ret += "google.protobuf.Int64Value"
	case "*int64":
		ret += "google.protobuf.Int64Value"
	case "*string":
		ret += "google.protobuf.StringValue"
	case "*bool":
		ret += "google.protobuf.BoolValue"
	default:
		ret += filType
	}
	return ret
}

func ToBaseType(aliasMap env.VariableAliasMap, filType string) string {
	val := aliasMap[filType]
	if val != "" {
		filType = val
	}
	switch filType {
	case "int":
		return "int64"
	case "UploadFile":
		return "string"
	}
	return filType
}

func IsBaseType(aliasMap env.VariableAliasMap, s string) bool {
	s = ToBaseType(aliasMap, s)
	switch s {
	case "int", "string", "int32", "int64", "bool":
		return true
	case "*int", "*string", "*int32", "*int64", "*bool":
		return true
	}
	return false
}
