package env

import (
	"genserver/gqlmodels2pb/cmd"
	"genserver/gqlmodels2pb/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type PathList []*ConfPath

type ConfPath struct {
	Name       string // 业务的名字
	Prefix     string // 前缀
	GoFileName string // go文件名
}

// Encode Encode
func (m PathList) Encode(filePath string) error {
	bs, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, bs, 0600)
	if err != nil {
		return err
	}
	return nil
}

// Decode Decode
func Decode(filePath string) (PathList, error) {
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var ret PathList
	err = yaml.Unmarshal(bs, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

// 配置
const (
	RenderObjects = true
	AlwaysCopy    = false
)

type ObjectMap map[string]*model.MyObject // model结构集合
type VariableAliasMap map[string]string   // 变量别名集合
type SpeStructMap map[string]struct{}     // 特殊结构集合
type AllObjectMap map[string]ObjectMap    // asset/bms/gate

type Env struct {
	ObjectMap        map[string]*model.MyObject // model结构集合
	VariableAliasMap map[string]string          // 变量别名集合
	SpeStructMap     map[string]struct{}        // 特殊结构集合
	AllObjectMap     map[string]ObjectMap       // asset/bms/gate
}

func InitEnv() (*Env, error) {
	env := &Env{
		ObjectMap:        make(ObjectMap),
		VariableAliasMap: make(VariableAliasMap),
		SpeStructMap:     make(SpeStructMap),
		AllObjectMap:     make(AllObjectMap),
	}
	pathList, err := Decode(`env/env.yaml`)
	if err != nil {
		return nil, err
	}
	for _, v := range pathList {
		cmd.RenderFile(env, v)
	}
	return env, nil
}

const (
	ProjectPath = "C:/Users/Administrator/GoProjects/src/solarland/backendv2/"
)
