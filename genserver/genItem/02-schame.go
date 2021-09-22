package genItem

import (
	"fmt"
	"genserver/genserver/gencore"
	"genserver/genserver/model"
)

//
type SchemaGenerate struct {
	schemagraphqlfile string
}

func (g *SchemaGenerate) PreCheck(env *model.MyEnv) {
	g.schemagraphqlfile = fmt.Sprintf("%v%v", env.ProtoPath, "proto/gate/schema.graphql")
	if !gencore.Exists(g.schemagraphqlfile) {
		panic("no file")
	}
}

var _ IGenerate = (*SchemaGenerate)(nil)

func (g SchemaGenerate) GenCode(env *model.MyEnv) {
	insertContentInputArr := []*gencore.InsertContentInput{
		// query方法
		{
			FilePath:     g.schemagraphqlfile,
			SearchSubStr: `type Query {`,
			Content: `    """批量获取用户邮件"""
    emailList(input:EmailListInput!):EmailListPayload!
    """获取用户单封邮件"""
    email(input:EmailInput!):Email!
    """获取邮件的点赞信息"""
    emailLikeList(input:EmailLikeListInput!):EmailLikeListPayload!
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// mutation方法
		{
			FilePath:     g.schemagraphqlfile,
			SearchSubStr: `type Mutation {`,
			Content: `    """发邮件给用户"""
    sendEmail2User(input:SendEmail2UserInput!): Boolean!
    """删除邮件"""
    deleteEmail(input:DeleteEmailInput!): Boolean!
    """改动邮件状态"""
    modifyEmail(input:ModifyEmailInput!): Boolean!
    """邮件点赞"""
    emailLike(input:EmailLikeInput!): Boolean!
`,
			PInsertType: gencore.StrPointNextLine,
		},
		// 结构定义
		{
			FilePath:     g.schemagraphqlfile,
			SearchSubStr: ``,
			Content: `input EmailListInput {
    receiverID:String!
    gameID:String!
    page:Int!
    perPageNum:Int!
}

type EmailListPayload {
    emailList:[Email!]
    totalCount:Int!
}

type Email {
    id :String!
    content :String!
    priority :Boolean!
    sendTime:Int!
    validTime:Int!
    receiveIDList :[String!]
    isReaded :Boolean!
    isOperate :Boolean!
}

type EmailLike {
    emailID :String!
    likerPlayer :String!
    likedPlayer :String!
}

input EmailInput {
    emailID:String!
    gameID :String!
}

input EmailLikeListInput{
    emailID:String!
    gameID :String!
}

type EmailLikeListPayload{
    emailLikeList:[EmailLike]
}

input SendEmail2UserInput{
    senderID:String!
    content:String!
    receiverIDList:[String!]
    isPriority:Boolean!
    sendTime:Int!
    gameID:String!
    validTime:Int!
}

input DeleteEmailInput{
    emailID:String!
    gameID :String!
}

input ModifyEmailInput{
    emailID:String!
    gameID:String!
    receiver:String!
    isRead:Boolean
    isOperated:Boolean
}

input EmailLikeInput {
    emailID:String!
    gameID:String!
    likerPlayer:String!
    likedPlayer:String!
}`,
			PInsertType: gencore.FileEnd,
		},
	}
	for _, v := range insertContentInputArr {
		gencore.CheckErr(gencore.InsertContent2File(v, env))
	}
}
