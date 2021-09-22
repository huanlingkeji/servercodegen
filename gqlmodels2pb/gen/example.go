package gen

import (
	"time"
)

type A struct {
	Age string
}

type BannerShowPolicy string

type BannerInput struct {
	ID           *int                      `json:"id"`
	IDArr        []*int                    `json:"id"`
	IDMap        map[string]*int           `json:"id"`
	IDMapArr     map[string][]*int         `json:"id"` // TODO 暂时不支持复杂 结构  map的嵌套  map和slice的嵌套
	Policy       *BannerShowPolicy         `json:"showPolicy"`
	PolicyArr    []*BannerShowPolicy       `json:"showPolicy"`
	PolicyMap    map[int]*BannerShowPolicy `json:"showPolicy"`
	Person       A
	Persons      []A
	PersonsPtr   []*A
	PersonMap    map[int]*A
	PersonMapInt map[int]interface{}
}

// Email 邮件记录
type Email struct {
	ID        string    `` // ID
	Priority  bool      `` // 优先级
	Content   string    `` // 内容
	SendTime  time.Time `` // 发送时间
	ValidTime time.Time `` // 有效截止时间
	ExtraData string    `` // 额外信息
	//Deleted   bool   // 标记删除
}

// EmailRecord 邮件记录
type EmailRecord struct {
	ID        string `` // ID
	SenderID  string `` // 发送方ID
	ReceiveID string `` // 接受方ID
	GameID    string `` // 游戏ID
	Status    string `` // 状态
	ExtraData string `` // 额外信息
	//Deleted   bool ``  // 标记删除
}

// EmailLike 邮件记录
type EmailLike struct {
	ID          string `` // ID
	LikerPlayer string `` // 点赞方ID
	LikedPlayer string `` // 被点赞方ID
	//Deleted   bool   // 标记删除
}
