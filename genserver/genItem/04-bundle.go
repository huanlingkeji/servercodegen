package genItem

import (
	"fmt"
	"solarland/backendv2/tools/genserver/gencore"
	"solarland/backendv2/tools/genserver/model"
)

// BundleGenerate BundleGenerate
type BundleGenerate struct {
	rpcgofile        string
	servergofile     string
	wiregofile       string
	grpcclientgofile string
	neventgofile     string
}

// PreCheck PreCheck
func (g *BundleGenerate) PreCheck(env *model.MyEnv) {
	g.rpcgofile = fmt.Sprintf("%v%v", env.BundlePath, `bundle/rpc.go`)
	g.servergofile = fmt.Sprintf("%v%v", env.BundlePath, `bundle/server.go`)
	g.wiregofile = fmt.Sprintf("%v%v", env.BundlePath, `bundle/wire.go`)
	g.grpcclientgofile = fmt.Sprintf("%v%v", env.BundlePath, `fundamental/grpcClient.go`)
	g.neventgofile = fmt.Sprintf("%v%v", env.BundlePath, `fundamental/nevent.go`)
	if !gencore.Exists(g.rpcgofile) {
		panic("no file")
	}
	if !gencore.Exists(g.servergofile) {
		panic("no file")
	}
	if !gencore.Exists(g.wiregofile) {
		panic("no file")
	}
	if !gencore.Exists(g.grpcclientgofile) {
		panic("no file")
	}
	if !gencore.Exists(g.neventgofile) {
		panic("no file")
	}
	gencore.CopyBackup(g.rpcgofile)
	gencore.CopyBackup(g.servergofile)
	gencore.CopyBackup(g.wiregofile)
	gencore.CopyBackup(g.grpcclientgofile)
	gencore.CopyBackup(g.neventgofile)
}

var _ IGenerate = (*BundleGenerate)(nil)

// GenCode GenCode
func (g BundleGenerate) GenCode(env *model.MyEnv) {
	g.GenRPCCode(env)
	g.GenServerCode(env)
	g.GenWireCode(env)
	g.GenGRPCClientCode(env)
	g.GenNeventCode(env)
}

