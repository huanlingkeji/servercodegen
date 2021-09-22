package main

import (
	"flag"
	"genserver/gqlmodels2pb/cmd"
	"genserver/gqlmodels2pb/env"
)

//	用于在本地生成pb文件

// 命令行参数列表
var cmdLineFlag map[string]string

// init 读取命令行参数
func init() {
	if cmdLineFlag == nil {
		cmdLineFlag = make(map[string]string)
	}
	var cmdName string
	flag.StringVar(&cmdName, "c", "undefine", showTips)
	flag.Parse()

	cmdLineFlag["cmd"] = cmdName
}

func getCmdLineFlag(key string) (string, bool) {
	if v, ok := cmdLineFlag[key]; ok {
		return v, true
	}
	return "", false
}

var showTips = `render xxxx
Gen 生成代码
new-env 生成配置`

func main() {
	cmdName, _ := getCmdLineFlag("cmd")
	switch cmdName {
	case "render":
		_, err := env.InitEnv()
		if err != nil {
			panic(err)
		}
	case "Gen":
		cmd.Gen()
	case "new-env":
		_ = env.PathList{{}, {}}.Encode(`env/env.yaml`)
	}
}
