package main

import (
	"fmt"
	"genserver/genserver/gencore"
	"genserver/genserver/helper"
	"genserver/genserver/model"
)

func main() {
	mv := model.MyEnv{
		ServerName:      "Email",
		UsePort:         "9233",
		ProjectBasePath: "C:/Users/Administrator/GoProjects/src/solarland/backendv2",
		EntityList: []*model.MyEntity{{
			ModelName: "Email",
			Fields: []*model.MyField{{
				Name: "",
				Type: "",
			}},
		}},
		ClusterPath: "/cluster", // /email/main.go
		//RepositoryPath: "/cluster/email/internal/repository",// /email.go
		//ServicePath:    "/cluster/email/internal/service",// /email.go
		//UsecasePath:    "/cluster/email/internal/usecase", // /email.go
		//EntityPath:     "/cluster/email/internal/domain/entity",// /email.go
		DeployPath:  "/deploy/app/base",              // /email.yaml
		ProtoPath:   "/deploy/app/base/proto/avatar", // /email/grpc.proto
		GraphqlPath: "/deploy/app/base/proto/avatar", // /email/grpc.proto
		GatePath:    "/cluster/gate/gate",            // /email.go
		ConfigPath:  "/cluster/config",
		BundlePath:  "/infra/wireinject/bundle",
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
	if !gencore.Exists(fmt.Sprintf("%v%v", m.ProjectBasePath, m.ClusterPath)) {
		fmt.Println("m.ClusterPath not found")
		return false
	}
	if !gencore.Exists(fmt.Sprintf("%v%v", m.ProjectBasePath, m.DeployPath)) {
		fmt.Println("m.DeployPath not found")
		return false
	}
	if !gencore.Exists(fmt.Sprintf("%v%v", m.ProjectBasePath, m.ProtoPath)) {
		fmt.Println("m.ProtoPath not found")
		return false
	}
	if !gencore.Exists(fmt.Sprintf("%v%v", m.ProjectBasePath, m.GraphqlPath)) {
		fmt.Println("m.GraphqlPath not found")
		return false
	}
	if !gencore.Exists(fmt.Sprintf("%v%v", m.ProjectBasePath, m.GatePath)) {
		fmt.Println("m.GatePath not found")
		return false
	}
	if !gencore.Exists(fmt.Sprintf("%v%v", m.ProjectBasePath, m.ConfigPath)) {
		fmt.Println("m.ConfigPath not found")
		return false
	}
	return true
}
