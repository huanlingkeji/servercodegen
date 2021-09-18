package genItem

import (
	"fmt"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
	"strings"
)

var (
	rpcgofile    = `/rpc.go`
	servergofile = `/server.go`
	wiregofile   = `/wire.go`
)

//
type BundleGenerate struct {
}

func (g BundleGenerate) PreCheck(env *model.MyEnv) {
	if !gencore.Exists(fmt.Sprintf("%v%v%v", env.ProjectBasePath, env.BundlePath, rpcgofile)) {
		panic("no file")
	}
	if !gencore.Exists(fmt.Sprintf("%v%v%v", env.ProjectBasePath, env.BundlePath, servergofile)) {
		panic("no file")
	}
	if !gencore.Exists(fmt.Sprintf("%v%v%v", env.ProjectBasePath, env.BundlePath, wiregofile)) {
		panic("no file")
	}
}

var _ IGenerate = (*BundleGenerate)(nil)

func (g BundleGenerate) GenCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		{
			FilePath:     fmt.Sprintf("%v%v%v", env.ProjectBasePath, env.BundlePath, rpcgofile),
			SearchSubStr: ``,
			Content: fmt.Sprintf(`// %vRPCBundle interface
type %vRPCBundle interface {
}
`, env.ServerName, env.ServerName),
			PInsertType: gencore.FileEnd,
		},
		{
			FilePath:     fmt.Sprintf("%v%v%v", env.ProjectBasePath, env.BundlePath, servergofile),
			SearchSubStr: ``,
			Content: strings.ReplaceAll(`// {{ .ServerName }}ServerBasisBundle type
type {{ .ServerName }}ServerBasisBundle struct {
	ServerBasisBundle
	rpc *t{{ .ServerName }}RPCBundle
}

// RPC func
func (t {{ .ServerName }}ServerBasisBundle) RPC() {{ .ServerName }}RPCBundle {
	return t.rpc
}

// t{{ .ServerName }}RPCBundle type
type t{{ .ServerName }}RPCBundle struct {
}
`, "{{ .ServerName }}", env.ServerName),
			PInsertType: gencore.FileEnd,
		},
		{
			FilePath:     fmt.Sprintf("%v%v%v", env.ProjectBasePath, env.BundlePath, wiregofile),
			SearchSubStr: `WireRPCBasisBundleSet = wire.NewSet(`,
			Content: strings.ReplaceAll(`	wire.Struct(new(t{{ .ServerName }}RPCBundle), "*"),
	wire.Bind(new({{ .ServerName }}RPCBundle), new(*t{{ .ServerName }}RPCBundle)),
`, "{{ .ServerName }}", env.ServerName),
			PInsertType: gencore.StrPointNextLine,
		},
		{
			FilePath:     fmt.Sprintf("%v%v%v", env.ProjectBasePath, env.BundlePath, wiregofile),
			SearchSubStr: `ServiceBasisBundleSet = wire.NewSet(`,
			Content: strings.ReplaceAll(`	wire.Struct(new({{ .ServerName }}ServerBasisBundle), "*"),
`, "{{ .ServerName }}", env.ServerName),
			PInsertType: gencore.StrPointNextLine,
		},
		{
			FilePath:     fmt.Sprintf("%v%v%v", env.ProjectBasePath, env.BundlePath, wiregofile),
			SearchSubStr: `InitializerSet = wire.NewSet(`,
			Content: fmt.Sprintf(`	Initialize%vServerBasisBundle,
`, env.ServerName),
			PInsertType: gencore.StrPointNextLine,
		},
		{
			FilePath:     fmt.Sprintf("%v%v%v", env.ProjectBasePath, env.BundlePath, wiregofile),
			SearchSubStr: `InitializerSet = wire.NewSet(`,
			Content: strings.ReplaceAll(`func Initialize{{ .ServerName }}ServerBasisBundle(ctx context.Context, cfg *viper.Viper) {{ .ServerName }}ServerBasisBundle {
	panic(wire.Build(Set))
}
`, "{{ .ServerName }}", env.ServerName),
			PInsertType: gencore.StrPointNextLine,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v))
	}
}
