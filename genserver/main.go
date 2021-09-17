package main

import (
	"fmt"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

func main() {
	mv := model.MyEnv{
		ServerName:      "",
		PortName:        "",
		ProjectBasePath: "C:/Users/Administrator/GoProjects/src/solarland/backendv2",
		Entity: []*model.MyEntity{{
			Name: "",
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
		ConfigPath:  "/cluster/config/config.yaml",
		// ConfigPath:     "/cluster/config/config.go",
	}
	checkErr(mv.Encode("yaml/env.yaml"))
	if !gencore.CheckPath(&mv) {
		panic("path not all right!!!")
	}
	port := 9233
	insertContent := fmt.Sprintf("\temail:\n\t\tport:%v\n", port)
	checkErr(gencore.InsertContent2File(fmt.Sprintf("%v%v", mv.ProjectBasePath, mv.ConfigPath), "service:", insertContent, gencore.PNextLine))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
