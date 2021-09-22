package main

import (
	"fmt"
	"genserver/genserver/gencore"
	"genserver/genserver/helper"
	"genserver/genserver/model"
)

func main() {
	projectBasePath := "I:/GoProjects/src/solarland/backendv2/"
	mv := model.MyEnv{
		ServerName: "Email",
		UsePort:    "9233",
		ModelName:  "Email",
		EntityList: []*model.MyEntity{{
			ModelName: "Email",
			Fields: []*model.MyField{{
				Name: "",
				Type: "",
			}},
		}},

		ProjectBasePath: projectBasePath,
		ClusterPath:     fmt.Sprintf("%v%v", projectBasePath, "cluster/"),
		DeployPath:      fmt.Sprintf("%v%v", projectBasePath, "deploy/app/base/"),
		ProtoPath:       fmt.Sprintf("%v%v", projectBasePath, "deploy/app/base/proto/avatar/"),
		GraphqlPath:     fmt.Sprintf("%v%v", projectBasePath, "deploy/app/base/proto/avatar/"),
		GatePath:        fmt.Sprintf("%v%v", projectBasePath, "cluster/gate/gate/"),
		ConfigPath:      fmt.Sprintf("%v%v", projectBasePath, "cluster/config/"),
		BundlePath:      fmt.Sprintf("%v%v", projectBasePath, "infra/wireinject/"),
	}
	gencore.CheckErr(mv.Encode("yaml/env.yaml"))
	if !CheckEnv(&mv) {
		panic("path not all right!!!")
	}

	generator := helper.MakeGenerator()
	generator.PreCheck(&mv)
	generator.GenAll(&mv)
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
