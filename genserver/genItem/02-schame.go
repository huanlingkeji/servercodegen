package genItem

import (
	"fmt"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

// graphql的结构至少定义一个字段！！！

//
type SchemaGenerate struct {
	schemagraphqlfile string
}

func (g *SchemaGenerate) PreCheck(env *model.MyEnv) {
	g.schemagraphqlfile = fmt.Sprintf("%v%v", env.ProtoPath, "gate/schema.graphql")
	if !gencore.Exists(g.schemagraphqlfile) {
		panic("no file")
	}
	gencore.CopyBackup(g.schemagraphqlfile)
}

var _ IGenerate = (*SchemaGenerate)(nil)

func (g SchemaGenerate) GenCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// query方法
		{
			FilePath:     g.schemagraphqlfile,
			SearchSubStr: `type Query {`,
			Content: `    """{{.ModelZH}}列表"""
    {{ .ModelName }}List(input:{{ .ModelName }}ListInput!):{{ .ModelName }}ListPayload!
    """{{.ModelZH}}"""
    {{ .ModelName }}(input:{{ .ModelName }}Input!):{{ .ModelName }}!
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// mutation方法
		{
			FilePath:     g.schemagraphqlfile,
			SearchSubStr: `type Mutation {`,
			Content: `    """创建{{.ModelZH}}"""
    create{{ .ModelName }}(input:Create{{ .ModelName }}Input!): Boolean!
    """删除{{.ModelZH}}"""
    delete{{ .ModelName }}(input:Delete{{ .ModelName }}Input!): Boolean!
    """修改{{.ModelZH}}"""
    modify{{ .ModelName }}(input:Modify{{ .ModelName }}Input!): Boolean!
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// 结构定义
		{
			FilePath:     g.schemagraphqlfile,
			SearchSubStr: ``,
			Content: `input {{ .ModelName }}ListInput {
	skip:Int!
	limit:Int!
    # TODO 填充自己的结构
  {{ if .ShowExample}}
   #receiverID:String!
   #gameID:String!
   #page:Int!
   #perPageNum:Int!
  {{- end }}
}

type {{ .ModelName }}ListPayload {
   	{{ .ModelName }}List:[{{ .ModelName }}!]
    totalCount:Int!
}

type {{ .ModelName }} {
	id :String!
   # TODO 填充自己的结构
  {{ if .ShowExample}}
    #id :String!
    #content :String!
    #priority :Boolean!
    #sendTime:Int!
    #validTime:Int!
    #receiveIDList :[String!]
    #isReaded :Boolean!
    #isOperate :Boolean!
  {{- end }}
}

input {{ .ModelName }}Input {
	{{ .ModelName }}ID:String!
   # TODO 填充自己的结构
  {{ if .ShowExample}}
    #{{ .ModelName }}ID:String!
    #gameID :String!
  {{- end }}
}

input Create{{ .ModelName }}Input{
	content:String!
   # TODO 填充自己的结构
  {{ if .ShowExample}}
    #senderID:String!
    #content:String!
    #receiverIDList:[String!]
    #isPriority:Boolean!
    #sendTime:Int!
    #gameID:String!
    #validTime:Int!
  {{- end }}
}

input Delete{{ .ModelName }}Input{
	{{ .ModelName }}ID:String!
  # TODO 填充自己的结构
  {{ if .ShowExample}}
	#{{ .ModelName }}ID:String!
    #gameID :String!
  {{- end }}
}

input Modify{{ .ModelName }}Input{
    {{ .ModelName }}ID:String!
  # TODO 填充自己的结构
  {{ if .ShowExample}}
    #gameID:String!
    #receiver:String!
    #isRead:Boolean
    #isOperated:Boolean
  {{- end }}
}

`,
			PInsertType: gencore.FileEnd,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v, env))
	}
}