// GenRPCCode GenRPCCode
func (g BundleGenerate) GenRPCCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// 包名
		{
			FilePath:     g.rpcgofile,
			SearchSubStr: `import (`,
			Content: `	{{ .ServerName | LowerFirstChar }}pb "solarland/backendv2/proto/gen/go/avatar/{{ .ServerName | LowerFirstChar }}"
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// RPCBundle
		{
			FilePath:     g.rpcgofile,
			SearchSubStr: ``,
			Content: `// {{ .ServerName  }}RPCBundle interface
type {{ .ServerName  }}RPCBundle interface {
}
`,
			PInsertType: gencore.FileEnd,
		},
		// Gate Depend
		{
			FilePath:     g.rpcgofile,
			SearchSubStr: `type GateRPCBundle interface {`,
			Content: `	{{ .ServerName  }}() {{ .ServerName | LowerFirstChar }}pb.{{ .ServerName  }}ServiceClient
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// Gate Depend
		{
			FilePath:     g.rpcgofile,
			SearchSubStr: `type tGateRPCBundle struct {`,
			Content: `	I{{ .ServerName  }}         {{ .ServerName | LowerFirstChar }}pb.{{ .ServerName  }}ServiceClient
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// rpc realize
		{
			FilePath:     g.rpcgofile,
			SearchSubStr: ``,
			Content: `
func (t *tGateRPCBundle) {{ .ServerName  }}() {{ .ServerName | LowerFirstChar }}pb.{{ .ServerName  }}ServiceClient {
	return t.I{{ .ServerName  }}
}
`,
			PInsertType: gencore.FileEnd,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v, env))
	}
}

// GenServerCode GenServerCode
func (g BundleGenerate) GenServerCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// RPCBundle
		{
			FilePath:     g.servergofile,
			SearchSubStr: ``,
			Content: `// {{ .ServerName  }}ServerBasisBundle type
type {{ .ServerName  }}ServerBasisBundle struct {
	ServerBasisBundle
	rpc *t{{ .ServerName  }}RPCBundle
}

// RPC func
func (t {{ .ServerName  }}ServerBasisBundle) RPC() {{ .ServerName  }}RPCBundle {
	return t.rpc
}

// t{{ .ServerName  }}RPCBundle type
type t{{ .ServerName  }}RPCBundle struct {
}
`,
			PInsertType: gencore.FileEnd,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v, env))
	}
}

// GenWireCode GenWireCode
func (g BundleGenerate) GenWireCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// rpc bundle
		{
			FilePath:     g.wiregofile,
			SearchSubStr: `var WireRPCBasisBundleSet = wire.NewSet(`,
			Content: `	wire.Struct(new(t{{ .ServerName  }}RPCBundle), "*"),
	wire.Bind(new({{ .ServerName  }}RPCBundle), new(*t{{ .ServerName  }}RPCBundle)),
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// service bundle
		{
			FilePath:     g.wiregofile,
			SearchSubStr: `var ServiceBasisBundleSet = wire.NewSet(`,
			Content: `	wire.Struct(new({{ .ServerName  }}ServerBasisBundle), "*"),
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// service set
		{
			FilePath:     g.wiregofile,
			SearchSubStr: `var InitializerSet = wire.NewSet(`,
			Content: `	Initialize{{ .ServerName  }}ServerBasisBundle,
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// server bundle
		{
			FilePath:     g.wiregofile,
			SearchSubStr: ``,
			Content: `func Initialize{{ .ServerName  }}ServerBasisBundle(ctx context.Context, cfg *viper.Viper) {{ .ServerName  }}ServerBasisBundle {
	panic(wire.Build(Set))
}
`,
			PInsertType: gencore.FileEnd,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v, env))
	}
}

// GenGRPCClientCode GenGRPCClientCode
func (g BundleGenerate) GenGRPCClientCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// 导入包
		{
			FilePath:     g.grpcclientgofile,
			SearchSubStr: `import (`,
			Content: `	{{ .ServerName | LowerFirstChar }}pb "solarland/backendv2/proto/gen/go/avatar/{{ .ServerName | LowerFirstChar }}"
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// client
		{
			FilePath:     g.grpcclientgofile,
			SearchSubStr: `var GrpcClientSet = wire.NewSet(`,
			Content: `	Provide{{ .ServerName  }}Client,
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// provide client
		{
			FilePath:     g.grpcclientgofile,
			SearchSubStr: ``,
			Content: `func Provide{{ .ServerName  }}Client(ctx context.Context, cfg *viper.Viper) {{ .ServerName | LowerFirstChar }}pb.{{ .ServerName  }}ServiceClient {
	conn, err := infra_grpc.NewClient(ctx, "{{ .ServerName | LowerFirstChar }}:"+config.{{ .ServerName  }}GRPCPort(), grpc.WithInsecure())
	check(err, "create grpc client for push failed")
	client := {{ .ServerName | LowerFirstChar }}pb.New{{ .ServerName  }}ServiceClient(conn)
	return client
}
`,
			PInsertType: gencore.FileEnd,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v, env))
	}
}

// GenNeventCode GenNeventCode
func (g BundleGenerate) GenNeventCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// 导入包
		{
			FilePath:     g.neventgofile,
			SearchSubStr: `import (`,
			Content: `	"solarland/backendv2/proto/gen/go/avatar/{{ .ServerName | LowerFirstChar }}"
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// client
		{
			FilePath:     g.neventgofile,
			SearchSubStr: `var NeventSet = wire.NewSet(`,
			Content: `	{{ .ServerName | LowerFirstChar }}.New{{ .ServerName }}EventClient,
`,
			PInsertType: gencore.StrPointNextLine,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v, env))
	}
}
