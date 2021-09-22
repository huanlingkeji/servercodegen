package genItem

import (
	"fmt"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

//
type BundleGenerate struct {
	rpcgofile        string
	servergofile     string
	wiregofile       string
	grpcclientgofile string
	neventgofile     string
}

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
}

var _ IGenerate = (*BundleGenerate)(nil)

func (g BundleGenerate) GenCode(env *model.MyEnv) {
	g.GenRpcCode(env)
	g.GenServerCode(env)
	g.GenWireCode(env)
	g.GenGRPCClientCode(env)
	g.GenNeventCode(env)
}

func (g BundleGenerate) GenRpcCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// 包名
		{
			FilePath:     g.rpcgofile,
			SearchSubStr: `import (`,
			Content:      `emailpb "solarland/backendv2/proto/gen/go/avatar/email"`,
			PInsertType:  gencore.StrPointNextLine,
		},
		// RPCBundle
		{
			FilePath:     g.rpcgofile,
			SearchSubStr: ``,
			Content: `// EmailRPCBundle interface
type EmailRPCBundle interface {
}
`,
			PInsertType: gencore.FileEnd,
		},
		// Gate Depend
		{
			FilePath:     g.rpcgofile,
			SearchSubStr: `type GateRPCBundle interface {`,
			Content: `	Email() emailpb.EmailServiceClient
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// Gate Depend
		{
			FilePath:     g.rpcgofile,
			SearchSubStr: `type tGateRPCBundle struct {`,
			Content: `	IEmail         emailpb.EmailServiceClient
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// rpc realize
		{
			FilePath:     g.rpcgofile,
			SearchSubStr: ``,
			Content: `
func (t *tGateRPCBundle) Email() emailpb.EmailServiceClient {
	return t.IEmail
}
`,
			PInsertType: gencore.FileEnd,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v, env))
	}
}

func (g BundleGenerate) GenServerCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// RPCBundle
		{
			FilePath:     g.servergofile,
			SearchSubStr: ``,
			Content: `// EmailServerBasisBundle type
type EmailServerBasisBundle struct {
	ServerBasisBundle
	rpc *tEmailRPCBundle
}

// RPC func
func (t EmailServerBasisBundle) RPC() EmailRPCBundle {
	return t.rpc
}

// tEmailRPCBundle type
type tEmailRPCBundle struct {
}
`,
			PInsertType: gencore.FileEnd,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v, env))
	}
}

func (g BundleGenerate) GenWireCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// rpc bundle
		{
			FilePath:     g.wiregofile,
			SearchSubStr: `var WireRPCBasisBundleSet = wire.NewSet(`,
			Content: `	wire.Struct(new(tEmailRPCBundle), "*"),
	wire.Bind(new(EmailRPCBundle), new(*tEmailRPCBundle)),
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// service bundle
		{
			FilePath:     g.wiregofile,
			SearchSubStr: `var ServiceBasisBundleSet = wire.NewSet(`,
			Content: `	wire.Struct(new(EmailServerBasisBundle), "*"),
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// service set
		{
			FilePath:     g.wiregofile,
			SearchSubStr: `var InitializerSet = wire.NewSet(`,
			Content: `	InitializeEmailServerBasisBundle,
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// server bundle
		{
			FilePath:     g.wiregofile,
			SearchSubStr: ``,
			Content: `func InitializeEmailServerBasisBundle(ctx context.Context, cfg *viper.Viper) EmailServerBasisBundle {
	panic(wire.Build(Set))
}`,
			PInsertType: gencore.FileEnd,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v, env))
	}
}

func (g BundleGenerate) GenGRPCClientCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// 导入包
		{
			FilePath:     g.grpcclientgofile,
			SearchSubStr: `import (`,
			Content: `	emailpb "solarland/backendv2/proto/gen/go/avatar/email"
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// client
		{
			FilePath:     g.grpcclientgofile,
			SearchSubStr: `var GrpcClientSet = wire.NewSet(`,
			Content: `	ProvideEmailClient,
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// provide client
		{
			FilePath:     g.grpcclientgofile,
			SearchSubStr: ``,
			Content: `func ProvideEmailClient(ctx context.Context, cfg *viper.Viper) emailpb.EmailServiceClient {
	conn, err := infra_grpc.NewClient(ctx, "email:"+config.EmailGRPCPort(), grpc.WithInsecure())
	check(err, "create grpc client for push failed")
	client := emailpb.NewEmailServiceClient(conn)
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

func (g BundleGenerate) GenNeventCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// 导入包
		{
			FilePath:     g.neventgofile,
			SearchSubStr: `import (`,
			Content: `	"solarland/backendv2/proto/gen/go/avatar/email"
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// client
		{
			FilePath:     g.neventgofile,
			SearchSubStr: `var NeventSet = wire.NewSet(`,
			Content: `	email.NewEmailEventClient,
`,
			PInsertType: gencore.StrPointNextLine,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v, env))
	}
}
