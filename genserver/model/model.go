package model

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// 环境定义
type MyEnv struct {
	ServerName      string      //微服务的名字
	PortName        string      //端口名字
	ProjectBasePath string      //项目路径
	Entity          []*MyEntity //新加的实体
	ClusterPath     string      //集群文件夹路径
	//RepositoryPath  string      //适配文夹路径
	//ServicePath     string      //服务文夹路径
	//UsecasePath     string      //用例文夹路径
	//EntityPath      string      //实体文夹路径
	DeployPath      string      //部署文夹路径
	ProtoPath       string      //pb文夹路径
	GraphqlPath     string      //graphql文夹路径
	GatePath        string      //gate文夹路径
	ConfigPath      string      //配置文夹路径
}

// 实体
type MyEntity struct {
	Name   string     //实体名字
	Fields []*MyField //实体设计的字段
}

// 字段
type MyField struct {
	Name string //实体的名字
	Type string //实体的类型
}

func (m MyEnv) Encode(filePath string) error {
	bs, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filePath, bs, 666)
	if err != nil {
		return err
	}
	return nil
}

func Decode(filePath string) (*MyEnv, error) {
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	ret := &MyEnv{}
	err = yaml.Unmarshal(bs, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
