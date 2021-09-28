package main

import (
	"fmt"
	"genserver/genserver/cmd"
	"genserver/genserver/gencore"
	"genserver/genserver/helper"
	"genserver/genserver/model"
)

func main() {
	//projectBasePath := "I:/GoProjects/src/solarland/backendv2/"
	//mv := &model.MyEnv{
	//	ServerName: "Avserver",
	//	UsePort:    "9244",
	//	EntityList: []*model.MyEntity{{
	//		ModelName: "MyEntity",
	//		Fields: []*model.MyField{{
	//			Name: "",
	//			Type: "",
	//		}},
	//	}},
	//	ModelName:   "MyEntity",
	//	ModelZH:     "我的结构体",
	//	ShowExample: false,
	//
	//	ProjectBasePath: projectBasePath,
	//	ClusterPath:     fmt.Sprintf("%v%v", projectBasePath, "cluster/"),
	//	DeployPath:      fmt.Sprintf("%v%v", projectBasePath, "deploy/app/base/"),
	//	ProtoPath:       fmt.Sprintf("%v%v", projectBasePath, "deploy/app/base/proto/avatar/"),
	//	GraphqlPath:     fmt.Sprintf("%v%v", projectBasePath, "deploy/app/base/proto/avatar/"),
	//	GatePath:        fmt.Sprintf("%v%v", projectBasePath, "cluster/gate/gate/"),
	//	ConfigPath:      fmt.Sprintf("%v%v", projectBasePath, "cluster/config/"),
	//	BundlePath:      fmt.Sprintf("%v%v", projectBasePath, "infra/wireinject/"),
	//}
	//gencore.CheckErr(mv.Encode("yaml/env.yaml"))

	mv, err := model.Decode("yaml/env.yaml")
	gencore.CheckErr(err)
	if !CheckEnv(mv) {
		panic("path not all right!!!")
	}

	generator := helper.MakeGenerator()
	generator.PreCheck(mv)
	generator.GenAll(mv)
	cmd.GitAdd(mv.ProjectBasePath)
}

// 检测环境是否正常
func CheckEnv(m *model.MyEnv) bool {
	if !CheckPath(m) {
		return false
	}
	return true
}

//检测文件路径是否都存在
func CheckPath(m *model.MyEnv) bool {
	if !gencore.Exists(m.ClusterPath) {
		fmt.Println("m.ClusterPath not found")
		return false
	}
	if !gencore.Exists(m.DeployPath) {
		fmt.Println("m.DeployPath not found")
		return false
	}
	if !gencore.Exists(m.ProtoPath) {
		fmt.Println("m.ProtoPath not found")
		return false
	}
	if !gencore.Exists(m.GraphqlPath) {
		fmt.Println("m.GraphqlPath not found")
		return false
	}
	if !gencore.Exists(m.GatePath) {
		fmt.Println("m.GatePath not found")
		return false
	}
	if !gencore.Exists(m.ConfigPath) {
		fmt.Println("m.ConfigPath not found")
		return false
	}
	return true
}
